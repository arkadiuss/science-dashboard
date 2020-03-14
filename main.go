package main

import (
	"fmt"
	"github.com/arkadiuss/science-dashboard/service"
	"github.com/arkadiuss/science-dashboard/repository"
	"net/http"
)

func main() {
	var sr repository.ISunRepository
	sr = repository.GetSunrestClient(&http.Client{})

	var ds service.IDashboardService
	ds = service.DashboardService{ sr }
	d,err := ds.GetDashboard()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Sr: %v, Ss: %v", d.Sunrise, d.Sunset)
}
