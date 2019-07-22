// go run mksysnum.go https://svn.freebsd.org/base/stable/11/sys/kern/syscalls.master
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build 386,freebsd

package unix

// SYS_NOSYS = 0;  // { int nosys(void); } syscall nosys_args int
// { void sys_exit(int rval); } exit sys_exit_args void
// { int fork(void); }
// { ssize_t read(int fd, void *buf, size_t nbyte); }
// { ssize_t write(int fd, const void *buf, size_t nbyte); }
// { int open(char *path, int flags, int mode); }
// { int close(int fd); }
// { int wait4(int pid, int *status, int options, struct rusage *rusage); }
// { int link(char *path, char *link); }
// { int unlink(char *path); }
// { int chdir(char *path); }
// { int fchdir(int fd); }
// { int mknod(char *path, int mode, int dev); }
// { int chmod(char *path, int mode); }
// { int chown(char *path, int uid, int gid); }
// { int obreak(char *nsize); } break obreak_args int
// { pid_t getpid(void); }
// { int mount(char *type, char *path, int flags, caddr_t data); }
// { int unmount(char *path, int flags); }
// { int setuid(uid_t uid); }
// { uid_t getuid(void); }
// { uid_t geteuid(void); }
// { int ptrace(int req, pid_t pid, caddr_t addr, int data); }
// { int recvmsg(int s, struct msghdr *msg, int flags); }
// { int sendmsg(int s, struct msghdr *msg, int flags); }
// { int recvfrom(int s, caddr_t buf, size_t len, int flags, struct sockaddr * __restrict from, __socklen_t * __restrict fromlenaddr); }
// { int accept(int s, struct sockaddr * __restrict name, __socklen_t * __restrict anamelen); }
// { int getpeername(int fdes, struct sockaddr * __restrict asa, __socklen_t * __restrict alen); }
// { int getsockname(int fdes, struct sockaddr * __restrict asa, __socklen_t * __restrict alen); }
// { int access(char *path, int amode); }
// { int chflags(const char *path, u_long flags); }
// { int fchflags(int fd, u_long flags); }
// { int sync(void); }
// { int kill(int pid, int signum); }
// { pid_t getppid(void); }
// { int dup(u_int fd); }
// { int pipe(void); }
// { gid_t getegid(void); }
// { int profil(caddr_t samples, size_t size, size_t offset, u_int scale); }
// { int ktrace(const char *fname, int ops, int facs, int pid); }
// { gid_t getgid(void); }
// { int getlogin(char *namebuf, u_int namelen); }
// { int setlogin(char *namebuf); }
// { int acct(char *path); }
// { int sigaltstack(stack_t *ss, stack_t *oss); }
// { int ioctl(int fd, u_long com, caddr_t data); }
// { int reboot(int opt); }
// { int revoke(char *path); }
// { int symlink(char *path, char *link); }
// { ssize_t readlink(char *path, char *buf, size_t count); }
// { int execve(char *fname, char **argv, char **envv); }
// { int umask(int newmask); } umask umask_args int
// { int chroot(char *path); }
// { int msync(void *addr, size_t len, int flags); }
// { int vfork(void); }
// { int sbrk(int incr); }
// { int sstk(int incr); }
// { int ovadvise(int anom); } vadvise ovadvise_args int
// { int munmap(void *addr, size_t len); }
// { int mprotect(const void *addr, size_t len, int prot); }
// { int madvise(void *addr, size_t len, int behav); }
// { int mincore(const void *addr, size_t len, char *vec); }
// { int getgroups(u_int gidsetsize, gid_t *gidset); }
// { int setgroups(u_int gidsetsize, gid_t *gidset); }
// { int getpgrp(void); }
// { int setpgid(int pid, int pgid); }
// { int setitimer(u_int which, struct itimerval *itv, struct itimerval *oitv); }
// { int swapon(char *name); }
// { int getitimer(u_int which, struct itimerval *itv); }
// { int getdtablesize(void); }
// { int dup2(u_int from, u_int to); }
// { int fcntl(int fd, int cmd, long arg); }
// { int select(int nd, fd_set *in, fd_set *ou, fd_set *ex, struct timeval *tv); }
// { int fsync(int fd); }
// { int setpriority(int which, int who, int prio); }
// { int socket(int domain, int type, int protocol); }
// { int connect(int s, caddr_t name, int namelen); }
// { int getpriority(int which, int who); }
// { int bind(int s, caddr_t name, int namelen); }
// { int setsockopt(int s, int level, int name, caddr_t val, int valsize); }
// { int listen(int s, int backlog); }
// { int gettimeofday(struct timeval *tp, struct timezone *tzp); }
// { int getrusage(int who, struct rusage *rusage); }
// { int getsockopt(int s, int level, int name, caddr_t val, int *avalsize); }
// { int readv(int fd, struct iovec *iovp, u_int iovcnt); }
// { int writev(int fd, struct iovec *iovp, u_int iovcnt); }
// { int settimeofday(struct timeval *tv, struct timezone *tzp); }
// { int fchown(int fd, int uid, int gid); }
// { int fchmod(int fd, int mode); }
// { int setreuid(int ruid, int euid); }
// { int setregid(int rgid, int egid); }
// { int rename(char *from, char *to); }
// { int flock(int fd, int how); }
// { int mkfifo(char *path, int mode); }
// { int sendto(int s, caddr_t buf, size_t len, int flags, caddr_t to, int tolen); }
// { int shutdown(int s, int how); }
// { int socketpair(int domain, int type, int protocol, int *rsv); }
// { int mkdir(char *path, int mode); }
// { int rmdir(char *path); }
// { int utimes(char *path, struct timeval *tptr); }
// { int adjtime(struct timeval *delta, struct timeval *olddelta); }
// { int setsid(void); }
// { int quotactl(char *path, int cmd, int uid, caddr_t arg); }
// { int nlm_syscall(int debug_level, int grace_period, int addr_count, char **addrs); }
// { int nfssvc(int flag, caddr_t argp); }
// { int lgetfh(char *fname, struct fhandle *fhp); }
// { int getfh(char *fname, struct fhandle *fhp); }
// { int sysarch(int op, char *parms); }
// { int rtprio(int function, pid_t pid, struct rtprio *rtp); }
// { int semsys(int which, int a2, int a3, int a4, int a5); }
// { int msgsys(int which, int a2, int a3, int a4, int a5, int a6); }
// { int shmsys(int which, int a2, int a3, int a4); }
// { int setfib(int fibnum); }
// { int ntp_adjtime(struct timex *tp); }
// { int setgid(gid_t gid); }
// { int setegid(gid_t egid); }
// { int seteuid(uid_t euid); }
// { int stat(char *path, struct stat *ub); }
// { int fstat(int fd, struct stat *sb); }
// { int lstat(char *path, struct stat *ub); }
// { int pathconf(char *path, int name); }
// { int fpathconf(int fd, int name); }
// { int getrlimit(u_int which, struct rlimit *rlp); } getrlimit __getrlimit_args int
// { int setrlimit(u_int which, struct rlimit *rlp); } setrlimit __setrlimit_args int
// { int getdirentries(int fd, char *buf, u_int count, long *basep); }
// { int __sysctl(int *name, u_int namelen, void *old, size_t *oldlenp, void *new, size_t newlen); } __sysctl sysctl_args int
// { int mlock(const void *addr, size_t len); }
// { int munlock(const void *addr, size_t len); }
// { int undelete(char *path); }
// { int futimes(int fd, struct timeval *tptr); }
// { int getpgid(pid_t pid); }
// { int poll(struct pollfd *fds, u_int nfds, int timeout); }
// { int semget(key_t key, int nsems, int semflg); }
// { int semop(int semid, struct sembuf *sops, size_t nsops); }
// { int msgget(key_t key, int msgflg); }
// { int msgsnd(int msqid, const void *msgp, size_t msgsz, int msgflg); }
// { int msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg); }
// { int shmat(int shmid, const void *shmaddr, int shmflg); }
// { int shmdt(const void *shmaddr); }
// { int shmget(key_t key, size_t size, int shmflg); }
// { int clock_gettime(clockid_t clock_id, struct timespec *tp); }
// { int clock_settime( clockid_t clock_id, const struct timespec *tp); }
// { int clock_getres(clockid_t clock_id, struct timespec *tp); }
// { int ktimer_create(clockid_t clock_id, struct sigevent *evp, int *timerid); }
// { int ktimer_delete(int timerid); }
// { int ktimer_settime(int timerid, int flags, const struct itimerspec *value, struct itimerspec *ovalue); }
// { int ktimer_gettime(int timerid, struct itimerspec *value); }
// { int ktimer_getoverrun(int timerid); }
// { int nanosleep(const struct timespec *rqtp, struct timespec *rmtp); }
// { int ffclock_getcounter(ffcounter *ffcount); }
// { int ffclock_setestimate( struct ffclock_estimate *cest); }
// { int ffclock_getestimate( struct ffclock_estimate *cest); }
// { int clock_nanosleep(clockid_t clock_id, int flags, const struct timespec *rqtp, struct timespec *rmtp); }
// { int clock_getcpuclockid2(id_t id,int which, clockid_t *clock_id); }
// { int ntp_gettime(struct ntptimeval *ntvp); }
// { int minherit(void *addr, size_t len, int inherit); }
// { int rfork(int flags); }
// { int openbsd_poll(struct pollfd *fds, u_int nfds, int timeout); }
// { int issetugid(void); }
// { int lchown(char *path, int uid, int gid); }
// { int aio_read(struct aiocb *aiocbp); }
// { int aio_write(struct aiocb *aiocbp); }
// { int lio_listio(int mode, struct aiocb * const *acb_list, int nent, struct sigevent *sig); }
// { int getdents(int fd, char *buf, size_t count); }
// { int lchmod(char *path, mode_t mode); }
// { int lutimes(char *path, struct timeval *tptr); }
// { int nstat(char *path, struct nstat *ub); }
// { int nfstat(int fd, struct nstat *sb); }
// { int nlstat(char *path, struct nstat *ub); }
// { ssize_t preadv(int fd, struct iovec *iovp, u_int iovcnt, off_t offset); }
// { ssize_t pwritev(int fd, struct iovec *iovp, u_int iovcnt, off_t offset); }
// { int fhopen(const struct fhandle *u_fhp, int flags); }
// { int fhstat(const struct fhandle *u_fhp, struct stat *sb); }
// { int modnext(int modid); }
// { int modstat(int modid, struct module_stat *stat); }
// { int modfnext(int modid); }
// { int modfind(const char *name); }
// { int kldload(const char *file); }
// { int kldunload(int fileid); }
// { int kldfind(const char *file); }
// { int kldnext(int fileid); }
// { int kldstat(int fileid, struct kld_file_stat* stat); }
// { int kldfirstmod(int fileid); }
// { int getsid(pid_t pid); }
// { int setresuid(uid_t ruid, uid_t euid, uid_t suid); }
// { int setresgid(gid_t rgid, gid_t egid, gid_t sgid); }
// { ssize_t aio_return(struct aiocb *aiocbp); }
// { int aio_suspend( struct aiocb * const * aiocbp, int nent, const struct timespec *timeout); }
// { int aio_cancel(int fd, struct aiocb *aiocbp); }
// { int aio_error(struct aiocb *aiocbp); }
// { int yield(void); }
// { int mlockall(int how); }
// { int munlockall(void); }
// { int __getcwd(char *buf, u_int buflen); }
// { int sched_setparam (pid_t pid, const struct sched_param *param); }
// { int sched_getparam (pid_t pid, struct sched_param *param); }
// { int sched_setscheduler (pid_t pid, int policy, const struct sched_param *param); }
// { int sched_getscheduler (pid_t pid); }
// { int sched_yield (void); }
// { int sched_get_priority_max (int policy); }
// { int sched_get_priority_min (int policy); }
// { int sched_rr_get_interval (pid_t pid, struct timespec *interval); }
// { int utrace(const void *addr, size_t len); }
// { int kldsym(int fileid, int cmd, void *data); }
// { int jail(struct jail *jail); }
// { int sigprocmask(int how, const sigset_t *set, sigset_t *oset); }
// { int sigsuspend(const sigset_t *sigmask); }
// { int sigpending(sigset_t *set); }
// { int sigtimedwait(const sigset_t *set, siginfo_t *info, const struct timespec *timeout); }
// { int sigwaitinfo(const sigset_t *set, siginfo_t *info); }
// { int __acl_get_file(const char *path, acl_type_t type, struct acl *aclp); }
// { int __acl_set_file(const char *path, acl_type_t type, struct acl *aclp); }
// { int __acl_get_fd(int filedes, acl_type_t type, struct acl *aclp); }
// { int __acl_set_fd(int filedes, acl_type_t type, struct acl *aclp); }
// { int __acl_delete_file(const char *path, acl_type_t type); }
// { int __acl_delete_fd(int filedes, acl_type_t type); }
// { int __acl_aclcheck_file(const char *path, acl_type_t type, struct acl *aclp); }
// { int __acl_aclcheck_fd(int filedes, acl_type_t type, struct acl *aclp); }
// { int extattrctl(const char *path, int cmd, const char *filename, int attrnamespace, const char *attrname); }
// { ssize_t extattr_set_file( const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { ssize_t extattr_get_file( const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { int extattr_delete_file(const char *path, int attrnamespace, const char *attrname); }
// { ssize_t aio_waitcomplete( struct aiocb **aiocbp, struct timespec *timeout); }
// { int getresuid(uid_t *ruid, uid_t *euid, uid_t *suid); }
// { int getresgid(gid_t *rgid, gid_t *egid, gid_t *sgid); }
// { int kqueue(void); }
// { int kevent(int fd, struct kevent *changelist, int nchanges, struct kevent *eventlist, int nevents, const struct timespec *timeout); }
// { ssize_t extattr_set_fd(int fd, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { ssize_t extattr_get_fd(int fd, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { int extattr_delete_fd(int fd, int attrnamespace, const char *attrname); }
// { int __setugid(int flag); }
// { int eaccess(char *path, int amode); }
// { int nmount(struct iovec *iovp, unsigned int iovcnt, int flags); }
// { int __mac_get_proc(struct mac *mac_p); }
// { int __mac_set_proc(struct mac *mac_p); }
// { int __mac_get_fd(int fd, struct mac *mac_p); }
// { int __mac_get_file(const char *path_p, struct mac *mac_p); }
// { int __mac_set_fd(int fd, struct mac *mac_p); }
// { int __mac_set_file(const char *path_p, struct mac *mac_p); }
// { int kenv(int what, const char *name, char *value, int len); }
// { int lchflags(const char *path, u_long flags); }
// { int uuidgen(struct uuid *store, int count); }
// { int sendfile(int fd, int s, off_t offset, size_t nbytes, struct sf_hdtr *hdtr, off_t *sbytes, int flags); }
// { int mac_syscall(const char *policy, int call, void *arg); }
// { int getfsstat(struct statfs *buf, long bufsize, int mode); }
// { int statfs(char *path, struct statfs *buf); }
// { int fstatfs(int fd, struct statfs *buf); }
// { int fhstatfs(const struct fhandle *u_fhp, struct statfs *buf); }
// { int ksem_close(semid_t id); }
// { int ksem_post(semid_t id); }
// { int ksem_wait(semid_t id); }
// { int ksem_trywait(semid_t id); }
// { int ksem_init(semid_t *idp, unsigned int value); }
// { int ksem_open(semid_t *idp, const char *name, int oflag, mode_t mode, unsigned int value); }
// { int ksem_unlink(const char *name); }
// { int ksem_getvalue(semid_t id, int *val); }
// { int ksem_destroy(semid_t id); }
// { int __mac_get_pid(pid_t pid, struct mac *mac_p); }
// { int __mac_get_link(const char *path_p, struct mac *mac_p); }
// { int __mac_set_link(const char *path_p, struct mac *mac_p); }
// { ssize_t extattr_set_link( const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { ssize_t extattr_get_link( const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); }
// { int extattr_delete_link( const char *path, int attrnamespace, const char *attrname); }
// { int __mac_execve(char *fname, char **argv, char **envv, struct mac *mac_p); }
// { int sigaction(int sig, const struct sigaction *act, struct sigaction *oact); }
// { int sigreturn( const struct __ucontext *sigcntxp); }
// { int getcontext(struct __ucontext *ucp); }
// { int setcontext( const struct __ucontext *ucp); }
// { int swapcontext(struct __ucontext *oucp, const struct __ucontext *ucp); }
// { int swapoff(const char *name); }
// { int __acl_get_link(const char *path, acl_type_t type, struct acl *aclp); }
// { int __acl_set_link(const char *path, acl_type_t type, struct acl *aclp); }
// { int __acl_delete_link(const char *path, acl_type_t type); }
// { int __acl_aclcheck_link(const char *path, acl_type_t type, struct acl *aclp); }
// { int sigwait(const sigset_t *set, int *sig); }
// { int thr_create(ucontext_t *ctx, long *id, int flags); }
// { void thr_exit(long *state); }
// { int thr_self(long *id); }
// { int thr_kill(long id, int sig); }
// { int jail_attach(int jid); }
// { ssize_t extattr_list_fd(int fd, int attrnamespace, void *data, size_t nbytes); }
// { ssize_t extattr_list_file( const char *path, int attrnamespace, void *data, size_t nbytes); }
// { ssize_t extattr_list_link( const char *path, int attrnamespace, void *data, size_t nbytes); }
// { int ksem_timedwait(semid_t id, const struct timespec *abstime); }
// { int thr_suspend( const struct timespec *timeout); }
// { int thr_wake(long id); }
// { int kldunloadf(int fileid, int flags); }
// { int audit(const void *record, u_int length); }
// { int auditon(int cmd, void *data, u_int length); }
// { int getauid(uid_t *auid); }
// { int setauid(uid_t *auid); }
// { int getaudit(struct auditinfo *auditinfo); }
// { int setaudit(struct auditinfo *auditinfo); }
// { int getaudit_addr( struct auditinfo_addr *auditinfo_addr, u_int length); }
// { int setaudit_addr( struct auditinfo_addr *auditinfo_addr, u_int length); }
// { int auditctl(char *path); }
// { int _umtx_op(void *obj, int op, u_long val, void *uaddr1, void *uaddr2); }
// { int thr_new(struct thr_param *param, int param_size); }
// { int sigqueue(pid_t pid, int signum, void *value); }
// { int kmq_open(const char *path, int flags, mode_t mode, const struct mq_attr *attr); }
// { int kmq_setattr(int mqd,		const struct mq_attr *attr,		struct mq_attr *oattr); }
// { int kmq_timedreceive(int mqd,	char *msg_ptr, size_t msg_len,	unsigned *msg_prio,			const struct timespec *abs_timeout); }
// { int kmq_timedsend(int mqd,		const char *msg_ptr, size_t msg_len,unsigned msg_prio,			const struct timespec *abs_timeout);}
// { int kmq_notify(int mqd,		const struct sigevent *sigev); }
// { int kmq_unlink(const char *path); }
// { int abort2(const char *why, int nargs, void **args); }
// { int thr_set_name(long id, const char *name); }
// { int aio_fsync(int op, struct aiocb *aiocbp); }
// { int rtprio_thread(int function, lwpid_t lwpid, struct rtprio *rtp); }
// { int sctp_peeloff(int sd, uint32_t name); }
// { int sctp_generic_sendmsg(int sd, caddr_t msg, int mlen, caddr_t to, __socklen_t tolen, struct sctp_sndrcvinfo *sinfo, int flags); }
// { int sctp_generic_sendmsg_iov(int sd, struct iovec *iov, int iovlen, caddr_t to, __socklen_t tolen, struct sctp_sndrcvinfo *sinfo, int flags); }
// { int sctp_generic_recvmsg(int sd, struct iovec *iov, int iovlen, struct sockaddr * from, __socklen_t *fromlenaddr, struct sctp_sndrcvinfo *sinfo, int *msg_flags); }
// { ssize_t pread(int fd, void *buf, size_t nbyte, off_t offset); }
// { ssize_t pwrite(int fd, const void *buf, size_t nbyte, off_t offset); }
// { caddr_t mmap(caddr_t addr, size_t len, int prot, int flags, int fd, off_t pos); }
// { off_t lseek(int fd, off_t offset, int whence); }
// { int truncate(char *path, off_t length); }
// { int ftruncate(int fd, off_t length); }
// { int thr_kill2(pid_t pid, long id, int sig); }
// { int shm_open(const char *path, int flags, mode_t mode); }
// { int shm_unlink(const char *path); }
// { int cpuset(cpusetid_t *setid); }
// { int cpuset_setid(cpuwhich_t which, id_t id, cpusetid_t setid); }
// { int cpuset_getid(cpulevel_t level, cpuwhich_t which, id_t id, cpusetid_t *setid); }
// { int cpuset_getaffinity(cpulevel_t level, cpuwhich_t which, id_t id, size_t cpusetsize, cpuset_t *mask); }
// { int cpuset_setaffinity(cpulevel_t level, cpuwhich_t which, id_t id, size_t cpusetsize, const cpuset_t *mask); }
// { int faccessat(int fd, char *path, int amode, int flag); }
// { int fchmodat(int fd, char *path, mode_t mode, int flag); }
// { int fchownat(int fd, char *path, uid_t uid, gid_t gid, int flag); }
// { int fexecve(int fd, char **argv, char **envv); }
// { int fstatat(int fd, char *path, struct stat *buf, int flag); }
// { int futimesat(int fd, char *path, struct timeval *times); }
// { int linkat(int fd1, char *path1, int fd2, char *path2, int flag); }
// { int mkdirat(int fd, char *path, mode_t mode); }
// { int mkfifoat(int fd, char *path, mode_t mode); }
// { int mknodat(int fd, char *path, mode_t mode, dev_t dev); }
// { int openat(int fd, char *path, int flag, mode_t mode); }
// { int readlinkat(int fd, char *path, char *buf, size_t bufsize); }
// { int renameat(int oldfd, char *old, int newfd, char *new); }
// { int symlinkat(char *path1, int fd, char *path2); }
// { int unlinkat(int fd, char *path, int flag); }
// { int posix_openpt(int flags); }
// { int gssd_syscall(char *path); }
// { int jail_get(struct iovec *iovp, unsigned int iovcnt, int flags); }
// { int jail_set(struct iovec *iovp, unsigned int iovcnt, int flags); }
// { int jail_remove(int jid); }
// { int closefrom(int lowfd); }
// { int __semctl(int semid, int semnum, int cmd, union semun *arg); }
// { int msgctl(int msqid, int cmd, struct msqid_ds *buf); }
// { int shmctl(int shmid, int cmd, struct shmid_ds *buf); }
// { int lpathconf(char *path, int name); }
// { int __cap_rights_get(int version, int fd, cap_rights_t *rightsp); }
// { int cap_enter(void); }
// { int cap_getmode(u_int *modep); }
// { int pdfork(int *fdp, int flags); }
// { int pdkill(int fd, int signum); }
// { int pdgetpid(int fd, pid_t *pidp); }
// { int pselect(int nd, fd_set *in, fd_set *ou, fd_set *ex, const struct timespec *ts, const sigset_t *sm); }
// { int getloginclass(char *namebuf, size_t namelen); }
// { int setloginclass(const char *namebuf); }
// { int rctl_get_racct(const void *inbufp, size_t inbuflen, void *outbufp, size_t outbuflen); }
// { int rctl_get_rules(const void *inbufp, size_t inbuflen, void *outbufp, size_t outbuflen); }
// { int rctl_get_limits(const void *inbufp, size_t inbuflen, void *outbufp, size_t outbuflen); }
// { int rctl_add_rule(const void *inbufp, size_t inbuflen, void *outbufp, size_t outbuflen); }
// { int rctl_remove_rule(const void *inbufp, size_t inbuflen, void *outbufp, size_t outbuflen); }
// { int posix_fallocate(int fd, off_t offset, off_t len); }
// { int posix_fadvise(int fd, off_t offset, off_t len, int advice); }
// { int wait6(idtype_t idtype, id_t id, int *status, int options, struct __wrusage *wrusage, siginfo_t *info); }
// { int cap_rights_limit(int fd, cap_rights_t *rightsp); }
// { int cap_ioctls_limit(int fd, const u_long *cmds, size_t ncmds); }
// { ssize_t cap_ioctls_get(int fd, u_long *cmds, size_t maxcmds); }
// { int cap_fcntls_limit(int fd, uint32_t fcntlrights); }
// { int cap_fcntls_get(int fd, uint32_t *fcntlrightsp); }
// { int bindat(int fd, int s, caddr_t name, int namelen); }
// { int connectat(int fd, int s, caddr_t name, int namelen); }
// { int chflagsat(int fd, const char *path, u_long flags, int atflag); }
// { int accept4(int s, struct sockaddr * __restrict name, __socklen_t * __restrict anamelen, int flags); }
// { int pipe2(int *fildes, int flags); }
// { int aio_mlock(struct aiocb *aiocbp); }
// { int procctl(idtype_t idtype, id_t id, int com, void *data); }
// { int ppoll(struct pollfd *fds, u_int nfds, const struct timespec *ts, const sigset_t *set); }
// { int futimens(int fd, struct timespec *times); }
// { int utimensat(int fd, char *path, struct timespec *times, int flag); }
// { int numa_getaffinity(cpuwhich_t which, id_t id, struct vm_domain_policy_entry *policy); }
// { int numa_setaffinity(cpuwhich_t which, id_t id, const struct vm_domain_policy_entry *policy); }
// { int fdatasync(int fd); }
