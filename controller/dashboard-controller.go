package controller

import (
	"fmt"
	"github.com/arkadiuss/science-dashboard/service"
	"net/http"
)

var dashboardService service.IDashboardService

func handler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Test!")
}

func SetupDashboardController(ds service.IDashboardService) {
	dashboardService = ds
	http.HandleFunc("/", handler)
}

