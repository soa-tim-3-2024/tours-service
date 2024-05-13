package tour

import (
	"context"
	"database-example/service"
	"fmt"
	"strconv"
)

type TourServer struct {
	TourService service.TourService
	UnimplementedMarketplaceTourServer
}

func (ts TourServer) GetPublishedTours(context context.Context, page *Page) (*TourResponseList, error) {
	fmt.Println("Pocetak, krenulo je")
	tours, err := ts.TourService.GetPublishedTours()
	if err != nil {
		return nil, err
	}
	if len(*tours) == 0 {
		return nil, nil
	}
	var toursResponse []*TourResponse
	for _, tour := range *tours {
		keyPoints := []*KeyPointResponse{}
		for _, kp := range tour.KeyPoints {
			keyPoints = append(keyPoints, &KeyPointResponse{Id: int64(kp.ID), Longitude: kp.Longitude, Latitude: kp.Latitude, LocationAddress: kp.LocationAddress,
				Name: kp.Name, Description: kp.Description, TourId: int64(kp.TourId), ImagePath: kp.ImagePath, Order: int64(kp.Order)})
		}

		toursResponse = append(toursResponse, &TourResponse{Id: int64(tour.ID), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponse_TourStatus(tour.TourStatus),
			Difficulty: tour.Difficulty, Category: TourResponse_TourCategory(tour.Category), Price: tour.Price, Distance: tour.Distance,
			KeyPoints: keyPoints, AuthorId: int64(tour.AuthorId)})
	}
	var toursResponses TourResponseList
	toursResponses.TourResponses = toursResponse
	return &toursResponses, nil
}

func (ts TourServer) GetAuthorTours(context context.Context, authorId *AuthorId) (*TourResponseList, error) {
	tours, err := ts.TourService.GetAuthorTours(int(authorId.AuthorId))
	if err != nil {
		return nil, err
	}
	if len(tours) == 0 {
		return nil, nil
	}
	var toursResponse []*TourResponse
	for _, tour := range tours {
		toursResponse = append(toursResponse, &TourResponse{Id: int64(tour.ID), AuthorId: int64(tour.AuthorId), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponse_TourStatus(tour.TourStatus),
			Difficulty: tour.Difficulty, Category: TourResponse_TourCategory(tour.Category), Price: tour.Price, Distance: tour.Distance})
	}
	var toursResponses TourResponseList
	toursResponses.TourResponses = toursResponse
	return &toursResponses, nil
}

func (ts TourServer) GetTour(context context.Context, tourId *TourId) (*TourResponse, error) {
	id := strconv.Itoa(int(tourId.TourId))
	tour, err := ts.TourService.FindTour(id)
	if err != nil {
		return nil, err
	}
	var toursResponse *TourResponse
	keyPoints := []*KeyPointResponse{}
	durations := []*TourDuration{}
	for _, kp := range tour.KeyPoints {
		keyPoints = append(keyPoints, &KeyPointResponse{Id: int64(kp.ID), Longitude: kp.Longitude, Latitude: kp.Latitude, LocationAddress: kp.LocationAddress,
			Name: kp.Name, Description: kp.Description, TourId: int64(kp.TourId), ImagePath: kp.ImagePath, Order: int64(kp.Order)})
	}
	for _, duration := range tour.Durations {
		durations = append(durations, &TourDuration{Duration: int32(duration.Duration), TransportType: TourDuration_TransportType(duration.TransportType)})
	}
	toursResponse = &TourResponse{Id: int64(tour.ID), Name: tour.Name, Description: tour.Description, Tags: tour.Tags, Status: TourResponse_TourStatus(tour.TourStatus),
		Difficulty: tour.Difficulty, Category: TourResponse_TourCategory(tour.Category), Price: tour.Price, Distance: tour.Distance, AuthorId: int64(tour.AuthorId),
		KeyPoints: keyPoints /*, Durations: durations*/}

	return toursResponse, nil
}
