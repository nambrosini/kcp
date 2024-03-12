package kafka

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	writer *kafka.Writer
}

func NewKafkaWriter(bootstrap, topic, pem, username, password string) KafkaWriter {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{bootstrap},
		Topic:   topic,
		Dialer:  newDialer(pem, username, password),
	})

	return KafkaWriter{
		writer,
	}
}

func (kw *KafkaWriter) Produce(key, value string) {
	defer kw.writer.Close()

	if value == "" {
		value = uuid.New().String()
	}

	err := kw.writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		},
	)
	if err != nil {
		log.Fatal("could not produce:", err)
	}

	log.Println("Message successfully produced...")
}
