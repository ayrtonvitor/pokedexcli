package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func setupCommands() {
	var commandsDict = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex\n",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message\n",
			callback:    commandHelp,
		},
	}
	commands = commandsDict
}
