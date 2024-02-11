// Package main includes the implementation for the "catch" command of the Pokedex CLI application.
// This file defines the functionality for attempting to catch a Pokemon by name.
package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// callbackCatch handles the "catch" command, which attempts to catch a specified Pokemon by name.
// It requires exactly one argument: the name of the Pokemon to catch.
func callbackCatch(cfg *config, args ...string) error {
	// Validates that exactly one argument, the Pokemon name, is provided.
	if len(args) != 1 {
		return errors.New("no pokemon name provided") // Returns an error if no Pokemon name is provided.
	}
	pokemonName := args[0] // The provided Pokemon name to catch.

	// Attempts to retrieve the specified Pokemon from the PokeAPI.
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err // Returns an error if the API call fails.
	}

	const threshold = 50                         // A constant threshold value for the catch success check.
	randNum := rand.Intn(pokemon.BaseExperience) // Generates a random number based on the Pokemon's base experience.

	if randNum > threshold {
		// If the random number exceeds the threshold, the attempt to catch the Pokemon fails.
		return fmt.Errorf("failed to catch %s!", pokemon.Name)
	}

	// If the attempt is successful, the Pokemon is added to the caughtPokemon map.
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("pokemon %s was caught!\n", pokemonName) // Prints a success message.

	return nil // Returns nil to indicate successful execution of the command.
}
