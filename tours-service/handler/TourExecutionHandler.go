package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TourExecutionHandler struct {
	TourExecutionService *service.TourExecutionService
}

func (handler *TourExecutionHandler) CanTourBeRated(writer http.ResponseWriter, req *http.Request) {
	idTour := mux.Vars(req)["tourId"]
	idUser := mux.Vars(req)["userId"]

	s1, err1 := strconv.Atoi(idTour)
	if err1 != nil {
		fmt.Println("Can't convert tour id to int!")
	}
	s2, err2 := strconv.Atoi(idUser)
	if err2 != nil {
		fmt.Println("Can't convert user id to int!")
	}
	tours, err := handler.TourExecutionService.CanBeRated(s1, s2)
	
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["tourId"]
	tourId, _ := strconv.Atoi(id)
	id = mux.Vars(req)["touristId"]
	touristId, _ := strconv.Atoi(id)
	execution, err := handler.TourExecutionService.Create(tourId, touristId)
	if err != nil {
		println("Error while creating tour execution session")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(execution)
}

func (handler *TourExecutionHandler) AbandonTour(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tourExecutionId, _ := strconv.Atoi(id)
	execution, err := handler.TourExecutionService.AbandonTour(tourExecutionId)
	if err != nil {
		println("Error while creating tour execution session")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(execution)

}

func (handler *TourExecutionHandler) CheckKeyPointCompletition(writer http.ResponseWriter, req *http.Request) {
	var position model.TouristPositionDto
	json.NewDecoder(req.Body).Decode(&position)
	execution, err := handler.TourExecutionService.CheckKeyPointCompletition(position)
	if err != nil {
		println("Error while checking key point completition")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(execution)
}
