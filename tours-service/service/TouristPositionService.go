package service

import (
	"database-example/model"
	"database-example/repo"
)

type TouristPositionService struct {
	TouristPositionRepo *repo.TouristPositionRepository
}

func (service *TouristPositionService) Create(position *model.TouristPosition) error {
	err := service.TouristPositionRepo.AddPosition(position)
	if err != nil {
		return err
	}
	return nil
}
