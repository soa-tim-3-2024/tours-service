package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// enum u go-u
type Status int

type StringArray []string

const (
	Draft Status = iota
	Published
	Archived
	Ready
)

type Tour struct {
	ID          int         `json:"id"`
	AuthorId    int         `json:"authorId"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Difficulty  int         `json:"difficulty"`
	Category    int         `json:"category"`
	Tags        StringArray `json:"tags"`
	TourStatus  Status      `json:"tourStatus"`
	Price       float64     `json:"price"`
	Distance    float64     `json:"distance"`
}

/*
func (tour *Tour) BeforeCreate(scope *gorm.DB) error {
	tour.ID = uuid.New()
	return nil
}
*/
// konvertuje tip podatka iz go-a u tip podatka u bazi (jer gorm ne moze sam da rukuje sa nizom stringova kao atributom)
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Konvertuje iz tipa podatka iz baze u tip podatka u go-u
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}
