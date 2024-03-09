package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type TourRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourRepository) GetAuthorTours(authorId int) ([]model.Tour, error) {
	tours := []model.Tour{}
	dbResult := repo.DatabaseConnection.Where("author_id = ?", authorId).Find(&tours)
	if dbResult != nil {
		return tours, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) UpdateTour(tour *model.Tour) error {
	dbResult := repo.DatabaseConnection.Save(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourRepository) CreateTour(tour *model.Tour) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.Tour{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	print("ID: ", maxID)
	tour.ID = int(maxID) + 1
	println("Tour: ", tour)
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
