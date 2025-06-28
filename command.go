package main

import (
	"fmt"
	"math/rand"
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
		"catch": {
			name:         "Catch",
			description:  "Tries to catch a pokemon. The chance of cathing it is based on a pokemon's base experience. The name of the pokemon  should be written after the catch command as an argument",
			numArguments: 1,
			callback:     commandCatch,
		},
		"inspect": {
			name:         "Inspect",
			description:  "Displays the stats of a pokemon that's already been caught. The name of the pokemon  should be written after the catch command as an argument",
			numArguments: 1,
			callback:     commandInspect,
		},
		"pokedex": {
			name:         "Pokedex",
			description:  "Displays all pokemon that have already been caught.",
			numArguments: 0,
			callback:     commandPokedex,
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

func commandCatch(configArg *pokeapi.Config, arg string) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	fmt.Printf("Throwing a Pokeball at %s...\n", arg)
	pokemon, err := pokeapi.GetPokemonStats(baseUrl+arg, configArg)
	if err != nil {
		return err
	}
	maxTreshold := 200
	catchTreshold := maxTreshold - pokemon.BaseExperience
	if rand.Intn(maxTreshold) < catchTreshold {
		fmt.Printf("%s was caught!\n", arg)
		fmt.Println("You may now inspect it with the inspect command.")
		configArg.Pokedex[arg] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", arg)
	}
	return nil
}

func commandInspect(configArg *pokeapi.Config, arg string) error {
	pokemon, ok := configArg.Pokedex[arg]
	if ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("	-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("	- %s\n", pokemonType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(configArg *pokeapi.Config, arg string) error {
	if len(configArg.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range configArg.Pokedex {
			fmt.Printf("	- %s\n", pokemon.Name)
		}
	}
	return nil
}
