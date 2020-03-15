package repository

import(
	"net/http"
	"net/url"
	"encoding/json"
	"github.com/arkadiuss/science-dashboard/models"
)

type ISunRepository interface {
	GetSunriseSunset(loc models.Location) (string, string, error)
}

type SunRestClient struct {
	client *http.Client
}

type SunResponse struct {
	Results SunResponseResult
}

type SunResponseResult struct {
	Sunrise, Sunset string
}

func GetSunrestClient(client *http.Client) *SunRestClient {
	return &SunRestClient { client }
}

func (sr SunRestClient) resolvePath(endpoint string) string {
	baseURL, _ := url.Parse("https://api.sunrise-sunset.org")
	endpointURL := &url.URL {Path: endpoint}

	return baseURL.ResolveReference(endpointURL).String()
}

func (sr SunRestClient) GetSunriseSunset(loc models.Location) (string, string, error) {
	path := sr.resolvePath("/json")
	req, err := http.NewRequest("GET", path, nil)
	q := req.URL.Query()
	q.Add("lat", "36.01231")
	q.Add("lng", "-4.52331")
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return "", "", err
	}
	res, err := sr.client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()
	var response SunResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return "", "", err
	}
	return response.Results.Sunrise, response.Results.Sunset, err
}
