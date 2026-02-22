
package main

import (
	"fmt"
	"os"
)
func commandExit(cfg *Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
