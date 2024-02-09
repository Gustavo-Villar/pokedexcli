package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("**Welcome to the Pokedex help Menu!**")
	fmt.Println("-- Here are your avaiable commands:")
	fmt.Println("")

	avaiableCommands := getCommands()

	for _, cmd := range avaiableCommands {
		fmt.Printf(" - %s:\n\t %s \n", cmd.name, cmd.description)
	}
	fmt.Println("")

	return nil
}
