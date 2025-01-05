package main

import "fmt"

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, v := range commands {
		fmt.Printf("%s: %s", v.name, v.description)
	}
	return nil
}
