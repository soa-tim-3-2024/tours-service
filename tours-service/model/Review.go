package model

type Review struct {
	ID            int         `json:"id"`
	Rating        int         `json:"rating"`
	Comment       string      `json:"comment"`
	TouristId     int         `json:"touristId"`
	CommentDate   string      `json:"commentDate"`
	TourVisitDate string      `json:"tourVisitDate"`
	TourId        int         `json:"tourId"`
	Images        StringArray `json:"images"`
}