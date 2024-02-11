// Package main includes the implementation for the "explore" command of the Pokedex CLI application.
// This file defines the functionality to list Pokemon in a specified location area.
package main

import (
	"errors"
	"fmt"
)

// callbackExplore handles the "explore" command, allowing users to list Pokemon in a given location area.
// It requires exactly one argument: the name of the location area to explore.
func callbackExplore(cfg *config, args ...string) error {
	// Checks if the correct number of arguments were provided.
	if len(args) != 1 {
		return errors.New("no location area provided") // Returns an error if no location area name is provided.
	}
	locationAreaName := args[0] // Assigns the provided argument as the location area name.

	// Attempts to retrieve the specified location area from the PokeAPI.
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err // Returns an error if the API call fails.
	}

	// Prints the name of the location area and lists all Pokemon encountered there.
	fmt.Printf("+ Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("\t* - %s\n", pokemon.Pokemon.Name) // Prints each Pokemon's name found in the location area.
	}

	return nil // Returns nil to indicate successful execution of the command.
}
