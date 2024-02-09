package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

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
	}
}

func startRelp(cfg *config) string {
	// Creates the input variable
	var input string = ""

	// Creates a new scanner
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Gets the input from the user
		fmt.Print(">_: ")      // For Aesthetics only
		scanner.Scan()         // Scans the input
		input = scanner.Text() // Gets the text fom the input

		cleanedInput := cleanInput(input)

		// Verifies if the user input is empty
		if len(cleanedInput) == 0 {
			continue
		}
		commandName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 1 {
			args = cleanedInput[1:]
		}

		avaiableCommands := getCommands()

		command, ok := avaiableCommands[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue

		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func cleanInput(str string) []string {
	loweredString := strings.ToLower(str)
	words := strings.Fields(loweredString)

	return words
}
