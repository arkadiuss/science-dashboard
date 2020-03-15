package repository

import (
	"net/http"
	"net/url"
	"encoding/json"
)

type CoronavirusRestClient struct {
	client *http.Client
}

type statsResponse struct {
	Cases, Deaths, Recovered int
}

func GetCoronavirusRestClient(client *http.Client) *CoronavirusRestClient {
	return &CoronavirusRestClient { client }
}

func (ls *CoronavirusRestClient) resolvePath(path string) string {
	baseURL, _ := url.Parse("https://corona.lmao.ninja")
	endpointURL := &url.URL { Path: path }
	return baseURL.ResolveReference(endpointURL).String()
}

func (ls *CoronavirusRestClient) GetGlobalStats() (int, int, int, error) {
	path := ls.resolvePath("all")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return 0, 0, 0, err
	}

	res, err := ls.client.Do(req)
	if err != nil {
		return 0, 0, 0, err
	}
	defer res.Body.Close()

	var response statsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	return response.Cases, response.Deaths, response.Recovered, err
}
