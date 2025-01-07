package pokeapi

import "github.com/ayrtonvitor/pokedexcli/internal/pokecache"

type apiConfig struct {
	Url struct {
		PokeApiBaseUrl string            `json:"poke-api-base-url"`
		Path           map[string]string `json:"path"`
	} `json:"url"`
	Timeout int                   `json:"timeout"`
	Cache   pokecache.CacheConfig `json:"cache"`
}
