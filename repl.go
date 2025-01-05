package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func run() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command := input[0]
		tryExec(command)
	}
}

func tryExec(input string) {
	reg, ok := commands[input]
	if !ok {
		fmt.Println("Unknown command")
		return
	}

	err := reg.callback()
	if err != nil {
		fmt.Printf("%w", err.Error())
	}
}
