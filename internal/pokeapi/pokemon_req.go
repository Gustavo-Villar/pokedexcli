// Package pokeapi provides the functionality to interact with the PokeAPI, including fetching Pokemon data.
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon fetches the details of a Pokemon by its name using the PokeAPI.
// It attempts to retrieve the data from a cache before making an HTTP request to the PokeAPI.
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName // Defines the specific API endpoint for the requested Pokemon.
	fullUrl := baseURL + endpoint         // Constructs the full URL for the API request.

	// Attempt to retrieve the Pokemon data from the cache.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// If the data is found in the cache, deserialize it into a Pokemon struct.
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			// Return an error if there's an issue with deserialization.
			return Pokemon{}, err
		}
		// Return the cached Pokemon data.
		return pokemon, nil
	}

	// If the data is not found in the cache, make an HTTP GET request to the PokeAPI.
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		// Return an error if there's an issue creating the request.
		return Pokemon{}, err
	}

	// Perform the HTTP request.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Return an error if the HTTP request fails.
		return Pokemon{}, err
	}
	defer resp.Body.Close() // Ensure the response body is closed to prevent resource leaks.

	// Check the response status code for errors.
	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// Read the response body.
	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		// Return an error if reading the response body fails.
		return Pokemon{}, err
	}

	// Deserialize the response data into a Pokemon struct.
	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		// Return an error if there's an issue with deserialization.
		return Pokemon{}, err
	}

	// Add the retrieved data to the cache for future requests.
	c.cache.Add(fullUrl, dat)

	// Return the fetched Pokemon data.
	return pokemon, nil
}
