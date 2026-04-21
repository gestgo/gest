//go:build linux || darwin

package version

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func termWidth() int {
	if s := os.Getenv("COLUMNS"); s != "" {
		if n, _ := strconv.Atoi(s); n > 0 {
			return n
		}
	}
	var ws [4]uint16 // rows, cols, xpixel, ypixel
	if _, _, e := syscall.Syscall(syscall.SYS_IOCTL, 1, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws[0]))); e == 0 && ws[1] > 0 {
		return int(ws[1])
	}
	return 80
}
