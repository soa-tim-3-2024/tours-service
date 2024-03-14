package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type TransportType int

const (
	Walking TransportType = iota
	Bicycle
	Car
)

type Duration struct {
	Duration      int           `json:"duration"`
	TransportType TransportType `json:"transportType"`
}

// Scan implements the sql.Scanner interface
func (d *Duration) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, d)
}

// Value implements the driver.Valuer interface
func (d Duration) Value() (driver.Value, error) {
	return json.Marshal(d)
}
