package kafka

import (
	"log"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)


func NewKafkaProducer() *confluentKafka.Producer {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),

	}

	producer, error := confluentKafka.NewProducer(configMap)

	if error != nil {
		log.Println(error.Error())
	}
	return producer
}

func Publish(messageToProduce string, topic string, producer *confluentKafka.Producer) error {
	message := &confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &topic, Partition: confluentKafka.PartitionAny},
		Value:          []byte(messageToProduce),
	}

	error := producer.Produce(message, nil)
	
	if error != nil {
		return error
	}

	return nil
}