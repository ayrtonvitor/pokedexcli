package main

import (
	"fmt"
	"math/rand"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func commandCatch(pokedex *map[string]pokeapi.Pokemon) func(pokeapi.Client, ...string) error {
	rng := rand.New(rand.NewSource(42))

	innerCommandcatch := func(apiClient pokeapi.Client, params ...string) error {
		if len(params) == 0 {
			return fmt.Errorf("Catch requires a Pokemon as parameter\n")
		}
		poke := params[0]

		pokemon, err := apiClient.GetPokemon(poke)
		if err != nil {
			return fmt.Errorf("Could not find pokemon: %w\n", err)
		}
		uncatchableExp, err := apiClient.GetUncatchableBaseExp()
		if err != nil || uncatchableExp < 1 {
			return fmt.Errorf("Could compare pokemon's XP\n")
		}

		fmt.Printf("Throwing a Pokeball at %s...\n", poke)
		if tryCatch(pokemon, uncatchableExp, rng) {
			fmt.Printf("%s was caught!\n", pokemon.Name)
			(*pokedex)[pokemon.Name] = pokemon
			return nil
		}
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
	return innerCommandcatch
}

func tryCatch(poke pokeapi.Pokemon, uncatchableExp int, rng *rand.Rand) bool {
	n := rng.Intn(uncatchableExp - 1)
	return n > poke.BaseExp
}
