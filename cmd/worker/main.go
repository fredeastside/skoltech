package main

import (
	"encoding/json"
	"log"
	"os"
	"skoltech/pkg/devices"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	storage := devices.NewStorage(os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"), os.Getenv("KAFKA_TOPIC"))
	client := devices.NewClient(os.Getenv("PARTNER_URL"), os.Getenv("PARTNER_PORT"))
	service := devices.NewService(storage)
	storage.Read(func(m kafka.Message, err error) {
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		log.Printf("Read message from kafka: %v", string(m.Value))
		d := service.CreateDevice()
		err = json.Unmarshal(m.Value, d)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
		_, err = client.Send(d)
		if err != nil {
			log.Printf("ERROR: %v", err)
			return
		}
	})
}
