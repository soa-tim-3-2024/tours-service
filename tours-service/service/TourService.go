package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) GetAuthorTours(authorId int) (*[]model.Tour, error) {
	tours, err := service.TourRepo.GetAuthorTours(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", "-2"))
	}
	return &tours, nil
}

func (service *TourService) Create(tour *model.Tour) error {
	err := service.TourRepo.CreateTour(tour)
	if err != nil {
		return err
	}
	return nil
}