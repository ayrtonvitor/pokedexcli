package pokeapi

import (
	"encoding/json"
)

type apiConfig struct {
	Url struct {
		PokeApiBaseUrl string            `json:"poke-api-base-url"`
		Path           map[string]string `json:"path"`
	} `json:"url"`
}

var configs apiConfig

func Setup(conf interface{}) {
	marshaledConf, _ := json.Marshal(conf.(map[string]interface{}))
	conf = json.Unmarshal(marshaledConf, &configs)
}
