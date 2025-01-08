package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemon struct {
	Name    string `json:"name"`
	BaseExp int    `json:"base_experience"`
}

func (c *Client) GetPokemon(pokeName string) (Pokemon, error) {
	url, err := buildURL(
		c.config.Url.PokeApiBaseUrl, c.config.Url.Path["pokemon"]+"/"+pokeName, nil)
	if err != nil {
		return Pokemon{}, err
	}

	var poke Pokemon
	if item, ok := c.cache.Get("url"); ok {
		err := json.Unmarshal(item, &poke)
		if err == nil {
			return poke, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Could not build request:\n%w\n", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("request error:\n%w\n", err)
	}
	if !(res.StatusCode >= 200 && res.StatusCode < 300) {
		return Pokemon{}, fmt.Errorf("Non-OK HTTP status: %s\n", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&poke); err != nil {
		return Pokemon{}, fmt.Errorf("Could not decode response: %w\n", err)
	}

	toCache, err := json.Marshal(poke)
	if err == nil {
		c.cache.Add(url, toCache)
	}

	return poke, nil
}

func (c *Client) GetUncatchableBaseExp() (int, error) {
	firstUncatchable := "mewtwo"
	poke, err := c.GetPokemon(firstUncatchable)
	if err != nil {
		return 0, err
	}
	return poke.BaseExp, nil
}
