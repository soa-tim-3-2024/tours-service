package model

type Preference struct {
	ID              int         `json:"id"`
	UserId          int         `json:"userId"`
	DifficultyLevel int         `json:"difficultyLevel"`
	WalkingRating   int         `json:"walkingRating"`
	CyclingRating   int         `json:"cyclingRating"`
	CarRating       int         `json:"carRating"`
	BoatRating      int         `json:"boatRating"`
	SelectedTags    StringArray `json:"selectedTags"`
}