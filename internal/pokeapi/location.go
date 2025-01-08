package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (c *Client) GetNextLocations(currentPage *int) ([]string, error) {
	*currentPage += 1
	return c.getLocations(*currentPage)
}

func (c *Client) GetPrevLocations(currentPage *int) ([]string, error) {
	if *currentPage > 1 {
		*currentPage -= 1
	}
	return c.getLocations(*currentPage)
}

func (c *Client) getLocations(currentPage int) ([]string, error) {
	query := [][2]string{
		[2]string{"offset", strconv.Itoa((currentPage - 1) * 20)},
		[2]string{"limit", "20"},
	}
	url, err := buildURL(
		c.config.Url.PokeApiBaseUrl, c.config.Url.Path["location"], query)
	if err != nil {
		return nil, err
	}

	if item, ok := c.cache.Get("url"); ok {
		var values []string
		err := json.Unmarshal(item, &values)
		if err == nil {
			return values, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not build request:\n%v\n", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request error:\n%v\n", err)
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode < 300 || res.StatusCode == 304) {
		return nil, fmt.Errorf("Non-OK HTTP status: %s\n", res.Status)
	}

	var data locationRes
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Could not decode response: %v\n", err.Error())
	}

	if len(data.Results) == 0 {
		return nil, emptySuccessfulRespError{}
	}

	locations := []string{}
	for _, loc := range data.Results {
		locations = append(locations, loc.Name)
	}
	toCache, err := json.Marshal(locations)
	if err == nil {
		c.cache.Add(url, toCache)
	}

	return locations, nil
}

func (c *Client) SearchLocation(sup string) bool {
	var err error = nil

	for i := 1; !errors.Is(emptySuccessfulRespError{}, err); i++ {
		locs, getErr := c.getLocations(i)
		err = getErr
		for _, loc := range locs {
			if loc == sup {
				return true
			}
		}
	}
	return false
}

func (c *Client) ExploretLocation(name string) ([]string, error) {
	url, err := buildURL(
		c.config.Url.PokeApiBaseUrl, c.config.Url.Path["location"]+"/"+name, nil)
	if err != nil {
		return nil, err
	}

	if item, ok := c.cache.Get("url"); ok {
		var values []string
		err := json.Unmarshal(item, &values)
		if err == nil {
			return values, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Could not build request:\n%v\n", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request error:\n%v\n", err)
	}
	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode < 300 || res.StatusCode == 304) {
		return nil, fmt.Errorf("Non-OK HTTP status: %s\n", res.Status)
	}

	var data locationPokeEncountersRes
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("Could not decode response: %v\n", err.Error())
	}

	if len(data.PokemonEncounters) == 0 {
		return nil, emptySuccessfulRespError{}
	}

	pokemons := []string{}
	for _, poke := range data.PokemonEncounters {
		pokemons = append(pokemons, poke.Pokemon.Name)
	}
	toCache, err := json.Marshal(pokemons)
	if err == nil {
		c.cache.Add(url, toCache)
	}

	return pokemons, nil
}

type locationRes struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type locationPokeEncountersRes struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		}
	} `json:"pokemon_encounters"`
}

type emptySuccessfulRespError struct {
	reqType string
}

func (e emptySuccessfulRespError) Error() string {
	switch e.reqType {
	case "location":
		return "No more locations\n"
	case "pokeEncounters":
		return "Could not find pokemon\n"
	default:
		return "Unknown request error\n"
	}
}
