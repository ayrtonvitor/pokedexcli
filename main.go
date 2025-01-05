package main

import (
	"log"

	"github.com/ayrtonvitor/pokedexcli/internal/pokeapi"
)

func main() {
	conf, err := loadConfig()
	if err != nil {
		log.Fatalf("Startup error: %v\n", err.Error())
	}
	apiConf, ok := conf["api"]
	if !ok {
		log.Fatalf("Could not get the api configs")
	}
	pokeapi.Setup(apiConf)

	setupCommands()
	run()
}
