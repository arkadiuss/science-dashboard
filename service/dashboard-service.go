package service

import (
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/models"
)

type IDashboardService interface {
	GetDashboard(string) (models.Dashboard, error)
}

type DashboardService struct {
	SunRepository repository.ISunRepository
	LocationRepository repository.ILocationRepository 
}

func (ds *DashboardService) GetDashboard(location string) (models.Dashboard, error) {
	d := models.Dashboard{}
	var err error
	d.Location, err = ds.LocationRepository.GetCoordinates(location)
	if err != nil {
		return models.Dashboard{}, err
	}
	d.Sunrise, d.Sunset, err = ds.SunRepository.GetSunriseSunset(d.Location)
	if err != nil {
		return models.Dashboard{}, err
	}
	return d, nil
}

