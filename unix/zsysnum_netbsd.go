// Generated code. DO NOT EDIT.

package unix

const (
	SYS_EXIT                 = 1
	SYS_FORK                 = 2
	SYS_READ                 = 3
	SYS_WRITE                = 4
	SYS_OPEN                 = 5
	SYS_CLOSE                = 6
	SYS_LINK                 = 9
	SYS_UNLINK               = 10
	SYS_CHDIR                = 12
	SYS_FCHDIR               = 13
	SYS_CHMOD                = 15
	SYS_CHOWN                = 16
	SYS_BREAK                = 17
	SYS_GETPID               = 20
	SYS_UNMOUNT              = 22
	SYS_SETUID               = 23
	SYS_GETUID               = 24
	SYS_GETEUID              = 25
	SYS_PTRACE               = 26
	SYS_RECVMSG              = 27
	SYS_SENDMSG              = 28
	SYS_RECVFROM             = 29
	SYS_ACCEPT               = 30
	SYS_GETPEERNAME          = 31
	SYS_GETSOCKNAME          = 32
	SYS_ACCESS               = 33
	SYS_CHFLAGS              = 34
	SYS_FCHFLAGS             = 35
	SYS_SYNC                 = 36
	SYS_KILL                 = 37
	SYS_GETPPID              = 39
	SYS_DUP                  = 41
	SYS_PIPE                 = 42
	SYS_GETEGID              = 43
	SYS_PROFIL               = 44
	SYS_KTRACE               = 45
	SYS_GETGID               = 47
	SYS___GETLOGIN           = 49
	SYS___SETLOGIN           = 50
	SYS_ACCT                 = 51
	SYS_IOCTL                = 54
	SYS_REVOKE               = 56
	SYS_SYMLINK              = 57
	SYS_READLINK             = 58
	SYS_EXECVE               = 59
	SYS_UMASK                = 60
	SYS_CHROOT               = 61
	SYS_VFORK                = 66
	SYS_SBRK                 = 69
	SYS_SSTK                 = 70
	SYS_VADVISE              = 72
	SYS_MUNMAP               = 73
	SYS_MPROTECT             = 74
	SYS_MADVISE              = 75
	SYS_MINCORE              = 78
	SYS_GETGROUPS            = 79
	SYS_SETGROUPS            = 80
	SYS_GETPGRP              = 81
	SYS_SETPGID              = 82
	SYS_DUP2                 = 90
	SYS_FCNTL                = 92
	SYS_FSYNC                = 95
	SYS_SETPRIORITY          = 96
	SYS_CONNECT              = 98
	SYS_GETPRIORITY          = 100
	SYS_BIND                 = 104
	SYS_SETSOCKOPT           = 105
	SYS_LISTEN               = 106
	SYS_GETSOCKOPT           = 118
	SYS_READV                = 120
	SYS_WRITEV               = 121
	SYS_FCHOWN               = 123
	SYS_FCHMOD               = 124
	SYS_SETREUID             = 126
	SYS_SETREGID             = 127
	SYS_RENAME               = 128
	SYS_FLOCK                = 131
	SYS_MKFIFO               = 132
	SYS_SENDTO               = 133
	SYS_SHUTDOWN             = 134
	SYS_SOCKETPAIR           = 135
	SYS_MKDIR                = 136
	SYS_RMDIR                = 137
	SYS_SETSID               = 147
	SYS_SYSARCH              = 165
	SYS_PREAD                = 173
	SYS_PWRITE               = 174
	SYS_NTP_ADJTIME          = 176
	SYS_SETGID               = 181
	SYS_SETEGID              = 182
	SYS_SETEUID              = 183
	SYS_PATHCONF             = 191
	SYS_FPATHCONF            = 192
	SYS_GETRLIMIT            = 194
	SYS_SETRLIMIT            = 195
	SYS_MMAP                 = 197
	SYS_LSEEK                = 199
	SYS_TRUNCATE             = 200
	SYS_FTRUNCATE            = 201
	SYS___SYSCTL             = 202
	SYS_MLOCK                = 203
	SYS_MUNLOCK              = 204
	SYS_UNDELETE             = 205
	SYS_GETPGID              = 207
	SYS_REBOOT               = 208
	SYS_POLL                 = 209
	SYS_SEMGET               = 221
	SYS_SEMOP                = 222
	SYS_SEMCONFIG            = 223
	SYS_MSGGET               = 225
	SYS_MSGSND               = 226
	SYS_MSGRCV               = 227
	SYS_SHMAT                = 228
	SYS_SHMDT                = 230
	SYS_SHMGET               = 231
	SYS_TIMER_CREATE         = 235
	SYS_TIMER_DELETE         = 236
	SYS_TIMER_GETOVERRUN     = 239
	SYS_FDATASYNC            = 241
	SYS_MLOCKALL             = 242
	SYS_MUNLOCKALL           = 243
	SYS_SIGQUEUEINFO         = 245
	SYS_MODCTL               = 246
	SYS___POSIX_RENAME       = 270
	SYS_SWAPCTL              = 271
	SYS_MINHERIT             = 273
	SYS_LCHMOD               = 274
	SYS_LCHOWN               = 275
	SYS_MSYNC                = 277
	SYS___POSIX_CHOWN        = 283
	SYS___POSIX_FCHOWN       = 284
	SYS___POSIX_LCHOWN       = 285
	SYS_GETSID               = 286
	SYS___CLONE              = 287
	SYS_FKTRACE              = 288
	SYS_PREADV               = 289
	SYS_PWRITEV              = 290
	SYS___GETCWD             = 296
	SYS_FCHROOT              = 297
	SYS_LCHFLAGS             = 304
	SYS_ISSETUGID            = 305
	SYS_UTRACE               = 306
	SYS_GETCONTEXT           = 307
	SYS_SETCONTEXT           = 308
	SYS__LWP_CREATE          = 309
	SYS__LWP_EXIT            = 310
	SYS__LWP_SELF            = 311
	SYS__LWP_WAIT            = 312
	SYS__LWP_SUSPEND         = 313
	SYS__LWP_CONTINUE        = 314
	SYS__LWP_WAKEUP          = 315
	SYS__LWP_GETPRIVATE      = 316
	SYS__LWP_SETPRIVATE      = 317
	SYS__LWP_KILL            = 318
	SYS__LWP_DETACH          = 319
	SYS__LWP_UNPARK          = 321
	SYS__LWP_UNPARK_ALL      = 322
	SYS__LWP_SETNAME         = 323
	SYS__LWP_GETNAME         = 324
	SYS__LWP_CTL             = 325
	SYS___SIGACTION_SIGTRAMP = 340
	SYS_PMC_GET_INFO         = 341
	SYS_PMC_CONTROL          = 342
	SYS_RASCTL               = 343
	SYS_KQUEUE               = 344
	SYS__SCHED_SETPARAM      = 346
	SYS__SCHED_GETPARAM      = 347
	SYS__SCHED_SETAFFINITY   = 348
	SYS__SCHED_GETAFFINITY   = 349
	SYS_SCHED_YIELD          = 350
	SYS_FSYNC_RANGE          = 354
	SYS_UUIDGEN              = 355
	SYS_GETVFSSTAT           = 356
	SYS_STATVFS1             = 357
	SYS_FSTATVFS1            = 358
	SYS_EXTATTRCTL           = 360
	SYS_EXTATTR_SET_FILE     = 361
	SYS_EXTATTR_GET_FILE     = 362
	SYS_EXTATTR_DELETE_FILE  = 363
	SYS_EXTATTR_SET_FD       = 364
	SYS_EXTATTR_GET_FD       = 365
	SYS_EXTATTR_DELETE_FD    = 366
	SYS_EXTATTR_SET_LINK     = 367
	SYS_EXTATTR_GET_LINK     = 368
	SYS_EXTATTR_DELETE_LINK  = 369
	SYS_EXTATTR_LIST_FD      = 370
	SYS_EXTATTR_LIST_FILE    = 371
	SYS_EXTATTR_LIST_LINK    = 372
	SYS_SETXATTR             = 375
	SYS_LSETXATTR            = 376
	SYS_FSETXATTR            = 377
	SYS_GETXATTR             = 378
	SYS_LGETXATTR            = 379
	SYS_FGETXATTR            = 380
	SYS_LISTXATTR            = 381
	SYS_LLISTXATTR           = 382
	SYS_FLISTXATTR           = 383
	SYS_REMOVEXATTR          = 384
	SYS_LREMOVEXATTR         = 385
	SYS_FREMOVEXATTR         = 386
	SYS_GETDENTS             = 390
	SYS_SOCKET               = 394
	SYS_GETFH                = 395
	SYS_MOUNT                = 410
	SYS_MREMAP               = 411
	SYS_PSET_CREATE          = 412
	SYS_PSET_DESTROY         = 413
	SYS_PSET_ASSIGN          = 414
	SYS__PSET_BIND           = 415
	SYS_POSIX_FADVISE        = 416
	SYS_SELECT               = 417
	SYS_GETTIMEOFDAY         = 418
	SYS_SETTIMEOFDAY         = 419
	SYS_UTIMES               = 420
	SYS_ADJTIME              = 421
	SYS_FUTIMES              = 423
	SYS_LUTIMES              = 424
	SYS_SETITIMER            = 425
	SYS_GETITIMER            = 426
	SYS_CLOCK_GETTIME        = 427
	SYS_CLOCK_SETTIME        = 428
	SYS_CLOCK_GETRES         = 429
	SYS_NANOSLEEP            = 430
	SYS___SIGTIMEDWAIT       = 431
	SYS__LWP_PARK            = 434
	SYS_KEVENT               = 435
	SYS_PSELECT              = 436
	SYS_POLLTS               = 437
	SYS_STAT                 = 439
	SYS_FSTAT                = 440
	SYS_LSTAT                = 441
	SYS___SEMCTL             = 442
	SYS_SHMCTL               = 443
	SYS_MSGCTL               = 444
	SYS_GETRUSAGE            = 445
	SYS_TIMER_SETTIME        = 446
	SYS_TIMER_GETTIME        = 447
	SYS_NTP_GETTIME          = 448
	SYS_WAIT4                = 449
	SYS_MKNOD                = 450
	SYS_FHSTAT               = 451
	SYS_PIPE2                = 453
	SYS_DUP3                 = 454
	SYS_KQUEUE1              = 455
	SYS_PACCEPT              = 456
	SYS_LINKAT               = 457
	SYS_RENAMEAT             = 458
	SYS_MKFIFOAT             = 459
	SYS_MKNODAT              = 460
	SYS_MKDIRAT              = 461
	SYS_FACCESSAT            = 462
	SYS_FCHMODAT             = 463
	SYS_FCHOWNAT             = 464
	SYS_FEXECVE              = 465
	SYS_FSTATAT              = 466
	SYS_UTIMENSAT            = 467
	SYS_OPENAT               = 468
	SYS_READLINKAT           = 469
	SYS_SYMLINKAT            = 470
	SYS_UNLINKAT             = 471
	SYS_FUTIMENS             = 472
	SYS___QUOTACTL           = 473
	SYS_POSIX_SPAWN          = 474
	SYS_RECVMMSG             = 475
	SYS_SENDMMSG             = 476
)
