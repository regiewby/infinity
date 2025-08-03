package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/infinity/identity-service/server/config"
	"github.com/sirupsen/logrus"
)

func NewKafkaProducer(kafkaConfig *config.KafkaConfig, log *logrus.Logger) *kafka.Producer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":            kafkaConfig.Brokers,
		"group.id":                     kafkaConfig.GroupID,
		"auto.offset.reset":            kafkaConfig.InitOffset,
		"client.id":                    kafkaConfig.ClientID,
		"acks":                         kafkaConfig.RequiredAcks,
		"message.timeout.ms":           5000,
		"queue.buffering.max.messages": 100000,
		"queue.buffering.max.kbytes":   1048576,
		"request.timeout.ms":           30000,
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	return producer
}
