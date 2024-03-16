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

func (handler *TourExecutionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["tourId"]
	tourId, _ := strconv.Atoi(id)
	id = mux.Vars(req)["touristId"]
	touristId, _ := strconv.Atoi(id)
	fmt.Println("sve ok?", tourId)
	fmt.Println("sve ok?", touristId)
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
