package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/audi-skripsi/lambda_event_presenter/cmd/consumer"
	"github.com/audi-skripsi/lambda_event_presenter/internal/component"
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/repository"
	"github.com/audi-skripsi/lambda_event_presenter/internal/service"
	"github.com/audi-skripsi/lambda_event_presenter/pkg/util/logutil"
)

func main() {
	config.Init()
	config := config.Get()

	logger := logutil.NewLogger(logutil.NewLoggerParams{
		PrettyPrint: true,
		ServiceName: config.AppName,
	})

	logger.Infof("app initialized with config of: %+v", config)

	kafkaConsumer, err := component.NewKafkaConsumer(config.KafkaConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing kafka consumer: %+v", err)
	}

	mongo, err := component.NewMongoDB(config.MongoDBConfig)
	if err != nil {
		logger.Fatalf("[main] error initializing mongodb: %+v", err)
	}

	logger.Infof("mongodb connected, ready to listen to connections")

	repository := repository.NewRepository(repository.NewRepositoryParams{
		Logger: logger,
		Config: config,
		Mongo:  mongo,
	})

	service := service.NewService(service.NewServiceParams{
		Logger:     logger,
		Repository: repository,
		Config:     config,
	})

	consumer := consumer.NewConsumer(consumer.NewConsumerParams{
		Logger:        logger,
		KafkaConsumer: kafkaConsumer,
		Config:        config,
		Service:       service,
	})

	consumer.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	go func() {
		go kafkaConsumer.Close()
	}()
	logger.Info("stopping service gracefully...")
	time.Sleep(2 * time.Second)
	logger.Info("service stopped gracefully")
}
