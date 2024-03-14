package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type PreferenceService struct {
	PrefRepo *repo.PreferenceRepository
}

func (service *PreferenceService) Create(pref *model.Preference) error {
	err := service.PrefRepo.CreatePreference(pref)
	if err != nil {
		return err
	}
	return nil
}

func (service *PreferenceService) FindPreferenceByUserId(id string) (*model.Preference, error) {
	pref, err := service.PrefRepo.FindByUserId(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &pref, nil
}

func (service *PreferenceService) DeleteById(id string) (*model.Preference, error) {
	pref, err := service.PrefRepo.DeleteById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &pref, nil
}

func (service *PreferenceService) Update(pref *model.Preference) error {
	err := service.PrefRepo.UpdateTour(pref)
	if err != nil {
		return err
	}
	return nil
}