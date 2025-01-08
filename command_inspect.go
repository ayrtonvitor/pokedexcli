package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func getInspectCommand() (func(pokeapi.Client, ...string) error, *map[string]pokeapi.Pokemon) {
	pokedex := make(map[string]pokeapi.Pokemon)
	getMap := func(_ pokeapi.Client, params ...string) error {
		if len(params) == 0 {
			return fmt.Errorf("Catch requires a Pokemon as parameter\n")
		}
		pokeName := params[0]

		poke, ok := pokedex[pokeName]
		if !ok {
			return fmt.Errorf("You did not catch such Pokemon.\n")
		}
		fmt.Println(poke)
		return nil
	}
	return getMap, &pokedex
}
