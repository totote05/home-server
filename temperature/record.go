package temperature

import (
	"encoding/json"
	"time"

	"github.com/go-playground/validator"
)

type Record struct {
	Source            string    `json:"src" validate:"required"`
	Temperature       float64   `json:"tem" validate:"required"`
	Humidity          float64   `json:"hum" validate:"required"`
	ComputedHeatIndex float64   `json:"chi" validate:"required"`
	Recorded          time.Time `json:"rec" validate:"required"`
}

func (r *Record) ToJSON() []byte {
	value, _ := json.Marshal(r)
	return value
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

	if record.Recorded.IsZero() {
		record.Recorded = time.Now().Truncate(time.Second).UTC()
	} else {
		record.Recorded = record.Recorded.UTC()
	}

	return record, ValidateRecord(record)
}
