package repository

import (
	"github.com/arkadiuss/science-dashboard/models"
)

type ISunRepository interface {
	GetSunriseSunset(loc models.Location) (string, string, error)
}


