package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	config     apiConfig
}

func NewClient(timeout time.Duration, conf interface{}) Client {
	var config apiConfig
	marshaledConf, _ := json.Marshal(conf.(map[string]interface{}))
	err := json.Unmarshal(marshaledConf, &config)
	if err != nil {
		log.Fatal("Could not setup the configs for the api.")
	}

	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		config: config,
	}
}
