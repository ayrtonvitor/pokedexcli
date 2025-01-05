package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func getMapCommands() (func(pokeapi.Client) error, func(pokeapi.Client) error) {
	currentPage := 0

	commandMap := func(apiClient pokeapi.Client) error {
		locations, err := apiClient.GetNextLocations(&currentPage)
		if err != nil {
			return err
		}
		for _, loc := range locations {
			fmt.Println(loc)
		}
		return nil
	}

	commandMapb := func(apiClient pokeapi.Client) error {
		locations, err := apiClient.GetPrevLocations(&currentPage)
		if err != nil {
			return err
		}
		if locations != nil && len(locations) == 0 {
			fmt.Println("You are on the first page")
		}
		for _, loc := range locations {
			fmt.Println(loc)
		}
		return nil
	}

	return commandMap, commandMapb
}
