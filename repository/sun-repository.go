package repository

import(
	"net/http"
	"time"
	"net/url"
	"encoding/json"
	"github.com/arkadiuss/science-dashboard/models"
)

type ISunRepository interface {
	GetSunriseSunset(loc models.Location) (time.Time, time.Time, error)
}

type SunRestClient struct {
	client *http.Client
}

type SunResponse struct {
	Result SunResponseResult
}

type SunResponseResult struct {
	Sunrise, Sunset time.Time
}

func GetSunrestClient(client *http.Client) *SunRestClient {
	return &SunRestClient { client }
}

func (sr SunRestClient) resolvePath(endpoint string) string {
	baseURL := &url.URL {Path: "https://api.sunrise-sunset.org/json"}
	endpointURL := &url.URL {Path: endpoint}

	return baseURL.ResolveReference(endpointURL).String()
}

func (sr SunRestClient) GetSunriseSunset(loc models.Location) (time.Time, time.Time, error) {
	path := sr.resolvePath("/")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	res, err := sr.client.Do(req)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	
	defer res.Body.Close()
	var response SunResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return response.Result.Sunrise, response.Result.Sunset, err
}
