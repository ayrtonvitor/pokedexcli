package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetNextLocations(currentPage *int, config ApiConfig) ([]string, error) {
	*currentPage += 1
	return GetLocations(*currentPage, config)
}

func GetPrevLocations(currentPage *int, config ApiConfig) ([]string, error) {
	if *currentPage > 1 {
		*currentPage -= 1
		return GetLocations(*currentPage, config)
	}
	return []string{}, nil
}

func GetLocations(currentPage int, config ApiConfig) ([]string, error) {
	query := make(map[string]string)
	query["offset"] = strconv.Itoa((currentPage - 1) * 20)
	query["limit"] = "20"
	url, err := buildURL(
		config.Url.PokeApiBaseUrl, config.Url.Path["location"], query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not build request:\n%w\n", err)
	}

	res, err := config.Client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request error:\n%w\n", err)
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode < 300 || res.StatusCode == 304) {
		return nil, fmt.Errorf("Non-OK HTTP status: %s\n", res.Status)
	}

	var data locationRes
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Could not decode response\n")
	}

	locations := []string{}
	for _, loc := range data.Results {
		locations = append(locations, loc.Name)
	}
	return locations, nil
}

type locationRes struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}
