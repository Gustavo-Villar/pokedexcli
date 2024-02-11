// Package main includes the REPL (Read-Eval-Print Loop) functionality for the Pokedex CLI application.
// It provides an interactive command-line interface for users to explore and interact with the Pokemon universe.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// cliCommand struct defines a CLI command within the Pokedex application.
// It includes the command name, a description, and a callback function that executes the command.
type cliCommand struct {
	name        string                         // Name of the command.
	description string                         // Description of what the command does.
	callback    func(*config, ...string) error // Callback function to execute when the command is invoked.
}

// getCommands returns a map of available CLI commands by their name.
// Each command is associated with its functionality, including help, exit, map navigation, exploration, catching, and inspecting Pokemon.
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "map",
			description: "Lists the last page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch the pokemon and add it to your pokedex",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Prints Pokemon information if already caught",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints entire pokedex of caught pokemon",
			callback:    callbackPokedex,
		},
	}
}

// startRelp starts the Read-Eval-Print Loop (REPL) for the Pokedex CLI application.
// It continuously reads user input, executes commands, and prints outputs until the exit command is invoked.
func startRelp(cfg *config) {
	// Initializes the command input loop.
	var input string

	// Creates a new scanner to read input from standard input.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">_: ")      // Prints a prompt for user input.
		scanner.Scan()         // Reads the next line of input.
		input = scanner.Text() // Retrieves the text from the scanner.

		cleanedInput := cleanInput(input) // Cleans and parses the input into a slice of strings.

		// Skips processing if the input is empty.
		if len(cleanedInput) == 0 {
			continue
		}

		// Separates the command name and arguments from the cleaned input.
		commandName, args := cleanedInput[0], cleanedInput[1:]

		// Retrieves the map of available commands.
		availableCommands := getCommands()

		// Executes the command if it exists.
		if command, ok := availableCommands[commandName]; ok {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err) // Prints any errors returned by the command.
			}
		} else {
			fmt.Println("Invalid Command") // Notifies the user of invalid commands.
		}
	}
}

// cleanInput takes a string input and returns a slice of strings representing the cleaned and lowercased input split into words.
func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)  // Converts the input string to lowercase.
	words := strings.Fields(loweredString) // Splits the string into a slice of words.

	return words
}
