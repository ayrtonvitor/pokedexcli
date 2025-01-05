package pokeapi

import (
	"encoding/json"
	"log"
)

type ApiConfig struct {
	Url struct {
		PokeApiBaseUrl string            `json:"poke-api-base-url"`
		Path           map[string]string `json:"path"`
	} `json:"url"`
	Client *Client
}

func Setup(conf interface{}, client *Client) ApiConfig {
	var configs ApiConfig
	marshaledConf, _ := json.Marshal(conf.(map[string]interface{}))
	err := json.Unmarshal(marshaledConf, &configs)
	if err != nil {
		log.Fatal("Could not setup the configs for the api.")
	}
	configs.Client = client

	return configs
}
