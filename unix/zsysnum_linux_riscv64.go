// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build riscv64,linux

package unix

const (
	SYS_FSTATAT = 79

	SYS_SYNC_FILE_RANGE = 84

	SYS_GETRLIMIT = 163

	SYS_SEMTIMEDOP = 192
	SYS_SEMOP      = 193

	SYS_ACCEPT = 202

	SYS_MMAP      = 222
	SYS_FADVISE64 = 223

	SYS_ARCH_SPECIFIC_SYSCALL = 244

	SYS_KEXEC_FILE_LOAD = 294
)
