package repository

import (
	"github.com/arkadiuss/science-dashboard/models"
	"time"
)

type IISSRepository interface {
	GetCurrentLocation() (models.Location, error)
	GetNextPasses(models.Location) ([]time.Time, error)
}
