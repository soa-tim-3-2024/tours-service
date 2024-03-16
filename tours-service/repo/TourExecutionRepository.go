package repo

import (
	"database-example/model"
	"time"

	"gorm.io/gorm"
)

type TourExecutionRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TourExecutionRepository) CreateTourExecution(tourExecution *model.TourExecution) error {
	var maxID uint
	result := repo.DatabaseConnection.Model(&model.TourExecution{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID)
	if result.Error != nil {
		panic("failed to find maximum ID")
	}
	tourExecution.ID = int(maxID) + 1
	dbResult := repo.DatabaseConnection.Create(tourExecution)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourExecutionRepository) UpdateTourExecution(execution *model.TourExecution) error {
	dbResult := repo.DatabaseConnection.Save(execution)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *TourExecutionRepository) GetTourExecution(tourId int, touristId int) (model.TourExecution, error) {
	execution := model.TourExecution{}
	dbResult := repo.DatabaseConnection.Find(&execution, "tour_id = ? and tourist_id = ? and status = 0", tourId, touristId)
	if dbResult != nil {
		return execution, dbResult.Error
	}
	return execution, nil
}

func (repo *TourExecutionRepository) GetTourExecutionById(tourExecutionId int) (model.TourExecution, error) {
	execution := model.TourExecution{}
	dbResult := repo.DatabaseConnection.Find(&execution, "id", tourExecutionId)
	if dbResult != nil {
		return execution, dbResult.Error
	}
	return execution, nil
}

func (repo *TourExecutionRepository) CompleteTourExecution(execution model.TourExecution) (model.TourExecution, error) {
	execution.LastActivity = time.Now()
	execution.NextKeyPointId = -1
	execution.Status = model.TourExecutionStatus(model.Completed)
	dbResult := repo.DatabaseConnection.Save(execution)
	if dbResult != nil {
		return execution, dbResult.Error
	}
	return execution, nil
}

func (repo *TourExecutionRepository) UpdateNextKeyPoint(keyPointId int, execution model.TourExecution) (model.TourExecution, error) {
	execution.LastActivity = time.Now()
	execution.NextKeyPointId = keyPointId
	dbResult := repo.DatabaseConnection.Save(execution)
	if dbResult != nil {
		return execution, dbResult.Error
	}
	return execution, nil
}
