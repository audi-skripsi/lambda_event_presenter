package service

import (
	"context"

	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	indto "github.com/audi-skripsi/lambda_event_presenter/internal/dto"
	"github.com/audi-skripsi/lambda_event_presenter/internal/repository"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/dto"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetAllMicroservicesData(ctx context.Context) (microservicesData dto.PublicMicroservicesNameResponse, err error)
	GetMicroserviceDataAnalytics(ctx context.Context, id string) (resp dto.PublicMicroserviceAnalyticsResponse, err error)
	GetMicroserviceEvents(ctx context.Context, id string, criteria indto.SearchEventCriteria) (resp dto.PublicMicroserviceEventSearchResponse, err error)

	StoreEvent(event dto.EventLog) (err error)
	Ping() (resp dto.PublicPingResponse)
}

type service struct {
	logger     *logrus.Entry
	repository repository.Repository
	config     *serviceConfig
	BatchMap   map[string]*indto.EventBatch
}

type serviceConfig struct {
	KafkaConfig *config.KafkaConfig
	BatchConfig *config.BatchConfig
}

type NewServiceParams struct {
	Logger     *logrus.Entry
	Repository repository.Repository
	Config     *config.Config
}

func NewService(params NewServiceParams) Service {
	s := &service{
		logger:     params.Logger,
		repository: params.Repository,
		config: &serviceConfig{
			KafkaConfig: &params.Config.KafkaConfig,
			BatchConfig: &params.Config.BatchConfig,
		},
	}
	s.BatchMap = make(map[string]*indto.EventBatch)
	s.initBatchCron()
	return s
}
