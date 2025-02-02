package main

import (
	"fmt"
	"os"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func commandExit(_ pokeapi.Client, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
