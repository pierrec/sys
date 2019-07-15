package merge

import (
	"fmt"
	"strings"
)

// builder adds a Printf method to strings.Builder for convenience.
type builder struct {
	strings.Builder
}

// Newline writes a new line to the underlying Builder.
func (b *builder) Newline() {
	b.Builder.WriteByte('\n')
}

// Printf follows the same usage as fmt.Printf.
func (b *builder) Printf(f string, args ...interface{}) {
	s := fmt.Sprintf(f, args...)
	b.Builder.WriteString(s)
}
