// cgo -godefs types_aix.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build ppc,aix

package unix

type (
	_C_long int32
)

type off int32

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type Time_t int32

type Utimbuf struct {
	Actime  int32
	Modtime int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type dev_t uint32

type Stat_t struct {
	Dev      uint32
	Ino      uint32
	Mode     uint32
	Nlink    int16
	Flag     uint16
	Uid      uint32
	Gid      uint32
	Rdev     uint32
	Size     int32
	Atim     Timespec
	Mtim     Timespec
	Ctim     Timespec
	Blksize  int32
	Blocks   int32
	Vfstype  int32
	Vfs      uint32
	Type     uint32
	Gen      uint32
	Reserved [9]uint32
}

type Dirent struct {
	Offset uint32
	Ino    uint32
	Reclen uint16
	Namlen uint16
	Name   [256]uint8
}

/* in_addr */

/* in6_addr */

type Iovec struct {
	Base *byte
	Len  uint32
}

/* in_addr */
/* in_addr */

/* in6_addr */

type FdSet struct {
	Bits [2048]int32
}

type Sigset_t struct {
	Losigs uint32
	Hisigs uint32
}

type Statfs_t struct {
	Version   int32
	Type      int32
	Bsize     uint32
	Blocks    uint32
	Bfree     uint32
	Bavail    uint32
	Files     uint32
	Ffree     uint32
	Fsid      Fsid_t
	Vfstype   int32
	Fsize     uint32
	Vfsnumber int32
	Vfsoff    int32
	Vfslen    int32
	Vfsvers   int32
	Fname     [32]uint8
	Fpack     [32]uint8
	Name_max  int32
}
