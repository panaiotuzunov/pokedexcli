package main

type cliCommand struct {
	name        string
	description string
	callback    func(map[string]cliCommand) error
}
