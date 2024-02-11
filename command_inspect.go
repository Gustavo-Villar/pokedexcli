package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return errors.New("this pokemon was not yet caught")
	}

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

	return nil
}
