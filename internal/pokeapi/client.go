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
		log.Fatalf("Could not setup the configs for the api:\n%v", err.Error())
	}
	timeout := config.Timeout
	if timeout == 0 {
		fmt.Println("Request timeout not specified in the configuration file. Using 10 seconds.")
		timeout = 10
	}

	if config.Cache == (pokecache.CacheConfig{}) {
		log.Println("Could not load cache configs for client. Using defaults")
		config.Cache = pokecache.CacheConfig{
			Interval: 60 * 1000,
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
