package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/panaiotuzunov/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	ConfigArg := pokeapi.Config{}
	startUrl := "https://pokeapi.co/api/v2/location-area"
	ConfigArg.Next = &startUrl
	ConfigArg.Previous = nil

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
			err := command.callback(&ConfigArg)
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
