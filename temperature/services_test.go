package temperature_test

import (
	"testing"

	"home_api.totote05.ar/temperature"
)

func TestHandleRegisterTemperature(t *testing.T) {
	repository := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repository)
	service := temperature.NewMockTemperatureService(temp)

	record := []byte(`"Temperature": 20.0, "Humidity": 19.0, "ComputedHeatIndex": 25.3}`)
	if err := service.HandleRegisterTemperature(record); err == nil {
		t.Logf("%s should not be parsed", record)
		t.Fail()
	}

	record = []byte(`{"Temperature": 20.0, "Humidity": 19.0, "ComputedHeatIndex": 25.3}`)
	if err := service.HandleRegisterTemperature(record); err == nil {
		t.Logf("%s missing fields", record)
		t.Fail()
	}

	record = []byte(`{"src": "galeria", "tem": 20.0, "hum": 19.0, "chi": 25.3}`)
	if err := service.HandleRegisterTemperature(record); err != nil {
		t.Logf("failed to register value %s, error: %s", record, err)
		t.Fail()
	}
}
