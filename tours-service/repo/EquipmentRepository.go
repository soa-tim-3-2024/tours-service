package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type EquipmentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *EquipmentRepository) FindByTourId(id string) ([]model.Equipment, error) {
	var tour model.Tour
	err := repo.DatabaseConnection.Model(&model.Tour{}).Omit("durations").Preload("Equipment").Where("id = ?", id).First(&tour).Error
	if err != nil {
		return tour.Equipment, err
	}
	return tour.Equipment, nil
}

func (repo *EquipmentRepository) DeleteById(id string) (model.Equipment, error) {
	equipment := model.Equipment{}
	repo.DatabaseConnection.First(&equipment, "id = ?", id)
	dbResult := repo.DatabaseConnection.Delete(&equipment)
	if dbResult != nil {
		return equipment, dbResult.Error
	}
	return equipment, nil
}

func (repo *EquipmentRepository) Create(equipment *model.Equipment) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.Equipment{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	print("ID: ", maxID)
	equipment.ID = int(maxID) + 1
	println("Equipment: ", equipment)
	dbResult := repo.DatabaseConnection.Create(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) Update(equipment *model.Equipment) error {
	dbResult := repo.DatabaseConnection.Save(equipment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) GetAll() ([]model.Equipment, error) {
	eqs := []model.Equipment{}
	dbResult := repo.DatabaseConnection.Find(&eqs)
	if dbResult != nil {
		return eqs, dbResult.Error
	}
	return eqs, nil
}

func (repo *EquipmentRepository) AddEquipmentTour(idEq string, idTour string) error {
	var eq model.Equipment
	dbResult := repo.DatabaseConnection.First(&eq, "id = ?", idEq)
	var tour model.Tour
	dbResult2 := repo.DatabaseConnection.Preload("KeyPoints").Omit("durations").First(&tour, "id = ?", idTour)
	repo.DatabaseConnection.Model(&tour).Association("Equipment").Append(&eq)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult2.Error != nil {
		return dbResult2.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *EquipmentRepository) RemoveEquipmentTour(idEq string, idTour string) error {
	var eq model.Equipment
	dbResult := repo.DatabaseConnection.First(&eq, "id = ?", idEq)
	var tour model.Tour
	dbResult2 := repo.DatabaseConnection.Preload("KeyPoints").Omit("durations").First(&tour, "id = ?", idTour)
	repo.DatabaseConnection.Model(&tour).Association("Equipment").Delete(&eq)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	if dbResult2.Error != nil {
		return dbResult2.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}