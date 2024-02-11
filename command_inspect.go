// Package main includes the implementation for the "inspect" command of the Pokedex CLI application.
// This file defines the functionality to display details of a caught Pokemon by name.
package main

import (
	"errors"
	"fmt"
)

// callbackInspect handles the "inspect" command, allowing users to view details of a caught Pokemon.
// It requires exactly one argument: the name of the Pokemon to inspect.
func callbackInspect(cfg *config, args ...string) error {
	// Validates that exactly one argument, the Pokemon name, is provided.
	if len(args) != 1 {
		return errors.New("no pokemon name provided") // Returns an error if no Pokemon name is provided.
	}
	pokemonName := args[0] // The provided Pokemon name to inspect.

	// Attempts to retrieve the specified Pokemon from the caughtPokemon map.
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("this pokemon was not yet caught") // Returns an error if the Pokemon hasn't been caught yet.
	}

	// Prints detailed information about the caught Pokemon.
	fmt.Printf("> Pokemon Name: %s\tID: %v\n", pokemon.Name, pokemon.ID)
	fmt.Printf("\t- Height: %v\n", pokemon.Height)
	fmt.Printf("\t- Weight: %v\n", pokemon.Weight)
	fmt.Print("\t- Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t\t+ %s: %v\n", stat.Stat.Name, stat.BaseStat) // Prints each stat of the Pokemon.
	}
	fmt.Print("\t- Types:\n")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t\t+ %s\n", typ.Type.Name) // Prints each type of the Pokemon.
	}

	return nil // Returns nil to indicate successful execution of the command.
}
