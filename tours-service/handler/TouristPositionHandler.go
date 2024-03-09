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

type TouristPositionHandler struct {
	TouristPositionService *service.TouristPositionService
}

func (handler *TouristPositionHandler) GetByTouristId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["tourist_id"]
	s, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Can't convert this to an int!")
	}
	position, err := handler.TouristPositionService.GetByTouristId(s)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(position)
}

func (handler *TouristPositionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var position model.TouristPosition
	err := json.NewDecoder(req.Body).Decode(&position)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	handler.TouristPositionService.Update(&position)

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(position)
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
