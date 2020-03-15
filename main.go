package main

import (
	"github.com/arkadiuss/science-dashboard/service"
	"github.com/arkadiuss/science-dashboard/repository"
	"github.com/arkadiuss/science-dashboard/controller"
	"net/http"
	"log"
)

func main() {
	httpClient := &http.Client {}

	var sr repository.ISunRepository
	sr = repository.GetSunrestClient(httpClient)
	var sl repository.ILocationRepository
	sl = repository.GetLocationRestClient(httpClient)

	var ds service.IDashboardService
	ds = &service.DashboardService{ sr, sl }

	controller.SetupDashboardController(ds)

	log.Println("Starting server on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
