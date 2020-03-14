package models

import(
	"time"
)

type Dashboard struct {
	Sunrise, Sunset time.Time
}
