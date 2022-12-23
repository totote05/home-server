package temperature

type MockTemperatureService struct {
	handler *Temperature
}

func NewMockTemperatureService(temp *Temperature) *MockTemperatureService {
	return &MockTemperatureService{handler: temp}
}

func (s *MockTemperatureService) HandleRegisterTemperature(value []byte) error {
	if record, err := RecordFromJson(value); err != nil {
		return err
	} else {
		s.handler.Register(record)
		return nil
	}
}

var _ TemperatureService = (*MockTemperatureService)(nil)
