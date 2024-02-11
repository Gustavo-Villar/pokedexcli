// Package pokeapi provides a client for interacting with the PokeAPI, including methods for fetching Pokemon data.
// It utilizes a caching mechanism to reduce the number of requests made to the PokeAPI.
package pokeapi

import (
	"net/http"
	"time"

	"github.com/gustavo-villar/pokedexcli/internal/pokecache"
)

// baseURL is the base URL for the PokeAPI, used to construct endpoint URLs for API requests.
const baseURL string = "https://pokeapi.co/api/v2"

// Client struct defines the structure of the PokeAPI client, including a cache for responses and an HTTP client for making requests.
type Client struct {
	cache      pokecache.Cache // cache is an instance of a caching mechanism to store and retrieve API responses.
	httpClient http.Client     // httpClient is used to make HTTP requests to the PokeAPI, with a configured timeout.
}

// NewClient initializes and returns a new instance of a PokeAPI client.
// It takes a cacheInterval argument, which specifies the duration for which cached responses should be considered valid.
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval), // Initializes the cache with the specified cache interval.
		httpClient: http.Client{
			Timeout: time.Minute, // Configures the HTTP client with a timeout of 1 minute to avoid hanging requests.
		},
	}
}
