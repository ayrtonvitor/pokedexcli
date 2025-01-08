package main

import (
	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(pokeapi.Client, ...string) error
}

func setupCommands() map[string]*cliCommand {
	commandMap, commandMapb := getMapCommands()
	commandInspect, pokedex := getInspectCommand()
	var commandsDict = map[string]*cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex\n",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message\n",
			callback:    commandDummy,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of locations in the map\n",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of locations in the map\n",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Lists all the pokemon in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Saves a pokemon to the pokedex",
			callback:    commandCatch(pokedex),
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon in the pokedex",
			callback:    commandInspect,
		},
	}
	commandsDict["help"].callback = getHelpCommand(commandsDict)

	return commandsDict
}

func commandDummy(pokeapi.Client, ...string) error {
	return nil
}
