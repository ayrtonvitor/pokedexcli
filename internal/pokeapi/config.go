package pokeapi

type apiConfig struct {
	Url struct {
		PokeApiBaseUrl string            `json:"poke-api-base-url"`
		Path           map[string]string `json:"path"`
	} `json:"url"`
}
