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

type ReviewHandler struct {
	ReviewService *service.ReviewService
}

func (handler *ReviewHandler) GetReviewsByTourId(writer http.ResponseWriter, req *http.Request) {
	fmt.Printf("dweefffrrgr")
	idTour := mux.Vars(req)["tourId"]
	s1, err1 := strconv.Atoi(idTour)
	if err1 != nil {
		fmt.Println("Can't convert tour id to int!")
	}
	reviews, err := handler.ReviewService.GetReviews(s1)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(reviews)
}

func (handler *ReviewHandler) Create(writer http.ResponseWriter, req *http.Request) {
	
	var review model.Review
	err := json.NewDecoder(req.Body).Decode(&review)
	fmt.Printf("%+v\n", review)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.ReviewService.Create(&review)
	if err != nil {
		println("Error while creating a new review")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(review)
}