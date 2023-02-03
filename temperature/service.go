package temperature

type TemperatureService interface {
	HandleRegisterTemperature(value []byte) error
	HandleTemperatureHistory(func(<-chan Record))
	HandleLastValue() []byte
}

type Service struct {
	handler *Temperature
}

func NewService(temp *Temperature) *Service {
	return &Service{handler: temp}
}

func (s *Service) HandleRegisterTemperature(value []byte) error {
	if record, err := RecordFromJson(value); err != nil {
		return err
	} else {
		s.handler.Register(record)
		return nil
	}
}

func (s *Service) HandleTemperatureHistory(resp func(records <-chan Record)) {
	resp(s.handler.GetHistory())
}

func (s *Service) HandleLastValue() []byte {
	last := s.handler.GetLastValue()

	if last == nil {
		return nil
	} else {
		return last.ToJSON()
	}
}

var _ TemperatureService = (*Service)(nil)
