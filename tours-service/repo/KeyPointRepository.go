package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type KeyPointRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *KeyPointRepository) FindById(id string) (model.KeyPoint, error) {
	kP := model.KeyPoint{}
	dbResult := repo.DatabaseConnection.First(&kP, "id = ?", id)
	if dbResult != nil {
		return kP, dbResult.Error
	}
	return kP, nil
}

func (repo *KeyPointRepository) CreateKeyPoint(kp *model.KeyPoint) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.KeyPoint{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	print("ID: ", maxID)
	kp.ID = int(maxID) + 1
	println("KeyPoint: ", kp)
	dbResult := repo.DatabaseConnection.Create(kp)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}