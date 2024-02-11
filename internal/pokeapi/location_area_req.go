// Package pokeapi includes functionality for fetching data about location areas from the PokeAPI.
package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocationAreas fetches a list of location areas from the PokeAPI or a specific page of location areas if a pageURL is provided.
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint // Constructs the URL for the location areas endpoint.

	if pageURL != nil {
		fullUrl = *pageURL // Uses the provided pageURL if available.
	}

	// Attempts to retrieve the location areas data from the cache.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// If data is found in the cache, unmarshal it into the LocationAreasResp struct.
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err // Returns error if unmarshaling fails.
		}
		return locationAreasResp, nil // Returns the cached response.
	}

	// If cache miss, make an HTTP GET request to the PokeAPI.
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Unmarshal the response body into the LocationAreasResp struct.
	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	// Add the response data to the cache.
	c.cache.Add(fullUrl, dat)

	return locationAreasResp, nil
}

// GetLocationArea fetches details of a specific location area by name from the PokeAPI.
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURL + endpoint // Constructs the URL for the specific location area endpoint.

	// Attempts to retrieve the location area data from the cache.
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// If data is found in the cache, unmarshal it into the LocationArea struct.
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err // Returns error if unmarshaling fails.
		}
		return locationArea, nil // Returns the cached response.
	}

	// If cache miss, make an HTTP GET request to the PokeAPI.
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Unmarshal the response body into the LocationArea struct.
	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	// Add the response data to the cache.
	c.cache.Add(fullUrl, dat)

	return locationArea, nil
}
