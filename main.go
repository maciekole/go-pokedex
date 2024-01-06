package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/maciekole/pokedex/pokeapi"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func commandHelp(cfg *config) error {
	fmt.Println("Help command invoked")
	return nil
}

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	fmt.Println("Map command invoked")
	locations, nextStartingLocation, nextStartingLocationBack, err := pokeapi.GetLocationsForward(cfg.Next)
	if err != nil {
		fmt.Println(err)
	}

	cfg.Next = nextStartingLocation
	cfg.Previous = nextStartingLocationBack

	for _, location := range locations {
		fmt.Println(location)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.Previous == nil {
		return errors.New("you are on first page")
	}
	fmt.Println("Map command invoked")
	locations, nextStartingLocation, nextStartingLocationBack, err := pokeapi.GetLocationsBackward(cfg.Previous)
	if err != nil {
		fmt.Println(err)
	}

	cfg.Next = nextStartingLocation
	cfg.Previous = nextStartingLocationBack

	for _, location := range locations {
		fmt.Println(location)
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
	cfg := &config{}
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
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(fmt.Sprintf("\nDEBUG: cfg %v", cfg))
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
