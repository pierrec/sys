// mkerrors.sh -Wall -Werror -static -I/tmp/include
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build mipsle,linux

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Wall -Werror -static -I/tmp/include _const.go

package unix

import "syscall"

const (
	MAP_RENAME = 0x800

	PTRACE_GETFPREGS = 0xe

	PTRACE_GET_THREAD_AREA      = 0x19
	PTRACE_GET_THREAD_AREA_3264 = 0xc4
	PTRACE_GET_WATCH_REGS       = 0xd0

	PTRACE_OLDSETOPTIONS = 0x15

	PTRACE_PEEKDATA_3264 = 0xc1

	PTRACE_PEEKTEXT_3264 = 0xc0

	PTRACE_POKEDATA_3264 = 0xc3

	PTRACE_POKETEXT_3264 = 0xc2

	PTRACE_SETFPREGS = 0xf

	PTRACE_SET_THREAD_AREA = 0x1a
	PTRACE_SET_WATCH_REGS  = 0xd1

	SO_STYLE = 0x1008

	TCGETS2 = 0x4030542a

	TCSETS2 = 0x8030542b

	TCSETSF2 = 0x8030542d

	TCSETSW2 = 0x8030542c

	TIOCGETP = 0x7408

	TIOCGLTC = 0x7474

	TIOCSER_TEMT = 0x1

	TIOCSETN = 0x740a
	TIOCSETP = 0x7409

	TIOCSLTC = 0x7475

	VSWTCH = 0x7
)

// Errors
const (
	EINIT = syscall.Errno(0x8d)

	EREMDEV = syscall.Errno(0x8e)
)

// Signals
const (
	SIGEMT = syscall.Signal(0x7)
)

// Error table
var errorList = [...]struct {
	num  syscall.Errno
	name string
	desc string
}{
	{1, "EPERM", "operation not permitted"},
	{2, "ENOENT", "no such file or directory"},
	{3, "ESRCH", "no such process"},
	{4, "EINTR", "interrupted system call"},
	{5, "EIO", "input/output error"},
	{6, "ENXIO", "no such device or address"},
	{7, "E2BIG", "argument list too long"},
	{8, "ENOEXEC", "exec format error"},
	{9, "EBADF", "bad file descriptor"},
	{10, "ECHILD", "no child processes"},
	{11, "EAGAIN", "resource temporarily unavailable"},
	{12, "ENOMEM", "cannot allocate memory"},
	{13, "EACCES", "permission denied"},
	{14, "EFAULT", "bad address"},
	{15, "ENOTBLK", "block device required"},
	{16, "EBUSY", "device or resource busy"},
	{17, "EEXIST", "file exists"},
	{18, "EXDEV", "invalid cross-device link"},
	{19, "ENODEV", "no such device"},
	{20, "ENOTDIR", "not a directory"},
	{21, "EISDIR", "is a directory"},
	{22, "EINVAL", "invalid argument"},
	{23, "ENFILE", "too many open files in system"},
	{24, "EMFILE", "too many open files"},
	{25, "ENOTTY", "inappropriate ioctl for device"},
	{26, "ETXTBSY", "text file busy"},
	{27, "EFBIG", "file too large"},
	{28, "ENOSPC", "no space left on device"},
	{29, "ESPIPE", "illegal seek"},
	{30, "EROFS", "read-only file system"},
	{31, "EMLINK", "too many links"},
	{32, "EPIPE", "broken pipe"},
	{33, "EDOM", "numerical argument out of domain"},
	{34, "ERANGE", "numerical result out of range"},
	{35, "ENOMSG", "no message of desired type"},
	{36, "EIDRM", "identifier removed"},
	{37, "ECHRNG", "channel number out of range"},
	{38, "EL2NSYNC", "level 2 not synchronized"},
	{39, "EL3HLT", "level 3 halted"},
	{40, "EL3RST", "level 3 reset"},
	{41, "ELNRNG", "link number out of range"},
	{42, "EUNATCH", "protocol driver not attached"},
	{43, "ENOCSI", "no CSI structure available"},
	{44, "EL2HLT", "level 2 halted"},
	{45, "EDEADLK", "resource deadlock avoided"},
	{46, "ENOLCK", "no locks available"},
	{50, "EBADE", "invalid exchange"},
	{51, "EBADR", "invalid request descriptor"},
	{52, "EXFULL", "exchange full"},
	{53, "ENOANO", "no anode"},
	{54, "EBADRQC", "invalid request code"},
	{55, "EBADSLT", "invalid slot"},
	{56, "EDEADLOCK", "file locking deadlock error"},
	{59, "EBFONT", "bad font file format"},
	{60, "ENOSTR", "device not a stream"},
	{61, "ENODATA", "no data available"},
	{62, "ETIME", "timer expired"},
	{63, "ENOSR", "out of streams resources"},
	{64, "ENONET", "machine is not on the network"},
	{65, "ENOPKG", "package not installed"},
	{66, "EREMOTE", "object is remote"},
	{67, "ENOLINK", "link has been severed"},
	{68, "EADV", "advertise error"},
	{69, "ESRMNT", "srmount error"},
	{70, "ECOMM", "communication error on send"},
	{71, "EPROTO", "protocol error"},
	{73, "EDOTDOT", "RFS specific error"},
	{74, "EMULTIHOP", "multihop attempted"},
	{77, "EBADMSG", "bad message"},
	{78, "ENAMETOOLONG", "file name too long"},
	{79, "EOVERFLOW", "value too large for defined data type"},
	{80, "ENOTUNIQ", "name not unique on network"},
	{81, "EBADFD", "file descriptor in bad state"},
	{82, "EREMCHG", "remote address changed"},
	{83, "ELIBACC", "can not access a needed shared library"},
	{84, "ELIBBAD", "accessing a corrupted shared library"},
	{85, "ELIBSCN", ".lib section in a.out corrupted"},
	{86, "ELIBMAX", "attempting to link in too many shared libraries"},
	{87, "ELIBEXEC", "cannot exec a shared library directly"},
	{88, "EILSEQ", "invalid or incomplete multibyte or wide character"},
	{89, "ENOSYS", "function not implemented"},
	{90, "ELOOP", "too many levels of symbolic links"},
	{91, "ERESTART", "interrupted system call should be restarted"},
	{92, "ESTRPIPE", "streams pipe error"},
	{93, "ENOTEMPTY", "directory not empty"},
	{94, "EUSERS", "too many users"},
	{95, "ENOTSOCK", "socket operation on non-socket"},
	{96, "EDESTADDRREQ", "destination address required"},
	{97, "EMSGSIZE", "message too long"},
	{98, "EPROTOTYPE", "protocol wrong type for socket"},
	{99, "ENOPROTOOPT", "protocol not available"},
	{120, "EPROTONOSUPPORT", "protocol not supported"},
	{121, "ESOCKTNOSUPPORT", "socket type not supported"},
	{122, "ENOTSUP", "operation not supported"},
	{123, "EPFNOSUPPORT", "protocol family not supported"},
	{124, "EAFNOSUPPORT", "address family not supported by protocol"},
	{125, "EADDRINUSE", "address already in use"},
	{126, "EADDRNOTAVAIL", "cannot assign requested address"},
	{127, "ENETDOWN", "network is down"},
	{128, "ENETUNREACH", "network is unreachable"},
	{129, "ENETRESET", "network dropped connection on reset"},
	{130, "ECONNABORTED", "software caused connection abort"},
	{131, "ECONNRESET", "connection reset by peer"},
	{132, "ENOBUFS", "no buffer space available"},
	{133, "EISCONN", "transport endpoint is already connected"},
	{134, "ENOTCONN", "transport endpoint is not connected"},
	{135, "EUCLEAN", "structure needs cleaning"},
	{137, "ENOTNAM", "not a XENIX named type file"},
	{138, "ENAVAIL", "no XENIX semaphores available"},
	{139, "EISNAM", "is a named type file"},
	{140, "EREMOTEIO", "remote I/O error"},
	{141, "EINIT", "unknown error 141"},
	{142, "EREMDEV", "unknown error 142"},
	{143, "ESHUTDOWN", "cannot send after transport endpoint shutdown"},
	{144, "ETOOMANYREFS", "too many references: cannot splice"},
	{145, "ETIMEDOUT", "connection timed out"},
	{146, "ECONNREFUSED", "connection refused"},
	{147, "EHOSTDOWN", "host is down"},
	{148, "EHOSTUNREACH", "no route to host"},
	{149, "EALREADY", "operation already in progress"},
	{150, "EINPROGRESS", "operation now in progress"},
	{151, "ESTALE", "stale file handle"},
	{158, "ECANCELED", "operation canceled"},
	{159, "ENOMEDIUM", "no medium found"},
	{160, "EMEDIUMTYPE", "wrong medium type"},
	{161, "ENOKEY", "required key not available"},
	{162, "EKEYEXPIRED", "key has expired"},
	{163, "EKEYREVOKED", "key has been revoked"},
	{164, "EKEYREJECTED", "key was rejected by service"},
	{165, "EOWNERDEAD", "owner died"},
	{166, "ENOTRECOVERABLE", "state not recoverable"},
	{167, "ERFKILL", "operation not possible due to RF-kill"},
	{168, "EHWPOISON", "memory page has hardware error"},
	{1133, "EDQUOT", "disk quota exceeded"},
}

// Signal table
var signalList = [...]struct {
	num  syscall.Signal
	name string
	desc string
}{
	{1, "SIGHUP", "hangup"},
	{2, "SIGINT", "interrupt"},
	{3, "SIGQUIT", "quit"},
	{4, "SIGILL", "illegal instruction"},
	{5, "SIGTRAP", "trace/breakpoint trap"},
	{6, "SIGABRT", "aborted"},
	{7, "SIGEMT", "EMT trap"},
	{8, "SIGFPE", "floating point exception"},
	{9, "SIGKILL", "killed"},
	{10, "SIGBUS", "bus error"},
	{11, "SIGSEGV", "segmentation fault"},
	{12, "SIGSYS", "bad system call"},
	{13, "SIGPIPE", "broken pipe"},
	{14, "SIGALRM", "alarm clock"},
	{15, "SIGTERM", "terminated"},
	{16, "SIGUSR1", "user defined signal 1"},
	{17, "SIGUSR2", "user defined signal 2"},
	{18, "SIGCHLD", "child exited"},
	{19, "SIGPWR", "power failure"},
	{20, "SIGWINCH", "window changed"},
	{21, "SIGURG", "urgent I/O condition"},
	{22, "SIGIO", "I/O possible"},
	{23, "SIGSTOP", "stopped (signal)"},
	{24, "SIGTSTP", "stopped"},
	{25, "SIGCONT", "continued"},
	{26, "SIGTTIN", "stopped (tty input)"},
	{27, "SIGTTOU", "stopped (tty output)"},
	{28, "SIGVTALRM", "virtual timer expired"},
	{29, "SIGPROF", "profiling timer expired"},
	{30, "SIGXCPU", "CPU time limit exceeded"},
	{31, "SIGXFSZ", "file size limit exceeded"},
}
