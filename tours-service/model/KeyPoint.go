package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type KeyPoint struct {
	ID                  int         	`json:"id"`
	TourId              int       		`json:"tourId"`
	Name                string			`json:"name"`
	Description         string    		`json:"description"`
	Longitude           float64   		`json:"longitude"`
	Latitude            float64   		`json:"latitude"`
	LocationAddress     string    		`json:"locationAddress"`
	ImagePath           string    		`json:"imagePath"`
	Order               int       		`json:"order"`
	IsEncounterRequired bool      		`json:"isEncounterRequired"`
	HasEncounter        bool      		`json:"hasEncounter"`
}

type KeyPointSecret struct {
	Description string    	`json:"description"`  
	Images      string 		`json:"images"`
}

func (r KeyPointSecret) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *KeyPointSecret) Scan(value interface{}) error {
	if value == nil {
		*r = KeyPointSecret{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Scan source is not []byte")
	}
	return json.Unmarshal(bytes, r)
}