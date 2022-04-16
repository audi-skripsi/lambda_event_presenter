package component

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(config config.RedisConfig) (redisClient *redis.Client, err error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Address,
		Password: config.Password,
		DB:       0, // using default db
	})

	return rdb, rdb.Ping(context.Background()).Err()
}
