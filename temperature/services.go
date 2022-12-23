package temperature

type TemperatureService interface {
	HandleRegisterTemperature(value []byte) error
	// Start()
}
