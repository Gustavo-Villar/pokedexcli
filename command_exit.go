// Package main includes the implementation for the exit command of the Pokedex CLI application.
// It defines a function that safely terminates the application when invoked.
package main

import (
	"fmt"
	"os"
)

// callbackExit handles the "exit" command within the Pokedex CLI application.
// It prints a message indicating the application is turning off and then terminates the process.
func callbackExit(cfg *config, args ...string) error {
	fmt.Println("turning off...") // Prints a farewell message to indicate the application is about to exit.
	os.Exit(0)                    // Terminates the application with a status code of 0, indicating a normal exit.
	return nil                    // This return statement is technically unnecessary due to os.Exit() above, but it's kept for consistency with the callback function signature.
}
