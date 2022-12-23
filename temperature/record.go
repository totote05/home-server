package temperature

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

type Record struct {
	Source            string  `json:"src" validate:"required"`
	Temperature       float64 `json:"tem" validate:"required"`
	Humidity          float64 `json:"hum" validate:"required"`
	ComputedHeatIndex float64 `json:"chi" validate:"required"`
}

func ValidateRecord(record Record) error {
	v := validator.New()
	return v.Struct(record)
}

func RecordFromJson(value []byte) (Record, error) {
	var record Record

	if err := json.Unmarshal(value, &record); err != nil {
		return record, err
	}

	return record, ValidateRecord(record)
}
