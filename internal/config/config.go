package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	AppAddress    string
	KafkaConfig   KafkaConfig
	MongoDBConfig MongoDBConfig
	RedisConfig   RedisConfig
}

var config *Config

func Init() {
	err := godotenv.Load("conf/.env")
	if err != nil {
		log.Printf("[Init] error on loading env from file: %+v", err)
	}

	config = &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppAddress: os.Getenv("APP_ADDRESS"),
		KafkaConfig: KafkaConfig{
			Address:       os.Getenv("KAFKA_ADDRESS"),
			InTopic:       os.Getenv("KAFKA_IN_TOPIC"),
			ConsumerGroup: os.Getenv("KAFKA_CONSUMER_GROUP"),
		},
		MongoDBConfig: MongoDBConfig{
			DBName:    os.Getenv("MONGODB_DB_NAME"),
			DBAddress: os.Getenv("MONGODB_ADDRESS"),
		},
		RedisConfig: RedisConfig{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}

	if config.AppName == "" {
		log.Panicf("[Init] app name cannot be empty")
	}

	if config.AppAddress == "" {
		log.Panicf("[Init] app address cannot be empty")
	}

	if config.KafkaConfig.Address == "" ||
		config.KafkaConfig.InTopic == "" {
		log.Panicf("[Init] kafka config cannot be empty")
	}

	if config.MongoDBConfig.DBAddress == "" ||
		config.MongoDBConfig.DBName == "" {
		log.Panic("[Init] mongodb config cannot be empty")
	}

	if config.RedisConfig.Address == "" {
		log.Panic("[Init] redis address cannot be empty")
	}

}

func Get() *Config {
	return config
}
