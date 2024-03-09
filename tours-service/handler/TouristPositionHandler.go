package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type TouristPositionHandler struct {
	TouristPositionService *service.TouristPositionService
}

func (handler *TouristPositionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var position model.TouristPosition
	err := json.NewDecoder(req.Body).Decode(&position)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TouristPositionService.Create(&position)
	if err != nil {
		println("Error while creating tourist position")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(position)
}
