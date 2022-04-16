package consumer

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/audi-skripsi/lambda_event_presenter/internal/service"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	logger        *logrus.Entry
	kafkaConsumer *kafka.Consumer
	config        *consumerConfig
	service       service.Service
}

type consumerConfig struct {
	KafkaConfig config.KafkaConfig
}

type NewConsumerParams struct {
	Logger        *logrus.Entry
	KafkaConsumer *kafka.Consumer
	Config        *config.Config
	Service       service.Service
}

func NewConsumer(params NewConsumerParams) *Consumer {
	return &Consumer{
		logger:        params.Logger,
		kafkaConsumer: params.KafkaConsumer,
		service:       params.Service,
		config: &consumerConfig{
			KafkaConfig: params.Config.KafkaConfig,
		},
	}
}

func (c *Consumer) Init() {
	err := c.kafkaConsumer.Subscribe(c.config.KafkaConfig.InTopic, nil)
	if err != nil {
		c.logger.Fatalf("error subscribing to topic: %+v", err)
		return
	}
	c.logger.Infof("kafka ready to listen to messages at topic: %s", c.config.KafkaConfig.InTopic)
	// go func() {
	// 	for {
	// 		msg, err := c.kafkaConsumer.ReadMessage(-1)
	// 		if err != nil {
	// 			c.logger.Errorf("error receiving message: %+v", err)
	// 			continue
	// 		}
	// 		c.logger.Infof("received message: %v", string(msg.Key))

	// 		var eventLog dto.EventLog
	// 		err = json.Unmarshal(msg.Value, &eventLog)
	// 		if err != nil {
	// 			c.logger.Errorf("error unmarshalling message: %+v", err)
	// 			c.kafkaConsumer.CommitMessage(msg)
	// 			continue
	// 		}

	// 		err = c.service.StandardizeLevelAndPush(eventLog)
	// 		if err != nil {
	// 			c.logger.Errorf("error handling event %+v: %+v", eventLog, err)
	// 		}

	// 		c.kafkaConsumer.CommitMessage(msg)
	// 	}
	// }()
}
