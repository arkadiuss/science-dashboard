package main

import (
	"fmt"
	"github.com/arkadiuss/science-dashboard/service"
	"github.com/arkadiuss/science-dashboard/repository"
)

func main() {
	var sr repository.ISunRepository
	sr = repository.SunRepository{}

	var ds service.IDashboardService
	ds = service.DashboardService{ sr }
	d := ds.GetDashboard()
	fmt.Printf("Sr: %v, Ss: %v", d.Sunrise, d.Sunset)
}
