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

type PreferenceHandler struct {
	PreferenceService *service.PreferenceService
}


func (handler *PreferenceHandler) GetByUserId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Preference sa id-em %s", id)
	pref, err := handler.PreferenceService.FindPreferenceByUserId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(pref)
}



func (handler *PreferenceHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var pref model.Preference
	err := json.NewDecoder(req.Body).Decode(&pref)
	fmt.Printf("%+v\n", pref)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PreferenceService.Create(&pref)
	if err != nil {
		println("Error while creating a new preference")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pref)
}

func (handler *PreferenceHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var pref model.Preference
	err := json.NewDecoder(req.Body).Decode(&pref)
	fmt.Printf("%+v\n", pref)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.PreferenceService.Update(&pref)
	if err != nil {
		println("Error while updating tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(pref)
}

func  (handler *PreferenceHandler) Delete(writer http.ResponseWriter, req *http.Request) {
    id := mux.Vars(req)["id"]
	log.Printf("Brisanje preference sa id-em %s", id)
	pref, err := handler.PreferenceService.DeleteById(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
    json.NewEncoder(writer).Encode(pref)
}