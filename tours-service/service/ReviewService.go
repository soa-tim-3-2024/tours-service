package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type ReviewService struct {
	ReviewRepo *repo.ReviewRepository
}

func (service *ReviewService) GetReviews(tourId int) (*[]model.Review, error) {
	reviews, err := service.ReviewRepo.GetReviews(tourId)
	if err != nil {
		return nil, fmt.Errorf("error while getting reviews")
	}
	return &reviews, nil
}

func (service *ReviewService) Create(review *model.Review) error {
	err := service.ReviewRepo.CreateReview(review)
	if err != nil {
		return err
	}
	return nil
}
