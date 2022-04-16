package service

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/repository"
	"github.com/sirupsen/logrus"
)

type Service interface {
}

type service struct {
	logger     *logrus.Entry
	repository repository.Repository
	config     *serviceConfig
}

type serviceConfig struct {
	KafkaConfig *config.KafkaConfig
}

type NewServiceParams struct {
	Logger     *logrus.Entry
	Repository repository.Repository
	Config     *config.Config
}

func NewService(params NewServiceParams) Service {
	return &service{
		logger:     params.Logger,
		repository: params.Repository,
		config: &serviceConfig{
			KafkaConfig: &params.Config.KafkaConfig,
		},
	}
}
