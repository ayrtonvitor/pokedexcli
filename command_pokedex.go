package main

import (
	"fmt"
	"strings"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func getPokedexCommand(pokedex *map[string]pokeapi.Pokemon) func(pokeapi.Client, ...string) error {
	return func(pokeapi.Client, ...string) error {
		if len(*pokedex) == 0 {
			return fmt.Errorf("No Pokemon was caught yet\n")
		}
		var pokeBuilder strings.Builder

		for _, poke := range *pokedex {
			pokeBuilder.WriteString(fmt.Sprintf(" - %s\n", poke.Name))
		}
		fmt.Print(pokeBuilder.String())
		return nil
	}
}
