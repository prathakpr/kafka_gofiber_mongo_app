package main

import (
	"context"
	"log"
	
	"github.com/segmentio/kafka-go"
)

func produceToKafka(message string) error {
	writer := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "test-topic",
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key"),
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Println("Failed to write message:", err)
		return err
	}

	log.Println("Message sent to Kafka:", message)
	return nil
}
