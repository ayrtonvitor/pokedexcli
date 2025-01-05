package pokeapi

import (
	"fmt"
	"net/url"
)

func buildURL(baseURL, path string, params [][2]string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("Could not parse base url:\n%w\n", err)
	}

	base.Path = base.Path + path

	query := base.Query()
	for _, pair := range params {
		query.Add(pair[0], pair[1])
	}
	base.RawQuery = query.Encode()

	return base.String(), nil
}
