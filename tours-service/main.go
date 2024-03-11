package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=super dbname=gorm port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Student{})
	database.AutoMigrate(&model.Tour{})
	database.AutoMigrate(&model.TouristPosition{})
	return database
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	studentRepo := &repo.StudentRepository{DatabaseConnection: database}
	studentService := &service.StudentService{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentService: studentService}

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	tourService := &service.TourService{TourRepo: tourRepo}
	tourHandler := &handler.TourHandler{TourService: tourService}

	touristPositionRepo := &repo.TouristPositionRepository{DatabaseConnection: database}
	touristPositionService := &service.TouristPositionService{TouristPositionRepo: touristPositionRepo}
	touristPositionHandler := &handler.TouristPositionHandler{TouristPositionService: touristPositionService}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", studentHandler.Get).Methods("GET")
	router.HandleFunc("/students", studentHandler.Create).Methods("POST")
	router.HandleFunc("/tours/{authorId}", tourHandler.Get).Methods("GET")
	router.HandleFunc("/tours", tourHandler.Create).Methods("POST")
	router.HandleFunc("/tours", tourHandler.Update).Methods("PUT")
	router.HandleFunc("/touristposition", touristPositionHandler.Create).Methods("POST")
	router.HandleFunc("/touristposition/{tourist_id}", touristPositionHandler.GetByTouristId).Methods("GET")
	router.HandleFunc("/touristposition", touristPositionHandler.Update).Methods("PUT")
	router.HandleFunc("/tour/{id}", tourHandler.GetById).Methods("GET")

	// Set up CORS middleware
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
