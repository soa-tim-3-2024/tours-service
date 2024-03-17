package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyPointHandler struct {
	KeyPointService *service.KeyPointService
}

func (handler *KeyPointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("KeyPoint sa id-em %s", id)
	keyPoint, err := handler.KeyPointService.FindKeyPoint(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(keyPoint)
}

func (handler *KeyPointHandler) GetKeyPoints(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tourId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Can't convert to int!")
	}
	keyPoint, err := handler.KeyPointService.FindKeyPoints(tourId)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(keyPoint)
}

func (handler *KeyPointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	//fmt.Println(req)
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	if err != nil {
		fmt.Println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Create(&keyPoint)

	if err != nil {
		fmt.Println("Error while creating a new keyPoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keyPoint)
}


func (handler *KeyPointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var keyPoint model.KeyPoint
	err := json.NewDecoder(req.Body).Decode(&keyPoint)
	fmt.Printf("%+v\n", keyPoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.KeyPointService.Update(&keyPoint)
	if err != nil {
		println("Error while updating keypoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(keyPoint)
}

func (handler *KeyPointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Brisanje KeyPointa sa id-em %s", id)
	kp, err := handler.KeyPointService.DeleteById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(kp)
}