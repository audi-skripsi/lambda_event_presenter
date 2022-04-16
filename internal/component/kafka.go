package component

import (
	"github.com/audi-skripsi/lambda_event_presenter/internal/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func NewKafkaConsumer(config config.KafkaConfig) (consumer *kafka.Consumer, err error) {
	conf := kafka.ConfigMap{}
	conf["bootstrap.servers"] = config.Address
	if config.ConsumerGroup != "" {
		conf["group.id"] = config.ConsumerGroup
	}
	conf["auto.offset.reset"] = "smallest"
	conf["enable.auto.commit"] = false

	consumer, err = kafka.NewConsumer(&conf)
	return
}
