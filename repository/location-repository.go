package repository

import (
	"github.com/arkadiuss/science-dashboard/models"
)

type ILocationRepository interface {
	GetCoordinates(ip string) (models.Location, error)
}
