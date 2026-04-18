package health

import (
	"context"
	"errors"
	"testing"
)

type mockLogger struct {
	lastMsg  string
	lastArgs []any
	errCalls int
	infoCalls int
}

func (l *mockLogger) Info(msg string, args ...any) {
	l.lastMsg = msg
	l.lastArgs = args
	l.infoCalls++
}

func (l *mockLogger) Error(msg string, args ...any) {
	l.lastMsg = msg
	l.lastArgs = args
	l.errCalls++
}

func TestService_Check_AllHealthy(t *testing.T) {
	e := NewExecutor(upStub("db"), upStub("cache"))
	log := &mockLogger{}
	svc := NewService(e, log)

	report, err := svc.Check(context.Background())

	if err != nil {
		t.Errorf("err = %v, want nil", err)
	}
	if report.Status != StatusOK {
		t.Errorf("Status = %q, want %q", report.Status, StatusOK)
	}
	if log.errCalls != 0 {
		t.Error("logger.Error should not be called when healthy")
	}
}

func TestService_Check_Unhealthy(t *testing.T) {
	e := NewExecutor(upStub("db"), downStub("cache"))
	log := &mockLogger{}
	svc := NewService(e, log)

	report, err := svc.Check(context.Background())

	if !errors.Is(err, ErrUnhealthy) {
		t.Errorf("err = %v, want ErrUnhealthy", err)
	}
	if report.Status != StatusError {
		t.Errorf("Status = %q, want %q", report.Status, StatusError)
	}
	if log.errCalls != 1 {
		t.Errorf("logger.Error called %d times, want 1", log.errCalls)
	}
}

func TestService_Check_ShuttingDown(t *testing.T) {
	e := NewExecutor(upStub("db"))
	svc := NewService(e, nil)
	svc.MarkShuttingDown()

	report, err := svc.Check(context.Background())

	if !errors.Is(err, ErrUnhealthy) {
		t.Errorf("err = %v, want ErrUnhealthy", err)
	}
	if report.Status != StatusShuttingDown {
		t.Errorf("Status = %q, want %q", report.Status, StatusShuttingDown)
	}
}

func TestService_Check_NilLogger(t *testing.T) {
	e := NewExecutor(downStub("db"))
	svc := NewService(e, nil)

	_, err := svc.Check(context.Background())

	if !errors.Is(err, ErrUnhealthy) {
		t.Errorf("err = %v, want ErrUnhealthy", err)
	}
}
