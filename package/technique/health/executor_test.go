package health

import (
	"context"
	"errors"
	"testing"
)

type stubIndicator struct {
	key    string
	result IndicatorResult
}

func (s *stubIndicator) Check(_ context.Context) IndicatorResult {
	return s.result
}

func upStub(key string) *stubIndicator {
	return &stubIndicator{key: key, result: Up(key, nil)}
}

func downStub(key string) *stubIndicator {
	return &stubIndicator{key: key, result: Down(key, errors.New("check failed"), nil)}
}

func TestExecutor_AllUp(t *testing.T) {
	e := NewExecutor(upStub("db"), upStub("redis"))

	report := e.Execute(context.Background())

	if report.Status != StatusOK {
		t.Errorf("Status = %q, want %q", report.Status, StatusOK)
	}
	if _, ok := report.Info["db"]; !ok {
		t.Error("Info should contain db")
	}
	if _, ok := report.Info["redis"]; !ok {
		t.Error("Info should contain redis")
	}
	if len(report.Error) != 0 {
		t.Errorf("Error map should be empty, got %v", report.Error)
	}
}

func TestExecutor_OneDown(t *testing.T) {
	e := NewExecutor(upStub("db"), downStub("redis"))

	report := e.Execute(context.Background())

	if report.Status != StatusError {
		t.Errorf("Status = %q, want %q", report.Status, StatusError)
	}
	if _, ok := report.Info["db"]; !ok {
		t.Error("Info should contain healthy db")
	}
	if _, ok := report.Error["redis"]; !ok {
		t.Error("Error should contain failing redis")
	}
}

func TestExecutor_AllDown(t *testing.T) {
	e := NewExecutor(downStub("db"), downStub("redis"))

	report := e.Execute(context.Background())

	if report.Status != StatusError {
		t.Errorf("Status = %q, want %q", report.Status, StatusError)
	}
	if len(report.Info) != 0 {
		t.Errorf("Info should be empty, got %v", report.Info)
	}
}

func TestExecutor_NoIndicators(t *testing.T) {
	e := NewExecutor()

	report := e.Execute(context.Background())

	if report.Status != StatusOK {
		t.Errorf("Status = %q, want %q, no indicators means healthy", report.Status, StatusOK)
	}
}

func TestExecutor_DetailsAlwaysPopulated(t *testing.T) {
	ind := &stubIndicator{
		key:    "db",
		result: Down("db", errors.New("timeout"), map[string]any{"latencyMs": 5000}),
	}
	e := NewExecutor(ind)

	report := e.Execute(context.Background())

	detail, ok := report.Details["db"].(map[string]any)
	if !ok {
		t.Fatal("Details[db] should be map[string]any")
	}
	if detail["latencyMs"] != 5000 {
		t.Errorf("Details should carry indicator details")
	}
	if detail["error"] != "timeout" {
		t.Errorf("Details should carry error string")
	}
}

func TestExecutor_ShuttingDown(t *testing.T) {
	e := NewExecutor(upStub("db"))
	e.MarkShuttingDown()

	report := e.Execute(context.Background())

	if report.Status != StatusShuttingDown {
		t.Errorf("Status = %q, want %q", report.Status, StatusShuttingDown)
	}
	if len(report.Info) != 0 || len(report.Error) != 0 {
		t.Error("ShuttingDown report should not run indicators")
	}
}

func TestExecutor_ShuttingDownSkipsIndicators(t *testing.T) {
	checked := false
	spy := &stubIndicator{key: "db"}
	spy.result = IndicatorResult{Key: "db", Up: true}

	e := NewExecutor(&spyIndicator{key: "db", checked: &checked})
	e.MarkShuttingDown()
	e.Execute(context.Background())

	if checked {
		t.Error("indicators should not be called when shutting down")
	}
}

type spyIndicator struct {
	key     string
	checked *bool
}

func (s *spyIndicator) Check(_ context.Context) IndicatorResult {
	*s.checked = true
	return Up(s.key, nil)
}
