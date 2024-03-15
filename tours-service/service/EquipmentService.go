package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type EquipmentService struct {
	EquipmentRepo *repo.EquipmentRepository
}

func (service *EquipmentService) Create(equipment *model.Equipment) error {
	err := service.EquipmentRepo.Create(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) FindByTourId(id string) (*[]model.Equipment, error) {
	equipment, err := service.EquipmentRepo.FindByTourId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) DeleteById(id string) (*model.Equipment, error) {
	equipment, err := service.EquipmentRepo.DeleteById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &equipment, nil
}

func (service *EquipmentService) Update(equipment *model.Equipment) error {
	err := service.EquipmentRepo.Update(equipment)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) GetAll() (*[]model.Equipment, error) {
	eqs, err := service.EquipmentRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", "-2"))
	}
	return &eqs, nil
}

func (service *EquipmentService) Add(idEq string, idTour string) error {
	err := service.EquipmentRepo.AddEquipmentTour(idEq, idTour)
	if err != nil {
		return err
	}
	return nil
}

func (service *EquipmentService) Remove(idEq string, idTour string) error {
	err := service.EquipmentRepo.RemoveEquipmentTour(idEq, idTour)
	if err != nil {
		return err
	}
	return nil
}
