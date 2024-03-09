package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TouristPositionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TouristPositionRepository) AddPosition(position *model.TouristPosition) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.TouristPosition{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	position.ID = int(maxID) + 1
	dbResult := repo.DatabaseConnection.Create(position)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
