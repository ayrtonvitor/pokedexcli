package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func commandMap() error {
	locations, err := pokeapi.GetNextLocations()
	if err != nil {
		return err
	}
	for _, loc := range locations {
		fmt.Println(loc)
	}
	return nil
}

func commandMapb() error {
	locations, err := pokeapi.GetPrevLocations()
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
