// Package main includes the implementation for the "pokedex" command of the Pokedex CLI application.
// This file defines functionality to display a list and complete details of caught Pokemon.
package main

import (
	"fmt"
	"sort"

	"github.com/gustavo-villar/pokedexcli/internal/pokeapi"
)

// callbackPokedex handles the "pokedex" command, displaying a summary and detailed information of all caught Pokemon.
// It does not require any arguments and will list Pokemon sorted by their ID.
func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("***** POKEDEX *****")       // Prints a header for the Pokedex output.
	pokemons := orderById(cfg.caughtPokemon) // Orders caught Pokemon by their ID for display.
	fmt.Println()

	fmt.Println("--- List of caught pokemon: ")
	for _, pokemon := range pokemons {
		// Prints a brief summary for each caught Pokemon, including its name and ID.
		fmt.Printf("> Pokemon Name: %s\tID: %v\n", pokemon.Name, pokemon.ID)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("--- Pokemon Complete Info: ")
	for _, pokemon := range pokemons {
		// Prints detailed information for each caught Pokemon.
		fmt.Printf("> Pokemon Name: %s\tID: %v\n", pokemon.Name, pokemon.ID)
		fmt.Printf("\t- Height: %v\n", pokemon.Height)
		fmt.Printf("\t- Weight: %v\n", pokemon.Weight)
		fmt.Print("\t- Stats:\n")
		for _, stat := range pokemon.Stats {
			// Lists each stat of the Pokemon.
			fmt.Printf("\t\t+ %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Print("\t- Types:\n")
		for _, typ := range pokemon.Types {
			// Lists each type of the Pokemon.
			fmt.Printf("\t\t+ %s\n", typ.Type.Name)
		}
	}

	return nil // Returns nil to indicate successful execution of the command.
}

// orderById takes a map of caught Pokemon and returns a slice of Pokemon sorted by their ID.
// It facilitates sorting and displaying Pokemon in the callbackPokedex function.
func orderById(caughtPokemon map[string]pokeapi.Pokemon) []pokeapi.Pokemon {
	// Converts the map of caught Pokemon into a slice for sorting.
	pokemons := make([]pokeapi.Pokemon, 0, len(caughtPokemon))
	for _, pokemon := range caughtPokemon {
		pokemons = append(pokemons, pokemon)
	}

	// Sorts the slice of Pokemon by ID in ascending order.
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	return pokemons // Returns the sorted slice of Pokemon.
}
