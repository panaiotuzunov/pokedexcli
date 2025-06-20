package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range getCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}

func commandMap() error {

	return nil
}

func commandMapb() error {

	return nil
}
