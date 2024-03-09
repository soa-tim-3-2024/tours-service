package model

type TouristPosition struct {
	ID        int     `json:"id"`
	TouristId int     `json:"touristId"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
