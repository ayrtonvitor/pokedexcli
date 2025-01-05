package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var locationOffset int

func GetLocations() ([]string, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	query := make(map[string]string)
	query["offset"] = strconv.Itoa(locationOffset)
	query["limit"] = "20"
	url, err := buildURL(
		configs.Url.PokeApiBaseUrl, configs.Url.Path["location"], query)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not build request:\n%w\n", err)
	}

	res, err := client.Do(req)
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

	locationOffset += 20
	locations := []string{}
	for _, loc := range data.Results {
		locations = append(locations, loc.Name)
	}
	return locations, nil
}

type locationRes struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string `json:"name"`
		Url  string
	} `json:"results"`
}
