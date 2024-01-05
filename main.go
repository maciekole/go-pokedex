package main

import (
	"bufio"
	"fmt"
	"github.com/maciekole/pokedex/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("Help command invoked")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func commandMap() error {
	fmt.Println("Map command invoked")
	locationName, err := pokeapi.GetLocation(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(locationName)
	return nil
}

func commandMapB() error {
	err := pokeapi.Xd()
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits program.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "The map command displays the names of 20 location areas in the Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "The map command displays the names of previous 20 location areas in the Pokemon world.",
			callback:    commandMapB,
		},
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		scanner.Scan()
		words := scanner.Text()

		if len(words) == 0 {
			continue
		}

		commandName := words

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
