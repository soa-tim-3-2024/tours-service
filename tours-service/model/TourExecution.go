package model

import (
	"time"
)

type TourExecutionStatus int

const (
	Started TourExecutionStatus = iota
	Abandoned
	Completed
)

type TourExecution struct {
	ID             int                 `json:"id"`
	TourId         int                 `json:"tourId"`
	TouristId      int                 `json:"touristId"`
	NextKeyPointId int                 `json:"nextKeyPointId"`
	Progress       float64             `json:"progress"`
	LastActivity   time.Time           `json:"lastActivity"`
	Status         TourExecutionStatus `json:"status"`
}
