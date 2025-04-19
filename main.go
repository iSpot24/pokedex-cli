package main

import (
	"time"

	"github.com/iSpot24/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		apiClient: pokeClient,
	}
	openPokedex(cfg)
}
