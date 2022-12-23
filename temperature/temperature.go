package temperature

type Temperature struct {
	repository TemperatureRepository
}

func NewTemperature(repository TemperatureRepository) *Temperature {
	return &Temperature{
		repository: repository,
	}
}

func (t *Temperature) Register(value Record) int {
	t.repository.Add(value)
	return t.repository.Size()
}

func (t *Temperature) GetHistory() <-chan Record {
	return t.repository.GetStream()
}

func (t *Temperature) GetLastValue() *Record {
	var last *Record

	for v := range t.repository.GetStream() {
		last = &v
	}

	return last
}
