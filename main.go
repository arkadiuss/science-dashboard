package main

import (
	"fmt"
	"github.com/arkadiuss/science-dashboard/service"
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/controller"
	"net/http"
)

func main() {
	restClient := &http.Client {}
	var sr repository.ISunRepository
	sr = repository.GetSunrestClient(restClient)
	var sl repository.ILocationRepository
	sl = repository.GetLocationRestClient(restClient)

	var ds service.IDashboardService
	ds = &service.DashboardService{ sr, sl }
	d,err := ds.GetDashboard("83.30.81.55")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Location: (%v, %v) \n", d.Location.Lat, d.Location.Lon)
	fmt.Printf("Sunrise: %v, Sunset: %v \n", d.Sunrise, d.Sunset)

	controller.SetupDashboardController(ds)
	http.ListenAndServe(":8080", nil)
}
