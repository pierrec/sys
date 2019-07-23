// cgo -godefs types_aix.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build ppc64,aix

package unix

type (
	_C_long int64
)

type off int64

type Timespec struct {
	Sec  int64
	Nsec int64
}

type Timeval struct {
	Sec  int64
	Usec int32
	_    [4]byte
}

type Time_t int64

type Utimbuf struct {
	Actime  int64
	Modtime int64
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int64
	Ixrss    int64
	Idrss    int64
	Isrss    int64
	Minflt   int64
	Majflt   int64
	Nswap    int64
	Inblock  int64
	Oublock  int64
	Msgsnd   int64
	Msgrcv   int64
	Nsignals int64
	Nvcsw    int64
	Nivcsw   int64
}

type dev_t uint64

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Mode     uint32
	Nlink    int16
	Flag     uint16
	Uid      uint32
	Gid      uint32
	Rdev     uint64
	Ssize    int32
	Atim     Timespec
	Mtim     Timespec
	Ctim     Timespec
	Blksize  int64
	Blocks   int64
	Vfstype  int32
	Vfs      uint32
	Type     uint32
	Gen      uint32
	Reserved [9]uint32
	Padto_ll uint32
	Size     int64
}

type Dirent struct {
	Offset uint64
	Ino    uint64
	Reclen uint16
	Namlen uint16
	Name   [256]uint8
	_      [4]byte
}

/* in_addr */

/* in6_addr */

type Iovec struct {
	Base *byte
	Len  uint64
}

/* in_addr */
/* in_addr */

/* in6_addr */

type FdSet struct {
	Bits [1024]int64
}

type Sigset_t struct {
	Set [4]uint64
}

type Statfs_t struct {
	Version   int32
	Type      int32
	Bsize     uint64
	Blocks    uint64
	Bfree     uint64
	Bavail    uint64
	Files     uint64
	Ffree     uint64
	Fsid      Fsid64_t
	Vfstype   int32
	Fsize     uint64
	Vfsnumber int32
	Vfsoff    int32
	Vfslen    int32
	Vfsvers   int32
	Fname     [32]uint8
	Fpack     [32]uint8
	Name_max  int32
	_         [4]byte
}
