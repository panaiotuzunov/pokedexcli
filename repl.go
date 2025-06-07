package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
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
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		command := words[0]
		_, ok := supportedCommands[command]
		if ok {
			supportedCommands[command].callback(supportedCommands)
			continue
		}
		fmt.Println("Unknown command")
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(_ map[string]cliCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for name, command := range commands {
		fmt.Printf("%s: %s\n", name, command.description)
	}
	return nil
}
