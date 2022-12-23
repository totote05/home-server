package temperature

type MockTemperatureRepository struct {
	values []Record
}

func NewMockTemperatureRepository(values []Record) *MockTemperatureRepository {
	return &MockTemperatureRepository{values: values}
}

func (r *MockTemperatureRepository) Add(value Record) {
	r.values = append(r.values, value)
}

func (r *MockTemperatureRepository) Size() int {
	return len(r.values)
}

func (r *MockTemperatureRepository) GetStream() <-chan Record {
	records := make(chan Record)

	go func() {
		for _, value := range r.values {
			records <- value
		}
		close(records)
	}()

	return records
}

var _ TemperatureRepository = (*MockTemperatureRepository)(nil)
