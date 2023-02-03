package temperature_test

import (
	"reflect"
	"testing"

	"home_api.totote05.ar/temperature"
)

func TestHandleRegisterTemperature(t *testing.T) {
	repository := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repository)
	service := temperature.NewService(temp)

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

	record = []byte(`{"src": "galeria", "tem": 20.0, "hum": 19.0, "chi": 25.3, "rec": "2023-02-02T23:01:00Z-03:00"}`)
	if err := service.HandleRegisterTemperature(record); err == nil {
		t.Logf("failed to register value %s, error: %s", record, err)
		t.Fail()
	}

	record = []byte(`{"src": "galeria", "tem": 20.0, "hum": 19.0, "chi": 25.3}`)
	if err := service.HandleRegisterTemperature(record); err != nil {
		t.Logf("failed to register value %s, error: %s", record, err)
		t.Fail()
	}
}

func TestHandleTemperatureHistory(t *testing.T) {
	values := []temperature.Record{
		{
			Source:            "galeria",
			Temperature:       20.0,
			Humidity:          19.0,
			ComputedHeatIndex: 25.3,
		},
		{
			Source:            "galeria",
			Temperature:       21.0,
			Humidity:          19.0,
			ComputedHeatIndex: 25.5,
		},
	}
	repository := temperature.NewMockTemperatureRepository(values)
	temp := temperature.NewTemperature(repository)
	service := temperature.NewService(temp)

	service.HandleTemperatureHistory(func(records <-chan temperature.Record) {
		var history []temperature.Record
		for record := range records {
			history = append(history, record)
		}

		if eq := reflect.DeepEqual(values, history); !eq {
			t.Log("History received is not equal to registered")
			t.Fail()
		}
	})
}

func TestHandleLastValue(t *testing.T) {
	json1 := []byte(`{
		"src": "galeria",
		"tem": 20.0,
		"hum": 19.0,
		"chi": 25.3
	}`)
	json2 := []byte(`{
		"src": "galeria",
		"tem": 21.0,
		"hum": 19.0,
		"chi": 25.5
	}`)
	json3 := []byte(`{
		"src": "galeria",
		"tem": 21.0,
		"hum": 19.0,
		"chi": 25.5
	}`)

	record1, err := temperature.RecordFromJson(json1)
	if err != nil {
		t.Log("can't decode json1", err)
		t.Fail()
	}

	record3, err := temperature.RecordFromJson(json3)
	if err != nil {
		t.Log("can't decode json3", err)
		t.Fail()
	}

	repository := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repository)
	service := temperature.NewService(temp)

	if value := service.HandleLastValue(); value != nil {
		t.Log("the last must be nil when is empty")
		t.Fail()
	}

	service.HandleRegisterTemperature(json1)

	value, _ := temperature.RecordFromJson(service.HandleLastValue())

	if !reflect.DeepEqual(value, record1) {
		t.Log("Last value is not equal to record 1")
		t.Fail()
	}

	service.HandleRegisterTemperature(json2)
	service.HandleRegisterTemperature(json3)

	value, _ = temperature.RecordFromJson(service.HandleLastValue())

	if !reflect.DeepEqual(value, record3) {
		t.Log("Last value is not equal to record 3")
		t.Fail()
	}
}
