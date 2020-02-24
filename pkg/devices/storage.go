package devices

import (
	"context"
	"encoding/json"
	kafka "github.com/segmentio/kafka-go"
	"log"
)

//Storage structure for storing requests in kafka
type Storage struct {
	url   string
	topic string
}

//NewStorage constructor for storage structure
func NewStorage(host, port, topic string) *Storage {
	return &Storage{url: host + ":" + port, topic: topic}
}

//Save device in storage (producing)
func (s *Storage) Save(d *Device) error {

	message, err := json.Marshal(d)
	if err != nil {
		return err
	}

	log.Printf("Write message to kafka: %v", string(message))
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{s.url},
		Topic:    s.topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	return w.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)
}

//Read messages from storage (consuming)
func (s *Storage) Read(f func(kafka.Message, error)) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{s.url},
		Topic:   s.topic,
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		f(m, err)
	}
}
