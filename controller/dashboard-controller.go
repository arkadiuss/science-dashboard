package controller

import (
	"fmt"
	"github.com/arkadiuss/science-dashboard/service"
	"net/http"
)

var dashboardService service.IDashboardService

func welcomeHandler(writer http.ResponseWriter, req *http.Request) {
	template := `
		<html>
			<head></head>
			<body>
				<h1>Welcome!</h1>
				<form action="/dashboard" method="GET">
					<input type="text" id="ip" name="ip" placeholder="IP address"><br>
					<input type="submit" value="Give me a dashboard! ">
				</form>
			<body>
		</html>
	`
	fmt.Fprintf(writer, template)
}

func dashboardHandler(writer http.ResponseWriter, req *http.Request) {
	template := `
		<html>
			<head></head>
			<body>
				<h1>Dashoboard for %v </h1>
				<h2>Information</h2>
				<table>
					<tr>
						<td>Location</td>
						<td>(%v, %v)</td>
					</tr>
					<tr>
						<td>Sunrise</td>
						<td>%v</td>
					</tr>
					<tr>
						<td>Sunset</td>
						<td>%v</td>
					</tr>
				</table>
				<h2>International Space Station</h2>
				<table>
					<tr>
						<td>Current location</td>
						<td>(%v, %v)</td>
					</tr>
					<tr>
						<td>Next pass over you head in</td>
						<td>%v min</td>
					</tr>
				</table>
			</body>
		<html>
	`
	ip := req.URL.Query()["ip"][0]
	dashboard, err := dashboardService.GetDashboard(ip)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(writer, template,
		ip,
		dashboard.Location.Lat,
		dashboard.Location.Lon,
		dashboard.Sunrise,
		dashboard.Sunset,
		dashboard.ISSLocation.Lat,
		dashboard.ISSLocation.Lon,
		dashboard.ISSNextPass)
}

func SetupDashboardController(ds service.IDashboardService) {
	dashboardService = ds
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
}

