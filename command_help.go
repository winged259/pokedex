package main

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range cliCommandMap {
		fmt.Printf("- %s: %s\n", command.name, command.description)
	}
	return nil
}