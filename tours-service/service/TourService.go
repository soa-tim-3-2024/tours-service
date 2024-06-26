package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type TourService struct {
	TourRepo *repo.TourRepository
}

func (service *TourService) GetAuthorTours(authorId int) ([]model.Tour, error) {
	tours, err := service.TourRepo.GetAuthorTours(authorId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("tours with author id %d not found", authorId))
	}
	return tours, nil
}

func (service *TourService) GetPublishedTours() (*[]model.Tour, error) {
	tours, err := service.TourRepo.GetPublishedTours()
	if err != nil {
		return nil, fmt.Errorf("error while getting published tours")
	}
	return &tours, nil
}

func (service *TourService) GetTours(Ids []int) (*[]model.Tour, error) {
	tours, err := service.TourRepo.GetTours(Ids)
	if err != nil {
		return nil, fmt.Errorf("error while getting published tours")
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

func (service *TourService) FindTour(id string) (*model.Tour, error) {
	tour, err := service.TourRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &tour, nil
}

func (service *TourService) Update(tour *model.Tour) error {
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) PublishTour(tour *model.Tour) error {
	tour.TourStatus = model.Status(model.Published)
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}

func (service *TourService) ArchiveTour(tour *model.Tour) error {
	tour.TourStatus = model.Status(model.Archived)
	err := service.TourRepo.UpdateTour(tour)
	if err != nil {
		return err
	}
	return nil
}
