package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	reader *kafka.Reader
}

func NewKafkaReader(bootstrap, groupId, topic, pem, username, password string) KafkaReader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{bootstrap},
		GroupID: groupId,
		Topic:   topic,
		Dialer:  newDialer(pem, username, password),
	})

	return KafkaReader{
		reader,
	}
}

func (kr *KafkaReader) Read() {
	defer kr.reader.Close()

	fmt.Println("start consuming...!!")

	for {
		m, err := kr.reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf(
			"message at topic:%v partition:%v offset:%v key: %s = value: %s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)
	}
}
