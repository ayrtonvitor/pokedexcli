package main

import (
	"fmt"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func getHelpCommand(commands map[string]*cliCommand) func(pokeapi.ApiConfig) error {
	innerHelp := func(_ pokeapi.ApiConfig) error {
		fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
		for _, v := range commands {
			fmt.Printf("%s: %s", v.name, v.description)
		}
		return nil
	}

	return innerHelp
}
