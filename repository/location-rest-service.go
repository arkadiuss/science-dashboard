package repository

import (
	"github.com/arkadiuss/science-dashboard/models"
	"net/http"
	"net/url"
	"encoding/json"
	"strconv"
	"fmt"
)

type LocationRestClient struct {
	client *http.Client
}

type LocationResponse struct {
	Latitude, Longitude string
}

func GetLocationRestClient(client *http.Client) *LocationRestClient {
	return &LocationRestClient { client }
}

func (ls *LocationRestClient) resolvePath(path string) string {
	baseURL, _ := url.Parse("http://free.ipwhois.io")
	endpointURL := &url.URL { Path: path }
	return baseURL.ResolveReference(endpointURL).String()
}

func (ls *LocationRestClient) GetCoordinates(ip string) (models.Location, error) {
	path := ls.resolvePath("json/"+ip)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return models.Location {}, err
	}

	res, err := ls.client.Do(req)
	if err != nil {
		return models.Location {}, err
	}
	defer res.Body.Close()

	var locationResponse LocationResponse
	err = json.NewDecoder(res.Body).Decode(&locationResponse)
	fmt.Println("tes")
	lat, _ := strconv.ParseFloat(locationResponse.Latitude, 8)
	lon, _ := strconv.ParseFloat(locationResponse.Longitude, 8)
	fmt.Println(locationResponse)
	return models.Location { lat, lon }, err
}
