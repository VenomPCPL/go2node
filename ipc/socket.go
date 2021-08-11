package ipc

import (
	"golang.org/x/sys/unix"
	"os"
	"strconv"
)

// Socketpair create a socketpair
func Socketpair() ([]*os.File, error) {
	fds, err := unix.Socketpair(unix.AF_LOCAL, unix.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}
	return []*os.File{
		os.NewFile(uintptr(fds[0]), "@/go2node/left"+strconv.Itoa(fds[0])),
		os.NewFile(uintptr(fds[1]), "@/go2node/right"+strconv.Itoa(fds[1])),
	}, err
}
