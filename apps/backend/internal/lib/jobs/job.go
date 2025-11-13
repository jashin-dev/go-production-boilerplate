package job

import (
	"github.com/hibiken/asynq"
	"github.com/jashin-dev/go-production-boilerplate/internal/config"
	"github.com/rs/zerolog"
)

type JobService struct {
	Client *asynq.Client
	server *asynq.Server
	logger *zerolog.Logger
}

func NewJobService(logger *zerolog.Logger, cfg *config.Config) *JobService {
	redisAddr := cfg.Redis.Address

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisAddr,
	})

	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redisAddr,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return &JobService{
		Client: client,
		server: server,
		logger: logger,
	}
}

func (j *JobService) Stop() {
	j.logger.Info().Msg("Stopping background job server")
	j.server.Shutdown()
	j.Client.Close()
}
