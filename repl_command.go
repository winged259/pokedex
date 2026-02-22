package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

var cliCommandMap map[string]cliCommand

func init() {
	cliCommandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next location",
			callback:    commandGetMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous location",
			callback:    commandGetMapBack,
		},
	}
}
