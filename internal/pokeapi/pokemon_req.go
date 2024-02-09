package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullUrl := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hit!")
		// cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)

		if err != nil {
			return Pokemon{}, err
		}

		return pokemon, nil
	}

	// fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v\n", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	// Save data to cache if needed
	c.cache.Add(fullUrl, dat)

	return pokemon, nil

}
