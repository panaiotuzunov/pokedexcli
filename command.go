package main

import (
	"fmt"
	"os"

	"github.com/panaiotuzunov/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name         string
	description  string
	numArguments int
	callback     func(*pokeapi.Config, string) error
}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:         "exit",
			description:  "Exit the Pokedex",
			numArguments: 0,
			callback:     commandExit,
		},
		"help": {
			name:         "Help",
			description:  "Returns information about app usage",
			numArguments: 0,
			callback:     commandHelp,
		},
		"map": {
			name:         "Map",
			description:  "Returns the names of the next 20 location areas",
			numArguments: 0,
			callback:     commandMap,
		},
		"mapb": {
			name:         "Map Back",
			description:  "Returns the names of the previous 20 location areas",
			numArguments: 0,
			callback:     commandMapb,
		},
		"explore": {
			name:         "Explore",
			description:  "Shows all pokemon in the selected location area. The name of the location area should be written after the explore command as an argument",
			numArguments: 1,
			callback:     commandExplore,
		},
	}
	return supportedCommands
}

func commandExit(configArg *pokeapi.Config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configArg *pokeapi.Config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap(configArg *pokeapi.Config, arg string) error {
	err := pokeapi.GetLocationAreas(*configArg.Next, configArg)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(configArg *pokeapi.Config, arg string) error {
	if configArg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	err := pokeapi.GetLocationAreas(*configArg.Previous, configArg)
	if err != nil {
		return err
	}
	return nil
}

func commandExplore(configArg *pokeapi.Config, arg string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	fmt.Printf("Exploring %s...", arg)
	fmt.Println()
	fmt.Println("Found Pokemon:")
	err := pokeapi.GetPokemonEncounters(baseUrl+arg, configArg)
	if err != nil {
		return err
	}
	return nil
}
