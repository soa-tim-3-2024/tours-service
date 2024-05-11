package tour

import (
	"context"
	"database-example/service"
	"fmt"
)

type TourServer struct {
	TourService service.TourService
	UnimplementedMarketplaceTourServer
}

func (ts TourServer) GetPublishedTours(context context.Context, page *Page) (*TourResponse, error) {
	fmt.Println("Pocetak, krenulo je")
	tours, err := ts.TourService.GetPublishedTours()
	if err != nil {
		return nil, err
	}
	if len(*tours) == 0 {
		return nil, nil
	}
	var toursResponse []TourResponse
	for _, tour := range *tours {
		toursResponse = append(toursResponse, TourResponse{Id: int64(tour.ID), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponse_TourStatus(tour.TourStatus)})
	}
	return &(toursResponse[0]), nil
}
