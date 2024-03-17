package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *ReviewRepository) GetReviews(tourId int) ([]model.Review, error) {
	reviews := []model.Review{}
	dbResult := repo.DatabaseConnection.Find(&reviews, "tour_id = ?", tourId)
	// for i := range reviews {
	// 	repo.DatabaseConnection.Model(&model.review{}).Where("id=?", reviews[i].ID).Pluck("durations", &durations)
	// }
	if dbResult != nil {
		return reviews, dbResult.Error
	}
	return reviews, nil
}

func (repo *ReviewRepository) CreateReview(review *model.Review) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.Review{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	review.ID = int(maxID) + 1
	dbResult := repo.DatabaseConnection.Create(review)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
