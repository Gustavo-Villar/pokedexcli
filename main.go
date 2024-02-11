// Package main defines the entry point for the Pokedex CLI application.
// It manages application configuration, including the PokeAPI client, and tracks the caught Pokemon.
package main

import (
	"time"

	"github.com/gustavo-villar/pokedexcli/internal/pokeapi"
)

// config struct holds the configuration for the Pokedex CLI application.
// It includes the PokeAPI client for fetching Pokemon data, URLs for navigating between location areas,
// and a map of caught Pokemon with their names as keys.
type config struct {
	pokeapiClient           pokeapi.Client             // PokeAPI client for making requests to the PokeAPI.
	nextLocationAreaURL     *string                    // URL to the next location area in the PokeAPI, used for navigation.
	previousLocationAreaURL *string                    // URL to the previous location area in the PokeAPI, used for navigation.
	caughtPokemon           map[string]pokeapi.Pokemon // Map of caught Pokemon, keyed by Pokemon name.
}

// main is the entry point of the Pokedex CLI application.
// It initializes the application configuration, including the PokeAPI client and the map of caught Pokemon,
// and starts the REPL (Read-Eval-Print Loop) for interacting with the application.
func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),     // Initializes the PokeAPI client with a cache duration of 1 hour.
		caughtPokemon: make(map[string]pokeapi.Pokemon), // Initializes the map of caught Pokemon.
	}

	startRelp(&cfg) // Starts the REPL with the initial application configuration.
}
