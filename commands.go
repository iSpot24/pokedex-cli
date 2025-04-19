package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location given")
	}

	fmt.Println("Exploring pastoria-city-area...")
	result, err := cfg.apiClient.ListLocation(args[0])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, loc := range result.PokemonEncounters {
		fmt.Println(" - " + loc.Pokemon.Name)
	}

	return nil
}

func commandMap(cfg *config, args ...string) error {
	if cfg.nextPageURL == nil && cfg.previousPageURL != nil {
		return errors.New("you're on the last page")
	}

	result, err := cfg.apiClient.ListLocations(cfg.nextPageURL)

	if err != nil {
		return err
	}

	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextPageURL = result.Next
	cfg.previousPageURL = result.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousPageURL == nil {
		return errors.New("you're on the first page")
	}

	result, err := cfg.apiClient.ListLocations(cfg.previousPageURL)

	if err != nil {
		return err
	}

	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextPageURL = result.Next
	cfg.previousPageURL = result.Previous

	return nil
}

func commandHelp(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println()
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
