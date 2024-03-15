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
	durations := []model.Duration{}
	dbResult := repo.DatabaseConnection.Where("author_id = ?", authorId).Omit("durations").Find(&tours)
	for i := range tours {
		repo.DatabaseConnection.Model(&model.Tour{}).Where("id=?", tours[i].ID).Pluck("durations", &durations)
		tours[i].Durations = durations
	}
	if dbResult != nil {
		return tours, dbResult.Error
	}
	return tours, nil
}

func (repo *TourRepository) FindById(id string) (model.Tour, error) {
	tour := model.Tour{}
	durations := []model.Duration{}
	dbResult := repo.DatabaseConnection.Preload("KeyPoints").Omit("durations").First(&tour, "id = ?", id)

	repo.DatabaseConnection.Model(&model.Tour{}).Pluck("durations", &durations).Where("id=?", tour.ID)
	tour.Durations = durations

	if tour.Durations == nil {
		tour.Durations = []model.Duration{}
	}
	if dbResult != nil {
		return tour, dbResult.Error
	}
	return tour, nil
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
	tour.ID = int(maxID) + 1
	dbResult := repo.DatabaseConnection.Create(tour)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
