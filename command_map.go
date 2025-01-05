package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	locations, err := pokeapi.GetLocations()
	if err != nil {
		return err
	}
	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}
