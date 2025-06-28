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
	Pokedex  map[string]Pokemon
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

type PokemonEncounter struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
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
	encounters := PokemonEncounter{}
	if err := json.Unmarshal(body, &encounters); err != nil {
		return err
	}
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

func GetPokemonStats(url string, configArg *Config) (Pokemon, error) {
	var body []byte
	if cachedData, found := configArg.Cache.Get(url); found {
		body = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		configArg.Cache.Add(url, body)
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
