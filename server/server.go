package server

import "home_api.totote05.ar/temperature"

type Server interface {
	Start()
	RegisterTemperatureHandler(temperature.TemperatureService)
}
