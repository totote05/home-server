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

func (s *MockTemperatureService) HandleTemperatureHistory(resp func(records <-chan Record)) {
	resp(s.handler.GetHistory())
}

func (s *MockTemperatureService) HandleLastValue() []byte {
	last := s.handler.GetLastValue()

	if last == nil {
		return nil
	} else {
		return last.ToJSON()
	}
}

func (s *MockTemperatureService) Start() {}
func (s *MockTemperatureService) Stop()  {}

var _ TemperatureService = (*MockTemperatureService)(nil)
