package authoringKeyPoint

import (
	context "context"
	"database-example/model"
	"database-example/service"
	"strconv"
)

type AuthoringKeyPointServer1 struct {
	KeyPointService service.KeyPointService
	UnimplementedAuthoringKeyPointServer
}

// KeyPoint u ovom fajlu nije KeyPoint iz modela, nego iz proto fajla :(
func (ts AuthoringKeyPointServer1) CreateKeyPoint(context context.Context, newKeyPoint *KeyPointCreate) (*KeyPoint, error) {
	keyPoint := model.KeyPoint{TourId: int(newKeyPoint.TourId), Name: newKeyPoint.Name, Description: newKeyPoint.Description, Longitude: newKeyPoint.Longitude,
		Latitude: newKeyPoint.Latitude, LocationAddress: newKeyPoint.LocationAddress, Order: int(newKeyPoint.Order), ImagePath: newKeyPoint.ImagePath, IsEncounterRequired: newKeyPoint.IsEncounterRequired, HasEncounter: newKeyPoint.HasEncounter}
	err := ts.KeyPointService.Create(&keyPoint)
	if err != nil {
		return nil, err
	}
	var keyPointResponse = KeyPoint{Name: keyPoint.Name, Description: keyPoint.Description, ImagePath: keyPoint.ImagePath, TourId: int64(keyPoint.TourId),
		Longitude: keyPoint.Longitude, Latitude: keyPoint.Latitude}
	return &(keyPointResponse), nil
}

func (ts AuthoringKeyPointServer1) UpdateKeyPoint(context context.Context, newKeyPoint *KeyPointUpdate) (*KeyPoint, error) {
	keyPoint := model.KeyPoint{ID: int(newKeyPoint.Id), TourId: int(newKeyPoint.TourId), Name: newKeyPoint.Name, Description: newKeyPoint.Description, Longitude: newKeyPoint.Longitude,
		Latitude: newKeyPoint.Latitude, LocationAddress: newKeyPoint.LocationAddress, Order: int(newKeyPoint.Order), ImagePath: newKeyPoint.ImagePath}
	err := ts.KeyPointService.Update(&keyPoint)
	if err != nil {
		return nil, err
	}
	var keyPointResponse = KeyPoint{Id: int64(keyPoint.ID), Name: keyPoint.Name, Description: keyPoint.Description, ImagePath: keyPoint.ImagePath, TourId: int64(keyPoint.TourId),
		Longitude: keyPoint.Longitude, Latitude: keyPoint.Latitude}
	return &(keyPointResponse), nil
}

func (ts AuthoringKeyPointServer1) DeleteKeyPoint(context context.Context, request *KeyPointId) (*KeyPoint, error) {
	id := strconv.Itoa(int(request.KeyPointId))
	keyPoint, err := ts.KeyPointService.DeleteById(id)
	if err != nil {
		return nil, err
	}
	var keyPointResponse = KeyPoint{Id: int64(keyPoint.ID), Name: keyPoint.Name, Description: keyPoint.Description, ImagePath: keyPoint.ImagePath, TourId: int64(keyPoint.TourId),
		Longitude: keyPoint.Longitude, Latitude: keyPoint.Latitude}
	return &(keyPointResponse), nil
}
