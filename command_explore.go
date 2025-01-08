package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func commandExplore(apiClient pokeapi.Client, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Explore requires a location name as parameter\n")
	}
	loc := params[0]
	if !apiClient.SearchLocation(loc) {
		return fmt.Errorf("No such loction %s\n", loc)
	}

	fmt.Printf("Exploring %s\n", loc)
	pokemons, err := apiClient.ExploretLocation(loc)
	if err != nil {
		return fmt.Errorf("commandExplore: %w", err)
	}

	fmt.Println("Found Pokemon:", loc)
	for _, poke := range pokemons {
		fmt.Printf(" - %s\n", poke)
	}
	return nil
}
