package indicators

import (
	"context"
	"testing"
)

func TestDiskIndicator_Up(t *testing.T) {
	// threshold 100% → always passes
	ind := NewDiskIndicator("disk", "/", 100.0)
	result := ind.Check(context.Background())

	if !result.Up {
		t.Errorf("Up = false, want true; error = %q", result.Error)
	}
	if _, ok := result.Details["usedPercent"]; !ok {
		t.Error("Details should contain usedPercent")
	}
	if _, ok := result.Details["thresholdPercent"]; !ok {
		t.Error("Details should contain thresholdPercent")
	}
}

func TestDiskIndicator_Down(t *testing.T) {
	// threshold 0% → always fails
	ind := NewDiskIndicator("disk", "/", 0.0)
	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false when threshold is 0%")
	}
	if result.Error == "" {
		t.Error("Error should be set when threshold exceeded")
	}
}

func TestDiskIndicator_InvalidPath(t *testing.T) {
	ind := NewDiskIndicator("disk", "/nonexistent/path/xyz", 80.0)
	result := ind.Check(context.Background())

	if result.Up {
		t.Error("Up = true, want false for invalid path")
	}
	if result.Error == "" {
		t.Error("Error should be set for invalid path")
	}
}

func TestDiskIndicator_Key(t *testing.T) {
	ind := NewDiskIndicator("root-disk", "/", 100.0)
	result := ind.Check(context.Background())

	if result.Key != "root-disk" {
		t.Errorf("Key = %q, want %q", result.Key, "root-disk")
	}
}

func TestDiskIndicator_DetailsFormat(t *testing.T) {
	ind := NewDiskIndicator("disk", "/", 100.0)
	result := ind.Check(context.Background())

	usedPct, ok := result.Details["usedPercent"].(string)
	if !ok {
		t.Fatalf("Details.usedPercent type = %T, want string", result.Details["usedPercent"])
	}
	// should look like "xx.xx%"
	if len(usedPct) < 2 || usedPct[len(usedPct)-1] != '%' {
		t.Errorf("usedPercent format unexpected: %q", usedPct)
	}
}
