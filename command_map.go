package main

import (
	"errors"
	"fmt"
	"strings"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		return err
	}

	fmt.Println("+ Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t* [%s] - %s\n", getAreaId(area.URL), area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("You're on the first page already")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)

	if err != nil {
		return err
	}

	fmt.Println("+ Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t* [%s] - %s\n", getAreaId(area.URL), area.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}

func getAreaId(areaURL string) string {

	parts := strings.Split(areaURL, "/")

	// Initialize an empty string for the ID
	var idStr string
	// Check if the last part is empty (due to trailing slash)
	if parts[len(parts)-1] == "" {
		// Use the second last part if the last is empty
		idStr = parts[len(parts)-2]
	} else {
		// Otherwise, use the last part
		idStr = parts[len(parts)-1]
	}

	return string(idStr)

}
