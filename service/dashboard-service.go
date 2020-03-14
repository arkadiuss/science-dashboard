package service

import "github.com/arkadiuss/science-dashboard/repository"
import "github.com/arkadiuss/science-dashboard/models"

type IDashboardService interface {
	GetDashboard() models.Dashboard
}

type DashboardService struct {
	SunRepository repository.ISunRepository
}

func (ds DashboardService) GetDashboard() models.Dashboard {
	d := models.Dashboard{}
	d.Sunrise, d.Sunset = ds.SunRepository.GetSunriseSunset()
	return d
}

