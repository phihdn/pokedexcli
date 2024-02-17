package main

import (
	"time"

	"github.com/phihdn/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(5 * time.Minute),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
