package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func getMapCommands() (func(pokeapi.ApiConfig) error, func(pokeapi.ApiConfig) error) {
	currentPage := 0

	commandMap := func(config pokeapi.ApiConfig) error {
		locations, err := pokeapi.GetNextLocations(&currentPage, config)
		if err != nil {
			return err
		}
		for _, loc := range locations {
			fmt.Println(loc)
		}
		return nil
	}

	commandMapb := func(config pokeapi.ApiConfig) error {
		locations, err := pokeapi.GetPrevLocations(&currentPage, config)
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
