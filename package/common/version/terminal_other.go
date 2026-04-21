//go:build !linux && !darwin

package version

import (
	"os"
	"strconv"
)

func termWidth() int {
	if s := os.Getenv("COLUMNS"); s != "" {
		if n, _ := strconv.Atoi(s); n > 0 {
			return n
		}
	}
	return 80
}
