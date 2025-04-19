package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
)

func commandPokedex(cfg *config, args ...string) error {
	pokedexLen := len(cfg.pokedex)

	if pokedexLen == 0 {
		return errors.New("your pokedex is empty... go catch some pokemon!")
	}

	fmt.Println("")
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  -%s \n", pokemon.Name)
	}
	fmt.Println("")

	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location given")
	}

	result, err := cfg.apiClient.GetLocation(args[0])

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	for _, loc := range result.PokemonEncounters {
		fmt.Println(" - " + loc.Pokemon.Name)
	}
	fmt.Println("")

	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name given")
	}

	pokemonName := args[0]
	pokemon, err := cfg.apiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	userChance := rand.Float64()
	difficulty := 1.0 - 1.0/math.Log10(float64(pokemon.BaseExperience))

	if userChance >= difficulty {
		cfg.pokedex[pokemonName] = pokemon
		fmt.Printf("%v was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%v escaped!\n", pokemonName)
	}
	fmt.Println("")

	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name given")
	}

	pokemon, found := cfg.pokedex[args[0]]

	if !found {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("")
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, item := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", item.Stat.Name, item.BaseStat)
	}

	fmt.Println("Types:")
	for _, item := range pokemon.Types {
		fmt.Printf("  - %s\n", item.Type.Name)
	}
	fmt.Println("")

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

	fmt.Println("")
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println("")

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

	fmt.Println("")
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}
	fmt.Println("")

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
