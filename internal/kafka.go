package internal

import (
	"log"

	"github.com/IBM/sarama"
)

func sendToKafka(msg string) {
	// Kafka broker address
	brokers := []string{"localhost:9092"}

	// Configuration for the producer
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new synchronous producer
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close Sarama producer: %v", err)
		}
	}()

	// Define the message
	message := &sarama.ProducerMessage{
		Topic: "my_topic",
		Value: sarama.StringEncoder(msg),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
}
