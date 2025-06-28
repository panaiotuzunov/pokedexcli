package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/panaiotuzunov/pokedexcli/internal/pokeapi"
	"github.com/panaiotuzunov/pokedexcli/internal/pokecache"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	ConfigArg := pokeapi.Config{}
	startUrl := "https://pokeapi.co/api/v2/location-area"
	ConfigArg.Next = &startUrl
	ConfigArg.Previous = nil
	ConfigArg.Cache = pokecache.NewCache(time.Minute * 10)
	ConfigArg.Pokedex = map[string]pokeapi.Pokemon{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		command, ok := getCommands()[commandName]
		if ok {
			if len(words)-1 != command.numArguments {
				fmt.Printf("Command %s takes exactly %d arguments", command.name, command.numArguments)
				fmt.Println()
				continue
			}
			arg := ""
			if len(words) > 1 {
				arg = words[1]
			}
			err := command.callback(&ConfigArg, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Println("Unknown command")
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
