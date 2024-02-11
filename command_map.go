// Package main includes functionalities for navigating and displaying location areas within the Pokedex CLI application.
// It defines callback functions for map navigation commands.
package main

import (
	"errors"
	"fmt"
	"strings"
)

// callbackMap fetches and displays the next page of location areas from the PokeAPI.
// It updates the application configuration with the URLs for the next and previous pages of location areas.
func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL) // Attempts to fetch the next page of location areas.

	if err != nil {
		return err // Returns an error if the API request fails.
	}

	fmt.Println("+ Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t* [%s] - %s\n", getAreaId(area.URL), area.Name) // Prints each location area with its ID and name.
	}
	cfg.nextLocationAreaURL = resp.Next         // Updates the URL for the next page of locations.
	cfg.previousLocationAreaURL = resp.Previous // Updates the URL for the previous page of locations.

	return nil
}

// callbackMapb fetches and displays the previous page of location areas from the PokeAPI.
// It performs a similar function to callbackMap but for navigating backwards.
func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("You're on the first page already") // Returns an error if there is no previous page.
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL) // Attempts to fetch the previous page of location areas.

	if err != nil {
		return err // Returns an error if the API request fails.
	}

	fmt.Println("+ Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t* [%s] - %s\n", getAreaId(area.URL), area.Name) // Prints each location area with its ID and name.
	}
	cfg.nextLocationAreaURL = resp.Next         // Updates the URL for the next page of locations.
	cfg.previousLocationAreaURL = resp.Previous // Updates the URL for the previous page of locations.

	return nil
}

// getAreaId extracts and returns the ID of a location area from its URL.
// It parses the URL string to find the area ID, accounting for potential trailing slashes.
func getAreaId(areaURL string) string {
	parts := strings.Split(areaURL, "/") // Splits the URL into parts.

	var idStr string // Initializes an empty string for the ID.
	if parts[len(parts)-1] == "" {
		idStr = parts[len(parts)-2] // Uses the second last part if the last is empty.
	} else {
		idStr = parts[len(parts)-1] // Otherwise, uses the last part.
	}

	return idStr
}
