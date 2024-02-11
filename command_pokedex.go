package main

import (
	"fmt"
	"sort"

	"github.com/gustavo-villar/pokedexcli/internal/pokeapi"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("***** POKEDEX *****")
	pokemons := orderById(cfg.caughtPokemon)
	fmt.Println()

	fmt.Println("--- List of caught pokemon: ")
	for _, pokemon := range pokemons {
		fmt.Printf("> Pokemon Name: %s\tID: %v\n", pokemon.Name, pokemon.ID)
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("--- Pokemon Complete Info: ")
	for _, pokemon := range pokemons {
		fmt.Printf("> Pokemon Name: %s\tID: %v\n", pokemon.Name, pokemon.ID)
		fmt.Printf("\t- Height: %v\n", pokemon.Height)
		fmt.Printf("\t- Weight: %v\n", pokemon.Weight)
		fmt.Print("\t- Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t\t+ %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Print("\t- Types:\n")
		for _, typ := range pokemon.Types {
			fmt.Printf("\t\t+ %s\n", typ.Type.Name)
		}
	}

	return nil
}

func orderById(caughtPokemon map[string]pokeapi.Pokemon) []pokeapi.Pokemon {
	// Convert map to slice of pokemons
	pokemons := make([]pokeapi.Pokemon, 0, len(caughtPokemon))
	for _, pokemon := range caughtPokemon {
		pokemons = append(pokemons, pokemon)
	}

	// Sort pokemons slice by ID
	sort.Slice(pokemons, func(i, j int) bool {
		return pokemons[i].ID < pokemons[j].ID
	})

	return pokemons
}
