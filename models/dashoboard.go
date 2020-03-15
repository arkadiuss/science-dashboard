package models


type Dashboard struct {
	Location Location
	Sunrise, Sunset string
	ISSLocation Location
	ISSNextPass int
}
