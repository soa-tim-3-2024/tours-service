package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TouristPositionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TouristPositionRepository) GetByTouristId(id int) (model.TouristPosition, error) {
	position := model.TouristPosition{}
	dbResult := repo.DatabaseConnection.First(&position, "tourist_id = ?", id)
	if dbResult.Error != nil {
		return position, dbResult.Error
	}
	return position, nil
}

func (repo *TouristPositionRepository) UpdatePosition(position *model.TouristPosition) {
	result := repo.DatabaseConnection.Model(&model.TouristPosition{}).Where("tourist_id = ?", position.TouristId).Update("longitude", position.Longitude)
	repo.DatabaseConnection.Model(&model.TouristPosition{}).Where("tourist_id = ?", position.TouristId).Update("latitude", position.Latitude)
	if result.Error != nil {
		panic("failed to update position")
	}
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
