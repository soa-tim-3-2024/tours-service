package authoring

import (
	context "context"
	"database-example/model"
	"database-example/service"
)

type AuthoringServer1 struct {
	TourService service.TourService
	UnimplementedAuthoringServer
}

func (ts AuthoringServer1) AddTour(context context.Context, newTour *TourCreate) (*TourResponseAuthor, error) {
	tour := model.Tour{AuthorId: int(newTour.AuthorId), Name: newTour.Name, Description: newTour.Description, Difficulty: newTour.Difficulty, Category: int(newTour.Category), Tags: newTour.Tags,
		Price: newTour.Price, TourStatus: model.Status(newTour.Status)}
	err := ts.TourService.Create(&tour)
	if err != nil {
		return nil, err
	}
	var toursResponse = TourResponseAuthor{Id: int64(tour.ID), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponseAuthor_TourStatus(tour.TourStatus)}
	return &(toursResponse), nil
}

func (ts AuthoringServer1) UpdateTour(context context.Context, newTour *TourUpdate) (*TourResponseAuthor, error) {
	tour := model.Tour{ID: int(newTour.Id), AuthorId: int(newTour.AuthorId), Name: newTour.Name, Description: newTour.Description, Difficulty: newTour.Difficulty, Category: int(newTour.Category), Tags: newTour.Tags,
		Price: newTour.Price, TourStatus: model.Status(newTour.Status)}
	err := ts.TourService.Update(&tour)
	if err != nil {
		return nil, err
	}
	var toursResponse = TourResponseAuthor{Id: int64(tour.ID), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponseAuthor_TourStatus(tour.TourStatus)}
	return &(toursResponse), nil
}

func (ts AuthoringServer1) PublishTour(context context.Context, newTour *TourUpdate) (*PublishResponse, error) {
	tour := model.Tour{ID: int(newTour.Id), AuthorId: int(newTour.AuthorId), Name: newTour.Name, Description: newTour.Description, Difficulty: newTour.Difficulty, Category: int(newTour.Category), Tags: newTour.Tags,
		Price: newTour.Price, TourStatus: model.Status(newTour.Status)}
	err := ts.TourService.PublishTour(&tour)
	if err != nil {
		return nil, err
	}
	var response = PublishResponse{Response: 1}
	return &(response), nil
}

func (ts AuthoringServer1) ArchiveTour(context context.Context, newTour *TourUpdate) (*PublishResponse, error) {
	tour := model.Tour{ID: int(newTour.Id), AuthorId: int(newTour.AuthorId), Name: newTour.Name, Description: newTour.Description, Difficulty: newTour.Difficulty, Category: int(newTour.Category), Tags: newTour.Tags,
		Price: newTour.Price, TourStatus: model.Status(newTour.Status)}
	err := ts.TourService.ArchiveTour(&tour)
	if err != nil {
		return nil, err
	}
	var response = PublishResponse{Response: 1}
	return &(response), nil
}
