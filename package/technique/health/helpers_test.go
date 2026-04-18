package health

import (
	"errors"
	"testing"
)

func TestUp(t *testing.T) {
	details := map[string]any{"responseTime": "5ms"}
	result := Up("db", details)

	if result.Key != "db" {
		t.Errorf("Key = %q, want %q", result.Key, "db")
	}
	if !result.Up {
		t.Error("Up should be true")
	}
	if result.Error != "" {
		t.Errorf("Error should be empty, got %q", result.Error)
	}
	if result.Details["responseTime"] != "5ms" {
		t.Errorf("Details not preserved")
	}
}

func TestUp_NilDetails(t *testing.T) {
	result := Up("memory", nil)

	if result.Key != "memory" {
		t.Errorf("Key = %q, want %q", result.Key, "memory")
	}
	if !result.Up {
		t.Error("Up should be true")
	}
	if result.Details != nil {
		t.Error("Details should be nil when not provided")
	}
}

func TestDown(t *testing.T) {
	err := errors.New("connection refused")
	details := map[string]any{"host": "localhost:5432"}
	result := Down("db", err, details)

	if result.Key != "db" {
		t.Errorf("Key = %q, want %q", result.Key, "db")
	}
	if result.Up {
		t.Error("Up should be false")
	}
	if result.Error != "connection refused" {
		t.Errorf("Error = %q, want %q", result.Error, "connection refused")
	}
	if result.Details["host"] != "localhost:5432" {
		t.Errorf("Details not preserved")
	}
}

func TestDown_NilError(t *testing.T) {
	result := Down("http", nil, nil)

	if result.Up {
		t.Error("Up should be false")
	}
	if result.Error != "" {
		t.Errorf("Error should be empty when nil err, got %q", result.Error)
	}
}
