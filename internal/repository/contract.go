package repository

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	SegragateCollection(name string) (err error)
	InsertEvent(event model.EventLog, collName string) (result model.EventLog, err error)
}

type repository struct {
	logger *logrus.Entry
	config *repositoryConfig
	mongo  *mongo.Database
	redis  *redis.Client
}

type repositoryConfig struct {
	kafkaConfig config.KafkaConfig
}

type NewRepositoryParams struct {
	Logger *logrus.Entry
	Config *config.Config
	Mongo  *mongo.Database
	Redis  *redis.Client
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger: params.Logger,
		mongo:  params.Mongo,
		redis:  params.Redis,
		config: &repositoryConfig{
			kafkaConfig: params.Config.KafkaConfig,
		},
	}
}
