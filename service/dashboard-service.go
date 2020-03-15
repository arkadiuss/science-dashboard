package service

import (
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/models"
	"time"
)

type IDashboardService interface {
	GetDashboard(string) (models.Dashboard, error)
}

type DashboardService struct {
	sunRepository repository.ISunRepository
	locationRepository repository.ILocationRepository 
	iissRepository repository.IISSRepository
}

func GetDashboardService(sr repository.ISunRepository, lr repository.ILocationRepository, ir repository.IISSRepository) *DashboardService {
	return &DashboardService { sr, lr, ir }
}

func (ds *DashboardService) GetDashboard(location string) (models.Dashboard, error) {
	d := models.Dashboard{}
	var err error
	d.Location, err = ds.locationRepository.GetCoordinates(location)
	if err != nil {
		return models.Dashboard{}, err
	}
	d.Sunrise, d.Sunset, err = ds.sunRepository.GetSunriseSunset(d.Location)
	if err != nil {
		return models.Dashboard{}, err
	}
	
	d.ISSLocation, err = ds.iissRepository.GetCurrentLocation()
	if err != nil {
		return models.Dashboard{}, err
	}
	passes, err := ds.iissRepository.GetNextPasses(d.Location)
	if err != nil {
		return models.Dashboard{}, err
	}
	d.ISSNextPass = int(passes[0].Sub(time.Now())/1e9/60)
	return d, nil
}

