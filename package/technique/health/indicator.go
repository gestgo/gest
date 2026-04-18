package health

import "context"

type Indicator interface {
	Check(ctx context.Context) IndicatorResult
}
