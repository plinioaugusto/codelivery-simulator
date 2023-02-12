package main

import (
	"fmt"
	"log"

	consumerKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	producerKafka "github.com/plinioaugusto/simulator-full-cycle-immersion-2023/application/kafka"
	"github.com/plinioaugusto/simulator-full-cycle-immersion-2023/infrastructure/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	messageChannel := make(chan *consumerKafka.Message)
	consumer := kafka.NewKafkaConsumer(messageChannel)
	go consumer.Consume()
	for message := range messageChannel {
		fmt.Println(string(message.Value))
		go producerKafka.Produce(message)
	}
}