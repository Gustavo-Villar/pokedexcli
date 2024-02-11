// Package pokeapi defines structures to interact with the PokeAPI, focusing on location areas and their associated data.
package pokeapi

type LocationAreasResp struct {
	// Count indicates the total number of location areas available.
	Count int `json:"count"`
	// Next is a pointer to the URL for the next page of location areas, if available.
	Next *string `json:"next"`
	// Previous is a pointer to the URL for the previous page of location areas, if available.
	Previous *string `json:"previous"`
	// Results is a slice of location area summaries, including their names and URLs.
	Results []struct {
		Name string `json:"name"` // The name of the location area.
		URL  string `json:"url"`  // The URL to fetch more data about the location area.
	} `json:"results"`
}

type LocationArea struct {
	// EncounterMethodRates includes the rates at which Pokemon can be encountered in this area, per method and game version.
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"` // The name of the encounter method.
			URL  string `json:"url"`  // The URL to fetch more data about the encounter method.
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"` // The rate of encountering Pokemon with this method in the specified version.
			Version struct {
				Name string `json:"name"` // The game version name.
				URL  string `json:"url"`  // The URL to fetch more data about the game version.
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`

	// GameIndex is an index for identifying this location across different game versions.
	GameIndex int `json:"game_index"`
	// ID is the identifier for this location area.
	ID int `json:"id"`
	// Location provides the higher-level location to which this area belongs.
	Location struct {
		Name string `json:"name"` // The name of the higher-level location.
		URL  string `json:"url"`  // The URL to fetch more data about the higher-level location.
	} `json:"location"`
	// Name is the name of the location area.
	Name string `json:"name"`
	// Names includes the names of this location area in various languages.
	Names []struct {
		Language struct {
			Name string `json:"name"` // The language name.
			URL  string `json:"url"`  // The URL to fetch more data about the language.
		} `json:"language"`
		Name string `json:"name"` // The name of the location area in the specified language.
	} `json:"names"`
	// PokemonEncounters lists the Pokemon that can be encountered in this area, including encounter details and versions.
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"` // The name of the Pokemon.
			URL  string `json:"url"`  // The URL to fetch more data about the Pokemon.
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`           // The chance of encountering the Pokemon.
				ConditionValues []interface{} `json:"condition_values"` // Conditions under which the Pokemon can be encountered.
				MaxLevel        int           `json:"max_level"`        // The maximum level at which the Pokemon can be encountered.
				Method          struct {
					Name string `json:"name"` // The encounter method name.
					URL  string `json:"url"`  // The URL to fetch more data about the encounter method.
				} `json:"method"`
				MinLevel int `json:"min_level"` // The minimum level at which the Pokemon can be encountered.
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"` // The maximum chance of encountering any Pokemon in this area.
			Version   struct {
				Name string `json:"name"` // The game version.
				URL  string `json:"url"`  // The URL to fetch more data about the game version.
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
