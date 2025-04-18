package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func openPokedex() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			fmt.Println("No command given")
			continue
		}
		arg := words[0]

		if command, ok := getCommands()[arg]; ok {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
