package repository

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/model"
	pkgdto "github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetAllMicroservicesName(ctx context.Context) (collections []string, err error)
	StoreMicroservicesData(ctx context.Context, microservicesData []pkgdto.PublicMicroserviceData) (err error)

	InsertEvent(event model.EventLog, collName string) (result model.EventLog, err error)
	BatchInsertEvent(eventBatch *dto.EventBatch) (err error)
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
