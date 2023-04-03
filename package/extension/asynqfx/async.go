package asynqfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type HandleJobFunc func(ctx context.Context, t *asynq.Task) error
type JobHandle struct {
	Type       string
	HandleFunc HandleJobFunc
}

func RegisterAsynqJobQueueHooks(
	lifecycle fx.Lifecycle, srv *asynq.Server, jobHandlers []JobHandle, logger *zap.SugaredLogger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				router.InitRouter(params.KafkaTopics)

				go func() error {
					if err := srv.Run(mux); err != nil {
						logger.Fatalf("could not run server: %v", err)
					}
					return nil
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}
