package pokeapi

import (
	"net/http"
	"time"

	"github.com/gustavo-villar/pokedexcli/internal/pokecache"
)

const baseURL string = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
