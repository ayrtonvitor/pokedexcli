package pokeapi

import (
	"fmt"
	"net/url"
)

func buildURL(baseURL, path string, params map[string]string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("Could not parse base url:\n%w\n", err)
	}

	base.Path = base.Path + path

	query := base.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	base.RawQuery = query.Encode()

	return base.String(), nil
}
