package model

type Review struct {
	ID            int         `json:"id"`
	Rating        int         `json:"rating"`
	Comment       string      `json:"comment"`
	Username      string      `json:"username"`
	TouristId     int         `json:"touristId"`
	CommentDate   string      `json:"commentDate"`
	TourVisitDate string      `json:"tourVisitDate"`
	TourId        int         `json:"tourId"`
	Images        StringArray `json:"images"`
}