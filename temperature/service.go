package temperature

type TemperatureService interface {
	HandleRegisterTemperature(value []byte) error
	HandleTemperatureHistory(func(<-chan Record))
	HandleLastValue() []byte
	Start()
	Stop()
}
