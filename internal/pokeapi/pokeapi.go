package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	loc := LocationAreas{}
	if err := json.Unmarshal(body, &loc); err != nil {
		return err
	}
	for _, result := range loc.Results {
		fmt.Println(result.Name)
	}
	return nil
}
