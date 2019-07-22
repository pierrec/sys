// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build sparc64,linux

package unix

const (
	SYS_FORK = 2

	SYS_OPEN = 5

	SYS_CREAT  = 8
	SYS_LINK   = 9
	SYS_UNLINK = 10
	SYS_EXECV  = 11

	SYS_CHOWN  = 13
	SYS_MKNOD  = 14
	SYS_CHMOD  = 15
	SYS_LCHOWN = 16

	SYS_PERFCTR = 18

	SYS_ALARM = 27

	SYS_PAUSE  = 29
	SYS_UTIME  = 30
	SYS_ACCESS = 33
	SYS_NICE   = 34

	SYS_STAT = 38

	SYS_LSTAT = 40

	SYS_PIPE = 42

	SYS_SIGNAL = 48

	SYS_MEMORY_ORDERING = 52

	SYS_SYMLINK  = 57
	SYS_READLINK = 58

	SYS_FSTAT64     = 63
	SYS_GETPAGESIZE = 64

	SYS_VFORK = 66

	SYS_MMAP = 71

	SYS_GETPGRP = 81

	SYS_DUP2 = 90

	SYS_SELECT = 93

	SYS_ACCEPT = 99

	SYS_RENAME = 128

	SYS_LSTAT64 = 132

	SYS_MKDIR      = 136
	SYS_RMDIR      = 137
	SYS_UTIMES     = 138
	SYS_STAT64     = 139
	SYS_SENDFILE64 = 140

	SYS_GETRLIMIT = 144

	SYS_PCICONFIG_READ  = 148
	SYS_PCICONFIG_WRITE = 149

	SYS_INOTIFY_INIT = 151

	SYS_POLL = 153

	SYS_UMOUNT             = 159
	SYS_SCHED_SET_AFFINITY = 160
	SYS_SCHED_GET_AFFINITY = 161
	SYS_GETDOMAINNAME      = 162

	SYS_UTRAP_INSTALL = 164

	SYS_USTAT = 168

	SYS_GETDENTS = 174

	SYS_SIGPENDING   = 183
	SYS_QUERY_MODULE = 184

	SYS_EPOLL_CREATE = 193

	SYS_EPOLL_WAIT = 195

	SYS_SIGACTION  = 198
	SYS_SGETMASK   = 199
	SYS_SSETMASK   = 200
	SYS_SIGSUSPEND = 201
	SYS_OLDLSTAT   = 202
	SYS_USELIB     = 203
	SYS_READDIR    = 204

	SYS_SOCKETCALL = 206

	SYS_FADVISE64    = 209
	SYS_FADVISE64_64 = 210

	SYS_WAITPID = 212

	SYS_IPC       = 215
	SYS_SIGRETURN = 216

	SYS_SIGPROCMASK   = 220
	SYS_CREATE_MODULE = 221

	SYS_GET_KERNEL_SYMS = 223

	SYS_BDFLUSH     = 225
	SYS_SYSFS       = 226
	SYS_AFS_SYSCALL = 227

	SYS__NEWSELECT = 230

	SYS_STIME     = 233
	SYS_STATFS64  = 234
	SYS_FSTATFS64 = 235
	SYS__LLSEEK   = 236

	SYS__SYSCTL = 251

	SYS_SYNC_FILE_RANGE = 255

	SYS_VSERVER = 267

	SYS_FUTIMESAT = 288
	SYS_FSTATAT64 = 289

	SYS_RENAMEAT = 291

	SYS_SIGNALFD = 311

	SYS_EVENTFD = 313

	SYS_KERN_FEATURES = 340

	SYS_SEMTIMEDOP = 392
)
