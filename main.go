package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"home_api.totote05.ar/temperature"
)

func main() {
	var service temperature.TemperatureService

	go service.Start()
	log.Println("Service started")

	s := make(chan<- os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)

	service.Stop()
	log.Println("Service stopped")
	log.Println("App terminated")
}
