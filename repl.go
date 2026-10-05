package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		fmt.Print("Pokedex > ")
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command, exists := getCommands()[words[0]]
		if exists {
			err := command.callback()
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
	callback		func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
