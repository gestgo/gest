package indicators

import (
	"context"
	"database/sql"

	"github.com/gestgo/gest/package/technique/health"
)

type DBIndicator struct {
	key string
	db  *sql.DB
}

func NewDBIndicator(key string, db *sql.DB) *DBIndicator {
	return &DBIndicator{key: key, db: db}
}

func (i *DBIndicator) Check(ctx context.Context) health.IndicatorResult {
	if err := i.db.PingContext(ctx); err != nil {
		return health.Down(i.key, err, nil)
	}
	return health.Up(i.key, nil)
}
