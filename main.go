package main

import (
	"log"

	"home_api.totote05.ar/core"
	"home_api.totote05.ar/server"
	"home_api.totote05.ar/temperature"
)

func main() {
	env := core.GetEnv()
	repo := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repo)
	svce := temperature.NewService(temp)
	serv := server.NewHttpServer(env.HttpHost)

	serv.RegisterTemperatureHandler(svce)
	serv.Start()
	log.Println("App terminated")
}
