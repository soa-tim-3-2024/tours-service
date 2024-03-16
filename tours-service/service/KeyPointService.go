package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type KeyPointService struct {
	KeyPointRepo *repo.KeyPointRepository
}

func (service *KeyPointService) FindKeyPoint(id string) (*model.KeyPoint, error) {
	kP, err := service.KeyPointRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &kP, nil
}

func (service *KeyPointService) FindKeyPoints(tourId int) (*[]model.KeyPoint, error) {
	kP, err := service.KeyPointRepo.FindByTourId(tourId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %d not found", tourId))
	}
	return &kP, nil
}

func (service *KeyPointService) Create(kP *model.KeyPoint) error {
	err := service.KeyPointRepo.CreateKeyPoint(kP)
	if err != nil {
		return err
	}
	return nil
}
