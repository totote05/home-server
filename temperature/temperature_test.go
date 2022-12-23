package temperature_test

import (
	"reflect"
	"testing"

	"home_api.totote05.ar/temperature"
)

func CheckRecordStructure(t *testing.T, record temperature.Record) {
	if err := temperature.ValidateRecord(record); err != nil {
		t.Logf("invalid structure of %#v", record)
		t.Fail()
	}
}

func TestRegisterTemperature(t *testing.T) {
	repository := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repository)

	record := temperature.Record{
		Source:            "galeria",
		Temperature:       20.0,
		Humidity:          19.0,
		ComputedHeatIndex: 25.3,
	}
	CheckRecordStructure(t, record)

	if size := temp.Register(record); size != 1 {
		t.Logf("Expected %d got %d\n", 1, size)
		t.Fail()
	}

	record = temperature.Record{
		Source:            "galeria",
		Temperature:       21.0,
		Humidity:          19.0,
		ComputedHeatIndex: 25.5,
	}
	CheckRecordStructure(t, record)
	if size := temp.Register(record); size != 2 {
		t.Logf("Expected %d got %d\n", 2, size)
		t.Fail()
	}
}

func TestGetTemperatureHistory(t *testing.T) {
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

	for _, value := range values {
		CheckRecordStructure(t, value)
	}
	repository := temperature.NewMockTemperatureRepository(values)
	temp := temperature.NewTemperature(repository)

	var res []temperature.Record
	for record := range temp.GetHistory() {
		res = append(res, record)
	}

	if eq := reflect.DeepEqual(values, res); !eq {
		t.Log("History received is not equal to registered")
		t.Fail()
	}
}

func TestGetLastValue(t *testing.T) {
	record1 := temperature.Record{
		Source:            "galeria",
		Temperature:       20.0,
		Humidity:          19.0,
		ComputedHeatIndex: 25.3,
	}
	CheckRecordStructure(t, record1)
	record2 := temperature.Record{
		Source:            "galeria",
		Temperature:       21.0,
		Humidity:          19.0,
		ComputedHeatIndex: 25.5,
	}
	CheckRecordStructure(t, record2)
	record3 := temperature.Record{
		Source:            "galeria",
		Temperature:       21.0,
		Humidity:          19.0,
		ComputedHeatIndex: 25.5,
	}
	CheckRecordStructure(t, record3)

	repository := temperature.NewMockTemperatureRepository([]temperature.Record{})
	temp := temperature.NewTemperature(repository)

	if value := temp.GetLastValue(); value != nil {
		t.Log("the last must be nil when is empty")
		t.Fail()
	}

	temp.Register(record1)

	if value := temp.GetLastValue(); !reflect.DeepEqual(*value, record1) {
		t.Log("Last value is not equal to record 1")
		t.Fail()
	}

	temp.Register(record2)
	temp.Register(record3)

	if value := temp.GetLastValue(); !reflect.DeepEqual(*value, record3) {
		t.Log("Last value is not equal to record 3")
		t.Fail()
	}
}
