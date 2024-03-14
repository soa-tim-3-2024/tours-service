package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type PreferenceRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *PreferenceRepository) FindByUserId(id string) (model.Preference, error) {
	pref := model.Preference{}
	dbResult := repo.DatabaseConnection.First(&pref, "user_id = ?", id)
	if dbResult != nil {
		return pref, dbResult.Error
	}
	return pref, nil
}

func (repo *PreferenceRepository) DeleteById(id string) (model.Preference, error) {
	pref := model.Preference{}
	repo.DatabaseConnection.First(&pref, "id = ?", id)
	dbResult := repo.DatabaseConnection.Delete(&pref)
	if dbResult != nil {
		return pref, dbResult.Error
	}
	return pref, nil
}

func (repo *PreferenceRepository) CreatePreference(pref *model.Preference) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.Preference{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	print("ID: ", maxID)
	pref.ID = int(maxID) + 1
	println("Preference: ", pref)
	dbResult := repo.DatabaseConnection.Create(pref)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *PreferenceRepository) UpdateTour(pref *model.Preference) error {
	dbResult := repo.DatabaseConnection.Save(pref)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}