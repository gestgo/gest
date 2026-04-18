package indicators

import (
	"context"
	"fmt"
	"syscall"

	"github.com/gestgo/gest/package/technique/health"
)

type DiskIndicator struct {
	key              string
	path             string
	thresholdPercent float64
}

func NewDiskIndicator(key, path string, thresholdPercent float64) *DiskIndicator {
	return &DiskIndicator{key: key, path: path, thresholdPercent: thresholdPercent}
}

func (i *DiskIndicator) Check(_ context.Context) health.IndicatorResult {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(i.path, &stat); err != nil {
		return health.Down(i.key, err, nil)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free
	usedPercent := float64(used) / float64(total) * 100

	details := map[string]any{
		"path":             i.path,
		"usedPercent":      fmt.Sprintf("%.2f%%", usedPercent),
		"thresholdPercent": fmt.Sprintf("%.2f%%", i.thresholdPercent),
	}

	if usedPercent > i.thresholdPercent {
		return health.Down(i.key, fmt.Errorf("disk usage %.2f%% exceeds threshold %.2f%%", usedPercent, i.thresholdPercent), details)
	}
	return health.Up(i.key, details)
}
