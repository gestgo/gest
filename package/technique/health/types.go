package health

import "errors"

type HealthStatus string

const (
	StatusOK           HealthStatus = "ok"
	StatusError        HealthStatus = "error"
	StatusShuttingDown HealthStatus = "shutting_down"
)

type IndicatorResult struct {
	Key     string
	Up      bool
	Details map[string]any
	Error   string
}

type HealthReport struct {
	Status  HealthStatus   `json:"status"`
	Info    map[string]any `json:"info,omitempty"`
	Error   map[string]any `json:"error,omitempty"`
	Details map[string]any `json:"details"`
}

var ErrUnhealthy = errors.New("health check failed")
