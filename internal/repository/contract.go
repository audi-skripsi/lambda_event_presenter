package repository

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
}

type repository struct {
	logger *logrus.Entry
	config *repositoryConfig
	mongo  *mongo.Database
}

type repositoryConfig struct {
	kafkaConfig config.KafkaConfig
}

type NewRepositoryParams struct {
	Logger *logrus.Entry
	Config *config.Config
	Mongo  *mongo.Database
}

func NewRepository(params NewRepositoryParams) Repository {
	return &repository{
		logger: params.Logger,
		mongo:  params.Mongo,
		config: &repositoryConfig{
			kafkaConfig: params.Config.KafkaConfig,
		},
	}
}
