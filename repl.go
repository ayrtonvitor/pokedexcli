package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func run(commands map[string]*cliCommand, apiClient pokeapi.Client) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command := input[0]
		tryExec(command, commands, apiClient)
	}
}

func tryExec(input string, commands map[string]*cliCommand, apiClient pokeapi.Client) {
	reg, ok := commands[input]
	if !ok {
		fmt.Println("Unknown command")
		return
	}

	err := reg.callback(apiClient)
	if err != nil {
		fmt.Printf("%w", err.Error())
	}
}
