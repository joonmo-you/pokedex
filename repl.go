package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joonmo-you/pokedexcli/internal/pokeapi"
)

type config struct {
	commands		  	 map[string]cliCommand
	pokeapiClient 	 pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon		 map[string]pokeapi.Pokemon
}

func startRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, exists := config.commands[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}


type cliCommand struct {
	name				string
	description	string
	callback		func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:		 commandCatch,

		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "exlore <location_name>",
			description: "Explore a location",
			callback:    commandExlore,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback:		 commandInspect,
		},
		"map": {
			name: 			 "map",
			description: "Display the names of 20 location areas in the Pokemon world",
			callback: 	 commandMapf,
		},
		"mapb": {
			name: 			 "mapb",
			description: "Displays the previous 20 locations areas in the Pokemon world",
			callback: 	 commandMapb,
		},
	}
}
