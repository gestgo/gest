package indicators

import (
	"context"
	"fmt"
	"runtime"

	"github.com/gestgo/gest/package/technique/health"
)

type MemoryIndicator struct {
	key              string
	heapThresholdMB  uint64
}

func NewMemoryIndicator(key string, heapThresholdMB uint64) *MemoryIndicator {
	return &MemoryIndicator{key: key, heapThresholdMB: heapThresholdMB}
}

func (i *MemoryIndicator) Check(_ context.Context) health.IndicatorResult {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	heapMB := m.HeapAlloc / 1024 / 1024
	details := map[string]any{
		"heapUsedMB":  heapMB,
		"thresholdMB": i.heapThresholdMB,
	}

	if heapMB > i.heapThresholdMB {
		return health.Down(i.key, fmt.Errorf("heap %dMB exceeds threshold %dMB", heapMB, i.heapThresholdMB), details)
	}
	return health.Up(i.key, details)
}
