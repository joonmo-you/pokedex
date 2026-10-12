package main

import (
	"time"

	"github.com/joonmo-you/pokedexcli/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	config := &config{
		commands: getCommands(),
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapiClient,
	}
	startRepl(config)
}
