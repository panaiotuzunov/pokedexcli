package main

import (
	"fmt"
	"os"

	"github.com/panaiotuzunov/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

func getCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "Help",
			description: "Returns information about app usage",
			callback:    commandHelp,
		},
		"map": {
			name:        "Map",
			description: "Returns the names of the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "Map Back",
			description: "Returns the names of the previous 20 location areas",
			callback:    commandMapb,
		},
	}
	return supportedCommands
}

func commandExit(configArg *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(configArg *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap(configArg *pokeapi.Config) error {
	err := pokeapi.GetLocationAreas(*configArg.Next, configArg)
	if err != nil {
		return err
	}
	return nil
}

func commandMapb(configArg *pokeapi.Config) error {
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
