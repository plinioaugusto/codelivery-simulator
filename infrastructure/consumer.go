package kafka

import (
	"fmt"
	"log"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)


type KafkaConsumer struct {
	MessageChannel chan *confluentKafka.Message
}


func NewKafkaConsumer(messageChannel chan *confluentKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MessageChannel: messageChannel,
	}
}
func (k *KafkaConsumer) Consume() {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),

	}

	confluentKafka, error := confluentKafka.NewConsumer(configMap)

	if error != nil {
		log.Fatalf("error consuming kafka message:" + error.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}

	confluentKafka.SubscribeTopics(topics, nil)

	fmt.Println("Kafka consumer has been started")

	for {
		message, error := confluentKafka.ReadMessage(-1)
		if error == nil {
			k.MessageChannel <- message
		}
	}
}