package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	if pageURL != nil {
		fullUrl = *pageURL
	}

	// check the cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hit!")
		// cache hit
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}

	// fmt.Println("cache miss!")

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
		return LocationAreasResp{}, fmt.Errorf("bad status code: %v\n", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	// Save data to cache if needed
	c.cache.Add(fullUrl, dat)

	return locationAreasResp, nil

}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullUrl := baseURL + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullUrl)
	if ok {
		// fmt.Println("cache hit!")
		// cache hit
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}

	// fmt.Println("cache miss!")

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
		return LocationArea{}, fmt.Errorf("bad status code: %v\n", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	// Save data to cache if needed
	c.cache.Add(fullUrl, dat)

	return locationArea, nil

}
