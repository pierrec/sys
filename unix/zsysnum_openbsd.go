// Generated code. DO NOT EDIT.

package unix

const (
	SYS_EXIT           = 1
	SYS_FORK           = 2
	SYS_READ           = 3
	SYS_WRITE          = 4
	SYS_OPEN           = 5
	SYS_CLOSE          = 6
	SYS_GETENTROPY     = 7
	SYS___TFORK        = 8
	SYS_LINK           = 9
	SYS_UNLINK         = 10
	SYS_WAIT4          = 11
	SYS_CHDIR          = 12
	SYS_FCHDIR         = 13
	SYS_MKNOD          = 14
	SYS_CHMOD          = 15
	SYS_CHOWN          = 16
	SYS_OBREAK         = 17
	SYS_GETDTABLECOUNT = 18
	SYS_GETRUSAGE      = 19
	SYS_GETPID         = 20
	SYS_MOUNT          = 21
	SYS_UNMOUNT        = 22
	SYS_SETUID         = 23
	SYS_GETUID         = 24
	SYS_GETEUID        = 25
	SYS_PTRACE         = 26
	SYS_RECVMSG        = 27
	SYS_SENDMSG        = 28
	SYS_RECVFROM       = 29
	SYS_ACCEPT         = 30
	SYS_GETPEERNAME    = 31
	SYS_GETSOCKNAME    = 32
	SYS_ACCESS         = 33
	SYS_CHFLAGS        = 34
	SYS_FCHFLAGS       = 35
	SYS_SYNC           = 36
	SYS_STAT           = 38
	SYS_GETPPID        = 39
	SYS_LSTAT          = 40
	SYS_DUP            = 41
	SYS_FSTATAT        = 42
	SYS_GETEGID        = 43
	SYS_PROFIL         = 44
	SYS_KTRACE         = 45
	SYS_SIGACTION      = 46
	SYS_GETGID         = 47
	SYS_SIGPROCMASK    = 48
	SYS_SETLOGIN       = 50
	SYS_ACCT           = 51
	SYS_SIGPENDING     = 52
	SYS_FSTAT          = 53
	SYS_IOCTL          = 54
	SYS_REBOOT         = 55
	SYS_REVOKE         = 56
	SYS_SYMLINK        = 57
	SYS_READLINK       = 58
	SYS_EXECVE         = 59
	SYS_UMASK          = 60
	SYS_CHROOT         = 61
	SYS_GETFSSTAT      = 62
	SYS_STATFS         = 63
	SYS_FSTATFS        = 64
	SYS_FHSTATFS       = 65
	SYS_VFORK          = 66
	SYS_GETTIMEOFDAY   = 67
	SYS_SETTIMEOFDAY   = 68
	SYS_SETITIMER      = 69
	SYS_GETITIMER      = 70
	SYS_SELECT         = 71
	SYS_KEVENT         = 72
	SYS_MUNMAP         = 73
	SYS_MPROTECT       = 74
	SYS_MADVISE        = 75
	SYS_UTIMES         = 76
	SYS_FUTIMES        = 77
	SYS_GETGROUPS      = 79
	SYS_SETGROUPS      = 80
	SYS_GETPGRP        = 81
	SYS_SETPGID        = 82
	SYS_FUTEX          = 83
	SYS_UTIMENSAT      = 84
	SYS_FUTIMENS       = 85
	SYS_KBIND          = 86
	SYS_CLOCK_GETTIME  = 87
	SYS_CLOCK_SETTIME  = 88
	SYS_CLOCK_GETRES   = 89
	SYS_DUP2           = 90
	SYS_NANOSLEEP      = 91
	SYS_FCNTL          = 92
	SYS_ACCEPT4        = 93
	SYS___THRSLEEP     = 94
	SYS_FSYNC          = 95
	SYS_SETPRIORITY    = 96
	SYS_SOCKET         = 97
	SYS_CONNECT        = 98
	SYS_GETDENTS       = 99
	SYS_GETPRIORITY    = 100
	SYS_PIPE2          = 101
	SYS_DUP3           = 102
	SYS_SIGRETURN      = 103
	SYS_BIND           = 104
	SYS_SETSOCKOPT     = 105
	SYS_LISTEN         = 106
	SYS_CHFLAGSAT      = 107
	SYS_PLEDGE         = 108
	SYS_PPOLL          = 109
	SYS_PSELECT        = 110
	SYS_SIGSUSPEND     = 111
	SYS_SENDSYSLOG     = 112
	SYS_UNVEIL         = 114
	SYS_GETSOCKOPT     = 118
	SYS_THRKILL        = 119
	SYS_READV          = 120
	SYS_WRITEV         = 121
	SYS_KILL           = 122
	SYS_FCHOWN         = 123
	SYS_FCHMOD         = 124
	SYS_SETREUID       = 126
	SYS_SETREGID       = 127
	SYS_RENAME         = 128
	SYS_FLOCK          = 131
	SYS_MKFIFO         = 132
	SYS_SENDTO         = 133
	SYS_SHUTDOWN       = 134
	SYS_SOCKETPAIR     = 135
	SYS_MKDIR          = 136
	SYS_RMDIR          = 137
	SYS_ADJTIME        = 140
	SYS_GETLOGIN_R     = 141
	SYS_SETSID         = 147
	SYS_QUOTACTL       = 148
	SYS_NFSSVC         = 155
	SYS_GETFH          = 161
	SYS_SYSARCH        = 165
	SYS_PREAD          = 173
	SYS_PWRITE         = 174
	SYS_SETGID         = 181
	SYS_SETEGID        = 182
	SYS_SETEUID        = 183
	SYS_PATHCONF       = 191
	SYS_FPATHCONF      = 192
	SYS_SWAPCTL        = 193
	SYS_GETRLIMIT      = 194
	SYS_SETRLIMIT      = 195
	SYS_MMAP           = 197
	SYS_LSEEK          = 199
	SYS_TRUNCATE       = 200
	SYS_FTRUNCATE      = 201
	SYS_SYSCTL         = 202
	SYS_MLOCK          = 203
	SYS_MUNLOCK        = 204
	SYS_GETPGID        = 207
	SYS_UTRACE         = 209
	SYS_SEMGET         = 221
	SYS_MSGGET         = 225
	SYS_MSGSND         = 226
	SYS_MSGRCV         = 227
	SYS_SHMAT          = 228
	SYS_SHMDT          = 230
	SYS_MINHERIT       = 250
	SYS_POLL           = 252
	SYS_ISSETUGID      = 253
	SYS_LCHOWN         = 254
	SYS_GETSID         = 255
	SYS_MSYNC          = 256
	SYS_PIPE           = 263
	SYS_FHOPEN         = 264
	SYS_PREADV         = 267
	SYS_PWRITEV        = 268
	SYS_KQUEUE         = 269
	SYS_MLOCKALL       = 271
	SYS_MUNLOCKALL     = 272
	SYS_GETRESUID      = 281
	SYS_SETRESUID      = 282
	SYS_GETRESGID      = 283
	SYS_SETRESGID      = 284
	SYS_MQUERY         = 286
	SYS_CLOSEFROM      = 287
	SYS_SIGALTSTACK    = 288
	SYS_SHMGET         = 289
	SYS_SEMOP          = 290
	SYS_FHSTAT         = 294
	SYS___SEMCTL       = 295
	SYS_SHMCTL         = 296
	SYS_MSGCTL         = 297
	SYS_SCHED_YIELD    = 298
	SYS_GETTHRID       = 299
	SYS___THRWAKEUP    = 301
	SYS___THREXIT      = 302
	SYS___THRSIGDIVERT = 303
	SYS___GETCWD       = 304
	SYS_ADJFREQ        = 305
	SYS_SETRTABLE      = 310
	SYS_GETRTABLE      = 311
	SYS_FACCESSAT      = 313
	SYS_FCHMODAT       = 314
	SYS_FCHOWNAT       = 315
	SYS_LINKAT         = 317
	SYS_MKDIRAT        = 318
	SYS_MKFIFOAT       = 319
	SYS_MKNODAT        = 320
	SYS_OPENAT         = 321
	SYS_READLINKAT     = 322
	SYS_RENAMEAT       = 323
	SYS_SYMLINKAT      = 324
	SYS_UNLINKAT       = 325
	SYS___SET_TCB      = 329
	SYS___GET_TCB      = 330
)
