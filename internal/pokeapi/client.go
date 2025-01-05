package pokeapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ayrtonvitor/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	config     apiConfig
	cache      *pokecache.Cache
}

func NewClient(conf map[string]interface{}) Client {
	var config apiConfig
	marshaledConf, _ := json.Marshal(conf)
	err := json.Unmarshal(marshaledConf, &config)
	if err != nil {
		log.Fatal("Could not setup the configs for the api.")
	}
	timeout := config.Timeout
	if timeout == 0 {
		fmt.Println("Request timeout not specified in the configuration file. Using 10 seconds.")
		timeout = 10
	}

	if config.Cache == nil {
		log.Println("Could not load cache configs. Using defaults")
		config.Cache = map[string]interface{}{
			"interval": 10,
		}
	}
	cache := pokecache.NewCache(config.Cache)

	return Client{
		httpClient: http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
		config: config,
		cache:  cache,
	}
}
