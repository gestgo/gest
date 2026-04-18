package health

import (
	"context"
	"sync/atomic"
)

type Executor struct {
	indicators     []Indicator
	isShuttingDown atomic.Bool
}

func NewExecutor(indicators ...Indicator) *Executor {
	return &Executor{indicators: indicators}
}

func (e *Executor) MarkShuttingDown() {
	e.isShuttingDown.Store(true)
}

func (e *Executor) Execute(ctx context.Context) HealthReport {
	if e.isShuttingDown.Load() {
		return HealthReport{
			Status:  StatusShuttingDown,
			Details: make(map[string]any),
		}
	}

	report := HealthReport{
		Status:  StatusOK,
		Info:    make(map[string]any),
		Error:   make(map[string]any),
		Details: make(map[string]any),
	}

	for _, ind := range e.indicators {
		result := ind.Check(ctx)
		detail := buildDetail(result)
		report.Details[result.Key] = detail

		if result.Up {
			report.Info[result.Key] = detail
		} else {
			report.Error[result.Key] = detail
			report.Status = StatusError
		}
	}

	return report
}

func buildDetail(r IndicatorResult) map[string]any {
	d := make(map[string]any)
	for k, v := range r.Details {
		d[k] = v
	}
	if r.Error != "" {
		d["error"] = r.Error
	}
	return d
}
