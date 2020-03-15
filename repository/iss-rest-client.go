package repository

import (
	"github.com/arkadiuss/science-dashboard/models"
	"time"
	"net/http"
	"net/url"
	"encoding/json"
	"fmt"
	"strconv"
)

type ISSRestClient struct {
	client *http.Client
}

type currentLocationResponse struct {
	Iss_position IssPosition
}

type IssPosition struct {
	Latitude, Longitude string
}

type passResponse struct {
	Response []issPass
}

type issPass struct {
	Risetime int64
}

func GetISSRestClient(httpClient *http.Client) *ISSRestClient {
	return &ISSRestClient { httpClient }
}

func (ir *ISSRestClient) resolvePath(path string) string {
	baseURL, _ := url.Parse("http://api.open-notify.org")
	pathURL := &url.URL { Path: path }
	return baseURL.ResolveReference(pathURL).String()
}

func (ir *ISSRestClient) GetCurrentLocation() (models.Location, error) {
	path := ir.resolvePath("iss-now.json")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return models.Location {}, err
	}

	res, err := ir.client.Do(req)
	if err != nil {
		return models.Location {}, err
	}
	defer res.Body.Close()

	var response currentLocationResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return models.Location {}, err
	}
	lat, _ := strconv.ParseFloat(response.Iss_position.Latitude, 8)
	lon, _ := strconv.ParseFloat(response.Iss_position.Longitude, 8)
	return models.Location { lat, lon }, err
}

func (ir *ISSRestClient) GetNextPasses(location models.Location) ([]time.Time, error) {
	path := ir.resolvePath("iss-pass.json")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err 
	}
	q := req.URL.Query()
	q.Add("lat", fmt.Sprintf("%v", location.Lat))
	q.Add("lon", fmt.Sprintf("%v", location.Lon))
	req.URL.RawQuery = q.Encode()

	res, err := ir.client.Do(req)
	if err != nil {
		return nil, err 
	}
	defer res.Body.Close()

	var response passResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	var times []time.Time
	for _, p := range response.Response {
		times = append(times, time.Unix(p.Risetime, 0))
	}
	return times, nil
}
