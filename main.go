package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		wordList := cleanInput(input)
		if len(wordList) == 0 {
			continue
		}
		commandName := wordList[0]
		command, exists := cliCommandMap[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
