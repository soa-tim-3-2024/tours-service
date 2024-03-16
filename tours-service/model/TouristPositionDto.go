package model

type TouristPositionDto struct {
	TourId    int     `json:"tourId"`
	TouristId int     `json:"touristId"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
