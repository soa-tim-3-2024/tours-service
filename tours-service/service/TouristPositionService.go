package service

import (
	"database-example/model"
	"database-example/repo"
)

type TouristPositionService struct {
	TouristPositionRepo *repo.TouristPositionRepository
}

func (service *TouristPositionService) GetByTouristId(id int) (model.TouristPosition, error) {
	position, err := service.TouristPositionRepo.GetByTouristId(id)
	if err != nil {
		return position, err
	}
	return position, nil
}

func (service *TouristPositionService) Create(position *model.TouristPosition) error {
	err := service.TouristPositionRepo.AddPosition(position)
	if err != nil {
		return err
	}
	return nil
}

func (service *TouristPositionService) Update(position *model.TouristPosition) {
	service.TouristPositionRepo.UpdatePosition(position)
}
