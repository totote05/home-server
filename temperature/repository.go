package temperature

type TemperatureRepository interface {
	Add(Record)
	Size() int
	GetStream() <-chan Record
}
