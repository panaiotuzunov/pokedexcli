package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/panaiotuzunov/pokedexcli/internal/pokecache"
)

type Config struct {
	Previous *string
	Next     *string
	Cache    *pokecache.Cache
}

type LocationAreas struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func GetLocationAreas(url string, configArg *Config) error {
	var body []byte
	if cachedData, found := configArg.Cache.Get(url); found {
		body = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		configArg.Cache.Add(url, body)
	}
	loc := LocationAreas{}
	if err := json.Unmarshal(body, &loc); err != nil {
		return err
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	configArg.Next = loc.Next
	configArg.Previous = loc.Previous
	return nil
}

func GetPokemonEncounters(url string, configArg *Config) error {
	var body []byte
	if cachedData, found := configArg.Cache.Get(url); found {
		body = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		configArg.Cache.Add(url, body)
	}
	encounters := Pokemon{}
	if err := json.Unmarshal(body, &encounters); err != nil {
		return err
	}
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
