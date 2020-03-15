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
	var il repository.IISSRepository
	il = repository.GetISSRestClient(httpClient)
	var cr repository.ICoronavirusRepository
	cr = repository.GetCoronavirusRestClient(httpClient)
	
	var ds service.IDashboardService
	ds = service.GetDashboardService(sr, sl, il, cr)

	controller.SetupDashboardController(ds)

	log.Println("Starting server on 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
