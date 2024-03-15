package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type EquipmentHandler struct {
	EquipmentService *service.EquipmentService
}

func (handler *EquipmentHandler) GetByTourId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Equipment sa id-em ture %s", id)
	equipment, err := handler.EquipmentService.FindByTourId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	fmt.Printf("%+v\n", equipment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.Create(&equipment)
	if err != nil {
		println("Error while creating a new Equipment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var equipment model.Equipment
	err := json.NewDecoder(req.Body).Decode(&equipment)
	fmt.Printf("%+v\n", equipment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.EquipmentService.Update(&equipment)
	if err != nil {
		println("Error while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Brisanje Equipment sa id-em %s", id)
	equipment, err := handler.EquipmentService.DeleteById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(equipment)
}

func (handler *EquipmentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	eqs, err := handler.EquipmentService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(eqs)
}

func (handler *EquipmentHandler) Add(writer http.ResponseWriter, req *http.Request) {
	idEq := mux.Vars(req)["idEq"]
	idTour := mux.Vars(req)["idTour"]
	log.Printf("Dodavanje Equipment sa id-em %s na turu sa id-e %s", idEq, idTour)
	
	err := handler.EquipmentService.Add(idEq, idTour)
	if err != nil {
		println("Error while creating a new EquipmentTour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *EquipmentHandler) Remove(writer http.ResponseWriter, req *http.Request) {
	idEq := mux.Vars(req)["idEq"]
	idTour := mux.Vars(req)["idTour"]
	log.Printf("Brisanje Equipment sa id-em %s sa ture sa id-e %s", idEq, idTour)
	
	err := handler.EquipmentService.Remove(idEq, idTour)
	if err != nil {
		println("Error while creating a new EquipmentTour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}