package main

import (
	"log"
	"net/http"
	"os"
	"skoltech/pkg/app"
	"skoltech/pkg/devices"
)

func main() {
	storage := devices.NewStorage(os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"), os.Getenv("KAFKA_TOPIC"))
	service := devices.NewService(storage)
	server := app.NewServer(service)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), server.Handlers()))
}
