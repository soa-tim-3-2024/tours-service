package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	"math"
	"time"
)

type TourExecutionService struct {
	TourExecutionRepo *repo.TourExecutionRepository
	KeyPointRepo      *repo.KeyPointRepository
}

func (service *TourExecutionService) CanBeRated(tourId int, userId int) (bool, error) {
	executions, err := service.TourExecutionRepo.GetTourExecutions(tourId, userId)
	if err != nil {
		return false, fmt.Errorf(fmt.Sprintf("tour ex with tour id %d and userId %d not found", tourId, userId))
	}
	if(len(executions) > 0){
		println(len(executions))
		return true, nil
	}else{
		println(len(executions))
		return false, nil
	}
}

func (service *TourExecutionService) Create(tourId int, touristId int) (model.TourExecution, error) {
	var execution model.TourExecution
	execution.TourId = tourId
	execution.TouristId = touristId
	execution.LastActivity = time.Now()
	execution.Progress = 0
	keyPoints, _ := service.KeyPointRepo.FindByTourId(tourId)
	execution.NextKeyPointId = keyPoints[0].ID
	execution.Status = model.TourExecutionStatus(model.Started)
	err := service.TourExecutionRepo.CreateTourExecution(&execution)
	if err != nil {
		return execution, err
	}
	return execution, nil
}

func (service *TourExecutionService) AbandonTour(tourExecutionId int) (model.TourExecution, error) {
	execution, _ := service.TourExecutionRepo.GetTourExecutionById(tourExecutionId)
	execution.Status = model.TourExecutionStatus(model.Abandoned)
	execution.LastActivity = time.Now()
	err := service.TourExecutionRepo.UpdateTourExecution(&execution)
	if err != nil {
		return execution, err
	}
	return execution, nil
}

func (service *TourExecutionService) CheckKeyPointCompletition(position model.TouristPositionDto) (model.TourExecution, error) {
	var execution model.TourExecution
	keyPoints, _ := service.KeyPointRepo.FindByTourId(position.TourId)
	execution, _ = service.TourExecutionRepo.GetTourExecution(position.TourId, position.TouristId)
	for i := 0; i < len(keyPoints); i++ {
		if keyPoints[i].ID == execution.NextKeyPointId {
			if calculateDistance(keyPoints[i].Longitude, keyPoints[i].Latitude, position.Longitude, position.Latitude) < 200 {
				if (i + 1) >= len(keyPoints) {
					execution, _ = service.TourExecutionRepo.CompleteTourExecution(execution)
				} else {
					execution, _ = service.TourExecutionRepo.UpdateNextKeyPoint(keyPoints[i+1].ID, execution)
				}
				break
			}
		}
	}
	return execution, nil
}

func calculateDistance(longitude1 float64, latitude1 float64, longitude2 float64, latitude2 float64) float64 {
	earthRadius := 6371000.0
	latitude1 = latitude1 * math.Pi / 180
	longitude1 = longitude1 * math.Pi / 180
	latitude2 = latitude2 * math.Pi / 180
	longitude2 = longitude2 * math.Pi / 180

	latitudeDistance := latitude2 - latitude1
	longitudeDistance := longitude2 - longitude1

	a := math.Sin(latitudeDistance/2)*math.Sin(latitudeDistance/2) +
		math.Cos(latitude1)*math.Cos(latitude2)*
			math.Sin(longitudeDistance/2)*math.Sin(longitudeDistance/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c

	return distance
}
