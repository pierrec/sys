// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package merge

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"io"
	"sort"
	"strings"
)

type (
	// Registry keeps the files merge is interested in grouped by their name:
	// - <name>_<goos>_<goarch>.go are the interesting files
	// - <name>_<goos>.go is the name of the file that will contain the merged objects and points at a *gofile
	Registry struct {
		m map[string]*gofile
	}
	// goarch represents an arch file.
	goarch struct {
		*ast.File
		kinds // Local objects for an arch
	}
	// gofile holds all the arch dependent files for a given interesting file.
	gofile struct {
		arch    map[string]*goarch
		kinds            // Factorized objects for all arch
		imports []string // Imported package names
	}
)

// NewRegistry creates a new Registry and loads the types to be consolidated.
func NewRegistry(pkg *ast.Package) *Registry {
	reg := make(map[string]*gofile)

	// Group files by name_os and arch.
	for fname, file := range pkg.Files {
		// Group by name_os and arch.
		i := strings.LastIndexByte(fname, '_')
		name := fname[:i]
		arch := fname[i+1:]

		if reg[name] == nil {
			reg[name] = &gofile{arch: make(map[string]*goarch)}
		}
		reg[name].arch[arch] = &goarch{File: file}
	}

	r := &Registry{m: reg}
	r.buildKinds()

	return r
}

// String lists the merged files and their architectures. Used for debugging.
func (r Registry) String() string {
	var b strings.Builder
	var fnames []string
	for file, gf := range r.m {
		b.WriteString(file)
		b.WriteByte('\n')
		for k := range gf.arch {
			fnames = append(fnames, k)
		}
		sort.Strings(fnames)
		for _, fname := range fnames {
			b.WriteString(fmt.Sprintf("  %s\n", fname))
		}
		fnames = fnames[:0]
	}
	return b.String()
}

// Build consolidates the objects.
func (r *Registry) Build() {
	// Build the intersection of all objects for all arch.
	for _, gf := range r.m {
		for _, ga := range gf.arch {
			gf.kinds.inter(&ga.kinds)
		}
	}
	// Remove consolidated objects for all arch.
	for _, gf := range r.m {
		for _, ga := range gf.arch {
			ga.kinds.diff(&gf.kinds)
			// Update the input file ast.
			trimFile(ga.File, &ga.kinds)
		}
	}
}

// buildKinds populates the kinds of every arch: constants, types and functions.
func (r *Registry) buildKinds() {
	for _, gf := range r.m {
		for _, ga := range gf.arch {
			k := &ga.kinds
			for _, decl := range ga.File.Decls {
				switch d := decl.(type) {
				case *ast.GenDecl:
					switch d.Tok {
					case token.CONST:
						k.pushConst(d)

					case token.TYPE:
						k.pushType(d)
					}

				case *ast.FuncDecl:
					k.pushFunc(d)
				}
			}
		}
	}
}

// buildImports records all imports required for files to be newly created.
func (r *Registry) buildImports() {
	for _, gf := range r.m {
		// Aggregate all the imports for a given OS and ARCH.
		for _, ga := range gf.arch {
			for _, imp := range ga.File.Imports {
				name := imp.Path.Value
				if !stringIn(name, gf.imports) {
					gf.imports = append(gf.imports, name)
				}
			}
		}
		// Keep the imports used in the new file.
		for i := 0; i < len(gf.imports); {
			if gf.kinds.hasImport(gf.imports[i]) {
				i++
				continue
			}
			gf.imports = stringDelAt(gf.imports, i)
		}
	}
}

// Print outputs the consolidated objects into their own file, and updates the correspoonding os_arch files.
func (r *Registry) Print(pkg *ast.Package, fset *token.FileSet) error {
	header := fmt.Sprintf(`// Generated code. DO NOT EDIT.

package %s

`, pkg.Name)

	handleClose := func(errp *error, c io.Closer) {
		if err := c.Close(); err != nil && *errp == nil {
			*errp = err
		}
	}

	doInput := func(name string, file *ast.File) (err error) {
		f, err := newbufferedFile(name)
		if err != nil {
			return err
		}
		defer handleClose(&err, f)
		return printer.Fprint(f, fset, file)
	}
	do := func(name string, gf *gofile) (err error) {
		f, err := newbufferedFile(name + ".go")
		if err != nil {
			return err
		}
		defer handleClose(&err, f)
		// Print header.
		if _, err := fmt.Fprintf(f, header); err != nil {
			return err
		}
		// Print imports, if any.
		switch {
		case len(gf.imports) == 1:
			name := gf.imports[0]
			if _, err := fmt.Fprintf(f, "import %s\n\n", name); err != nil {
				return err
			}

		case len(gf.imports) > 1:
			if _, err := fmt.Fprintf(f, "import (\n"); err != nil {
				return err
			}
			for _, name := range gf.imports {
				if _, err := fmt.Fprintf(f, "%s\n", name); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprintf(f, ")\n\n"); err != nil {
				return err
			}
		}
		// Print consolidated objects.
		if err := gf.kinds.print(f); err != nil {
			return err
		}
		// Update input files.
		for arch, ga := range gf.arch {
			name := fmt.Sprintf("%s_%s", name, arch)
			if err := doInput(name, ga.File); err != nil {
				return err
			}
		}
		return nil
	}
	r.buildImports()
	for name, gf := range r.m {
		if err := do(name, gf); err != nil {
			return err
		}
	}
	return nil
}

// Stats populates the statistics.
func (r *Registry) ReadStats(s *Stats) {
	s.clear()
	for name, gf := range r.m {
		var as AggStats
		as.FileStats.set(name, &gf.kinds)
		for arch, ga := range gf.arch {
			var fs FileStats
			fs.set(arch, &ga.kinds)
			as.Arch = append(as.Arch, fs)
		}
		// Sort by file name to ease the readings.
		sort.Slice(as.Arch, func(i, j int) bool {
			return as.Arch[i].Name < as.Arch[j].Name
		})
		s.Agg = append(s.Agg, as)
	}
	// Sort by file name to ease the readings.
	sort.Slice(s.Agg, func(i, j int) bool {
		return s.Agg[i].FileStats.Name < s.Agg[j].FileStats.Name
	})
}

// trimFile removes objects from f that are not in k.
func trimFile(f *ast.File, k *kinds) {
	for i := 0; i < len(f.Decls); i++ {
		switch d := f.Decls[i].(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				for i := 0; i < len(d.Specs); i++ {
					val := d.Specs[i].(*ast.ValueSpec)
					if len(k.consts) == 0 {
						d.Specs = typeDelAt(d.Specs, i)
						i--
						continue
					}
					for _, v := range k.consts {
						if exprEqual(val.Type, v.Type) {
							valInter(val, v)
							if len(val.Names) == 0 {
								d.Specs = typeDelAt(d.Specs, i)
								i--
							}
							break
						}
					}
				}
				if len(d.Specs) == 0 {
					f.Decls = declDelAt(f.Decls, i)
					i--
				}

			case token.TYPE:
				for i := 0; i < len(d.Specs); {
					spec := d.Specs[i].(*ast.TypeSpec)
					if typeIn(spec, k.types) {
						i++
						continue
					}
					d.Specs = typeDelAt(d.Specs, i)
				}
				if len(d.Specs) == 0 {
					f.Decls = declDelAt(f.Decls, i)
					i--
				}
			}

		case *ast.FuncDecl:
			if !funcIn(d, k.funcs) {
				f.Decls = declDelAt(f.Decls, i)
				i--
			}
		}
	}
}

//-----------------------------------------------------------------------------

func stringIn(key string, s []string) bool {
	for _, k := range s {
		if k == key {
			return true
		}
	}
	return false
}

const nostring = ""

func stringDelAt(s []string, i int) []string {
	if i+1 < len(s) {
		copy(s[i:], s[i+1:])
	}
	s[len(s)-1] = nostring
	return s[:len(s)-1]
}
