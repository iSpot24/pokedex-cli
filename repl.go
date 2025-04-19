package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iSpot24/pokedex-cli/internal/pokeapi"
)

func openPokedex(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("No command given")
			continue
		}
		commandName := words[0]
		args := words[1:]

		if command, ok := getCommands()[commandName]; ok {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")

	}
}

type config struct {
	apiClient       pokeapi.Client
	nextPageURL     *string
	previousPageURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "List available locations or navigate to next page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Navigate to previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"help": {
			name:        "help",
			description: "How to use the pokedex",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
