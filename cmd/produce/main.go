package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)
// Importação está bugada

func main() {
	producer := NewKafkaProducer()
	Publish(msg:"Mensagem", topic:"teste", producer, key:nil)
	producer.Flush(1000)
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka:9092",
	}

	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte) error {
	message:= &kafka.Message {
		Value: []byte(msg),
		TopicPartition: kafka.TopicPartition{
			Topic: topic,
			Partition: kafka.PartitionAny,
		},
		Key: key,
	}

	err := producer.Produce(message, deliveryChan: nil)

	if err!= nil {
		return err
	}

	return nil
}