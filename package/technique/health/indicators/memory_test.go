package indicators

import (
	"context"
	"runtime"
	"testing"
)

func TestMemoryIndicator_Up(t *testing.T) {
	// set threshold well above current usage
	ind := NewMemoryIndicator("memory", 99999)
	result := ind.Check(context.Background())

	if !result.Up {
		t.Errorf("Up = false, want true; error = %q", result.Error)
	}
	if _, ok := result.Details["heapUsedMB"]; !ok {
		t.Error("Details should contain heapUsedMB")
	}
	if _, ok := result.Details["thresholdMB"]; !ok {
		t.Error("Details should contain thresholdMB")
	}
}

func TestMemoryIndicator_Down(t *testing.T) {
	// allocate a slice to ensure heap > 0 MB, then set threshold below current usage
	buf := make([]byte, 2*1024*1024) // 2 MB
	_ = buf

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	heapMB := m.HeapAlloc / 1024 / 1024
	if heapMB == 0 {
		t.Skip("heap rounded to 0 MB — cannot test Down reliably")
	}

	ind := NewMemoryIndicator("memory", heapMB-1)
	result := ind.Check(context.Background())

	if result.Up {
		t.Errorf("Up = true, want false (heap %d MB > threshold %d MB)", heapMB, heapMB-1)
	}
	if result.Error == "" {
		t.Error("Error should be set when threshold exceeded")
	}
}

func TestMemoryIndicator_DetailsReflectActualHeap(t *testing.T) {
	ind := NewMemoryIndicator("memory", 99999)
	result := ind.Check(context.Background())

	heapUsed, ok := result.Details["heapUsedMB"].(uint64)
	if !ok {
		t.Fatalf("Details.heapUsedMB type = %T, want uint64", result.Details["heapUsedMB"])
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	expected := m.HeapAlloc / 1024 / 1024

	// allow ±5 MB drift between two consecutive reads
	if heapUsed > expected+5 || (expected > 5 && heapUsed < expected-5) {
		t.Errorf("heapUsedMB = %d, expected ~%d", heapUsed, expected)
	}
}

func TestMemoryIndicator_Key(t *testing.T) {
	ind := NewMemoryIndicator("heap", 99999)
	result := ind.Check(context.Background())

	if result.Key != "heap" {
		t.Errorf("Key = %q, want %q", result.Key, "heap")
	}
}
