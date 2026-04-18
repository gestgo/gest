package health

import "context"

type Service struct {
	executor *Executor
	logger   Logger
}

func NewService(executor *Executor, logger Logger) *Service {
	return &Service{executor: executor, logger: logger}
}

func (s *Service) Check(ctx context.Context) (HealthReport, error) {
	report := s.executor.Execute(ctx)
	if report.Status != StatusOK {
		if s.logger != nil {
			s.logger.Error("health check failed", "status", report.Status, "errors", report.Error)
		}
		return report, ErrUnhealthy
	}
	return report, nil
}

func (s *Service) MarkShuttingDown() {
	s.executor.MarkShuttingDown()
}
