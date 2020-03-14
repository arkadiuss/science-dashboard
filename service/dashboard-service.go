package service

import (
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/models"
)

type IDashboardService interface {
	GetDashboard() (models.Dashboard, error)
}

type DashboardService struct {
	SunRepository repository.ISunRepository
}

func (ds DashboardService) GetDashboard() (models.Dashboard, error) {
	d := models.Dashboard{}
	var err error
	d.Sunrise, d.Sunset, err = ds.SunRepository.GetSunriseSunset(models.Location {})
	if err != nil {
		return models.Dashboard{}, err
	}
	return d, nil
}

