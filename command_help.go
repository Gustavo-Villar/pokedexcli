// Package main includes the implementation of the help command for the Pokedex CLI application.
// It defines a function to display a help menu with available commands and their descriptions.
package main

import "fmt"

// callbackHelp displays the help menu for the Pokedex CLI application.
// It iterates over the available commands, printing each command's name and description to the user.
func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("**Welcome to the Pokedex help Menu!**") // Prints a welcome message for the help menu.
	fmt.Println("-- Here are your available commands:")  // Introduces the list of available commands.
	fmt.Println("")

	availableCommands := getCommands() // Retrieves the map of available commands.

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s:\n\t %s \n", cmd.name, cmd.description) // Prints each command's name and description.
	}
	fmt.Println("")

	return nil // Returns nil to indicate successful execution of the help command.
}
