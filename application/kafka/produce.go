package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/plinioaugusto/simulator-full-cycle-immersion-2023/application/route"
	"github.com/plinioaugusto/simulator-full-cycle-immersion-2023/infrastructure/kafka"
)


func Produce(message *confluentKafka.Message) {
	producer := kafka.NewKafkaProducer()

	route := route.NewRoute()

	json.Unmarshal(message.Value, &route)

	route.LoadPositions()

	positions, error := route.ExportJsonPositions()

	if error != nil {
		log.Println(error.Error())
	}

	for _, position := range positions {
		kafka.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}