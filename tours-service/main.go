package main

import (
	"database-example/handler"
	"database-example/model"
	tour "database-example/proto"
	"database-example/proto/authoring"
	"database-example/proto/authoringKeyPoint"
	"database-example/repo"
	"database-example/saga"
	"database-example/saga/nats"
	"database-example/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	QueueGroup = "order_service"
)

func initDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=super dbname=gorm port=5432 sslmode=disable"
	//dsn := "host=database user=postgres password=super dbname=nzm port=5432 sslmode=disable"
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
	database.AutoMigrate(&model.Review{})
	return database
}

func initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		"nats", "4222",
		"ruser", "T0pS3cr3t", subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		"nats", "4222",
		"ruser", "T0pS3cr3t", subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func initDeleteTourOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.DeleteTourOrchestrator {
	orchestrator, err := service.NewDeleteTourOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func initDeleteTourHandler(service *service.TourService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewDeleteTourCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func initTourService(repo repo.TourRepository, orchestrator *service.DeleteTourOrchestrator) *service.TourService {
	return service.NewTourService(repo, orchestrator)
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	commandPublisher := initPublisher("tour.delete.command")
	replySubscriber := initSubscriber("tour.delete.reply", QueueGroup)
	deleteTourOrchestrator := initDeleteTourOrchestrator(commandPublisher, replySubscriber)

	tourRepo := &repo.TourRepository{DatabaseConnection: database}
	//tourService := &service.TourService{TourRepo: tourRepo}
	tourService := initTourService(*tourRepo, deleteTourOrchestrator)
	keyPointRepo := &repo.KeyPointRepository{DatabaseConnection: database}
	keyPointService := &service.KeyPointService{KeyPointRepo: keyPointRepo}

	commandSubscriber := initSubscriber("tour.delete.command", QueueGroup)
	replyPublisher := initPublisher("tour.delete.reply")
	initDeleteTourHandler(tourService, replyPublisher, commandSubscriber)

	listener, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	tourServer := tour.TourServer{TourService: *tourService}
	tour.RegisterMarketplaceTourServer(grpcServer, tourServer)

	authoringServer := authoring.AuthoringServer1{TourService: *tourService}
	authoring.RegisterAuthoringServer(grpcServer, authoringServer)

	authoringKeyPointServer := authoringKeyPoint.AuthoringKeyPointServer1{KeyPointService: *keyPointService}
	authoringKeyPoint.RegisterAuthoringKeyPointServer(grpcServer, authoringKeyPointServer)

	fmt.Print("server started")
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

}
