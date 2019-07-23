// mkerrors.sh -Wall -Werror -static -I/tmp/include -m32
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,linux

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Wall -Werror -static -I/tmp/include -m32 _const.go

package unix

import "syscall"

const (
	FP_XSTATE_MAGIC2 = 0x46505845

	MAP_32BIT = 0x40

	MAP_SYNC = 0x80000

	PTRACE_GETFPREGS  = 0xe
	PTRACE_GETFPXREGS = 0x12

	PTRACE_GET_THREAD_AREA = 0x19

	PTRACE_OLDSETOPTIONS = 0x15

	PTRACE_SETFPREGS  = 0xf
	PTRACE_SETFPXREGS = 0x13

	PTRACE_SET_THREAD_AREA = 0x1a
	PTRACE_SINGLEBLOCK     = 0x21

	PTRACE_SYSEMU            = 0x1f
	PTRACE_SYSEMU_SINGLESTEP = 0x20

	TCGETS2 = 0x802c542a
	TCGETX  = 0x5432

	TCSETS2 = 0x402c542b

	TCSETSF2 = 0x402c542d

	TCSETSW2 = 0x402c542c
	TCSETX   = 0x5433
	TCSETXF  = 0x5434
	TCSETXW  = 0x5435

	TIOCSER_TEMT = 0x1

	X86_FXSR_MAGIC = 0x0
)

// Errors

// Signals
const (
	SIGSTKFLT = syscall.Signal(0x10)
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
	{35, "EDEADLK", "resource deadlock avoided"},
	{36, "ENAMETOOLONG", "file name too long"},
	{37, "ENOLCK", "no locks available"},
	{38, "ENOSYS", "function not implemented"},
	{39, "ENOTEMPTY", "directory not empty"},
	{40, "ELOOP", "too many levels of symbolic links"},
	{42, "ENOMSG", "no message of desired type"},
	{43, "EIDRM", "identifier removed"},
	{44, "ECHRNG", "channel number out of range"},
	{45, "EL2NSYNC", "level 2 not synchronized"},
	{46, "EL3HLT", "level 3 halted"},
	{47, "EL3RST", "level 3 reset"},
	{48, "ELNRNG", "link number out of range"},
	{49, "EUNATCH", "protocol driver not attached"},
	{50, "ENOCSI", "no CSI structure available"},
	{51, "EL2HLT", "level 2 halted"},
	{52, "EBADE", "invalid exchange"},
	{53, "EBADR", "invalid request descriptor"},
	{54, "EXFULL", "exchange full"},
	{55, "ENOANO", "no anode"},
	{56, "EBADRQC", "invalid request code"},
	{57, "EBADSLT", "invalid slot"},
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
	{72, "EMULTIHOP", "multihop attempted"},
	{73, "EDOTDOT", "RFS specific error"},
	{74, "EBADMSG", "bad message"},
	{75, "EOVERFLOW", "value too large for defined data type"},
	{76, "ENOTUNIQ", "name not unique on network"},
	{77, "EBADFD", "file descriptor in bad state"},
	{78, "EREMCHG", "remote address changed"},
	{79, "ELIBACC", "can not access a needed shared library"},
	{80, "ELIBBAD", "accessing a corrupted shared library"},
	{81, "ELIBSCN", ".lib section in a.out corrupted"},
	{82, "ELIBMAX", "attempting to link in too many shared libraries"},
	{83, "ELIBEXEC", "cannot exec a shared library directly"},
	{84, "EILSEQ", "invalid or incomplete multibyte or wide character"},
	{85, "ERESTART", "interrupted system call should be restarted"},
	{86, "ESTRPIPE", "streams pipe error"},
	{87, "EUSERS", "too many users"},
	{88, "ENOTSOCK", "socket operation on non-socket"},
	{89, "EDESTADDRREQ", "destination address required"},
	{90, "EMSGSIZE", "message too long"},
	{91, "EPROTOTYPE", "protocol wrong type for socket"},
	{92, "ENOPROTOOPT", "protocol not available"},
	{93, "EPROTONOSUPPORT", "protocol not supported"},
	{94, "ESOCKTNOSUPPORT", "socket type not supported"},
	{95, "ENOTSUP", "operation not supported"},
	{96, "EPFNOSUPPORT", "protocol family not supported"},
	{97, "EAFNOSUPPORT", "address family not supported by protocol"},
	{98, "EADDRINUSE", "address already in use"},
	{99, "EADDRNOTAVAIL", "cannot assign requested address"},
	{100, "ENETDOWN", "network is down"},
	{101, "ENETUNREACH", "network is unreachable"},
	{102, "ENETRESET", "network dropped connection on reset"},
	{103, "ECONNABORTED", "software caused connection abort"},
	{104, "ECONNRESET", "connection reset by peer"},
	{105, "ENOBUFS", "no buffer space available"},
	{106, "EISCONN", "transport endpoint is already connected"},
	{107, "ENOTCONN", "transport endpoint is not connected"},
	{108, "ESHUTDOWN", "cannot send after transport endpoint shutdown"},
	{109, "ETOOMANYREFS", "too many references: cannot splice"},
	{110, "ETIMEDOUT", "connection timed out"},
	{111, "ECONNREFUSED", "connection refused"},
	{112, "EHOSTDOWN", "host is down"},
	{113, "EHOSTUNREACH", "no route to host"},
	{114, "EALREADY", "operation already in progress"},
	{115, "EINPROGRESS", "operation now in progress"},
	{116, "ESTALE", "stale file handle"},
	{117, "EUCLEAN", "structure needs cleaning"},
	{118, "ENOTNAM", "not a XENIX named type file"},
	{119, "ENAVAIL", "no XENIX semaphores available"},
	{120, "EISNAM", "is a named type file"},
	{121, "EREMOTEIO", "remote I/O error"},
	{122, "EDQUOT", "disk quota exceeded"},
	{123, "ENOMEDIUM", "no medium found"},
	{124, "EMEDIUMTYPE", "wrong medium type"},
	{125, "ECANCELED", "operation canceled"},
	{126, "ENOKEY", "required key not available"},
	{127, "EKEYEXPIRED", "key has expired"},
	{128, "EKEYREVOKED", "key has been revoked"},
	{129, "EKEYREJECTED", "key was rejected by service"},
	{130, "EOWNERDEAD", "owner died"},
	{131, "ENOTRECOVERABLE", "state not recoverable"},
	{132, "ERFKILL", "operation not possible due to RF-kill"},
	{133, "EHWPOISON", "memory page has hardware error"},
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
	{7, "SIGBUS", "bus error"},
	{8, "SIGFPE", "floating point exception"},
	{9, "SIGKILL", "killed"},
	{10, "SIGUSR1", "user defined signal 1"},
	{11, "SIGSEGV", "segmentation fault"},
	{12, "SIGUSR2", "user defined signal 2"},
	{13, "SIGPIPE", "broken pipe"},
	{14, "SIGALRM", "alarm clock"},
	{15, "SIGTERM", "terminated"},
	{16, "SIGSTKFLT", "stack fault"},
	{17, "SIGCHLD", "child exited"},
	{18, "SIGCONT", "continued"},
	{19, "SIGSTOP", "stopped (signal)"},
	{20, "SIGTSTP", "stopped"},
	{21, "SIGTTIN", "stopped (tty input)"},
	{22, "SIGTTOU", "stopped (tty output)"},
	{23, "SIGURG", "urgent I/O condition"},
	{24, "SIGXCPU", "CPU time limit exceeded"},
	{25, "SIGXFSZ", "file size limit exceeded"},
	{26, "SIGVTALRM", "virtual timer expired"},
	{27, "SIGPROF", "profiling timer expired"},
	{28, "SIGWINCH", "window changed"},
	{29, "SIGIO", "I/O possible"},
	{30, "SIGPWR", "power failure"},
	{31, "SIGSYS", "bad system call"},
}
