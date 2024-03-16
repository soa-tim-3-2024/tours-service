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
	database.AutoMigrate(&model.KeyPoint{})
	database.AutoMigrate(&model.Preference{})
	database.AutoMigrate(&model.Equipment{})
	database.AutoMigrate(&model.TourExecution{})
	return database
}

func configureToursHandler(router *mux.Router, tourHandler *handler.TourHandler) {
	router.HandleFunc("/tours/published/all", tourHandler.FindPublishedTours).Methods("GET")
	router.HandleFunc("/tours/tours-list", tourHandler.GetToursByIds).Methods("POST")
}

func configureTourExecutionHandler(router *mux.Router, tourExecutionHandler *handler.TourExecutionHandler) {
	router.HandleFunc("/tour-execution/{tourId}/{touristId}", tourExecutionHandler.Create).Methods("GET")
	router.HandleFunc("/tour-execution/check-completition", tourExecutionHandler.CheckKeyPointCompletition).Methods("POST")
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

	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepo}
	keyPointHandler := &handler.KeyPointHandler{KeyPointService: keyPointService}

	touristPositionRepo := &repo.TouristPositionRepository{DatabaseConnection: database}
	touristPositionService := &service.TouristPositionService{TouristPositionRepo: touristPositionRepo}
	touristPositionHandler := &handler.TouristPositionHandler{TouristPositionService: touristPositionService}

	preferenceRepo := &repo.PreferenceRepository{DatabaseConnection: database}
	preferenceService := &service.PreferenceService{PrefRepo: preferenceRepo}
	preferenceHandler := &handler.PreferenceHandler{PreferenceService: preferenceService}

	eqRepo := &repo.EquipmentRepository{DatabaseConnection: database}
	eqService := &service.EquipmentService{EquipmentRepo: eqRepo}
	eqHandler := &handler.EquipmentHandler{EquipmentService: eqService}

	tourExecutionRepo := &repo.TourExecutionRepository{DatabaseConnection: database}
	tourExecutionService := &service.TourExecutionService{TourExecutionRepo: tourExecutionRepo, KeyPointRepo: keyPointRepo}
	tourExecutionHandler := &handler.TourExecutionHandler{TourExecutionService: tourExecutionService}

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
	router.HandleFunc("/tours/publish", tourHandler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive", tourHandler.Archive).Methods("PUT")
	router.HandleFunc("/keyPoints", keyPointHandler.Create).Methods("POST")
	router.HandleFunc("/preference/{id}", preferenceHandler.GetByUserId).Methods("GET")
	router.HandleFunc("/preference", preferenceHandler.Create).Methods("POST")
	router.HandleFunc("/preference", preferenceHandler.Update).Methods("PUT")
	router.HandleFunc("/preference/{id}", preferenceHandler.Delete).Methods("DELETE")
	router.HandleFunc("/equipment/all", eqHandler.GetAll).Methods("GET")
	router.HandleFunc("/equipment/tour/{id}", eqHandler.GetByTourId).Methods("GET")
	router.HandleFunc("/equipment", eqHandler.Create).Methods("POST")
	router.HandleFunc("/equipment", eqHandler.Update).Methods("PUT")
	router.HandleFunc("/equipment/{id}", eqHandler.Delete).Methods("DELETE")
	router.HandleFunc("/equipment/{idEq}/{idTour}", eqHandler.Add).Methods("POST")
	router.HandleFunc("/equipment/{idEq}/{idTour}", eqHandler.Remove).Methods("DELETE")

	configureToursHandler(router, tourHandler)
	configureTourExecutionHandler(router, tourExecutionHandler)

	// Set up CORS middleware
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
