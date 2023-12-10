package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/AnhTuan4198/Pokedex/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}
type CLICommand struct {
	name        string
	description string
	callback    func(*config) error
}

func (cmd CLICommand) getInformation() {
	fmt.Printf("%v: %v \n", cmd.name, cmd.description)
	fmt.Println("")
}

func startRELP(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cmdInput := scanner.Text()
		if cmdInput == "" {
			continue
		}
		cmds := getCommands()
		cmd, ok := cmds[cmdInput]
		if !ok {
			fmt.Printf("Command not found!\n")
			continue
		} else {
			shouldExit := cmd.callback(cfg)
			if shouldExit != nil {
				return
			}
			continue
		}
	}
}

func getCommands() map[string]CLICommand {
	var cmds = make(map[string]CLICommand)

	cmds["exit"] = CLICommand{
		name:        "exit",
		description: "Using to stop application",
		callback:    commandExit,
	}

	cmds["help"] = CLICommand{
		name:        "help",
		description: "Show all available commands",
		callback:    commandHelp,
	}

	cmds["map"] = CLICommand{
		name:        "map",
		description: "Show location areas",
		callback:    commandMap,
	}

	cmds["mapb"] = CLICommand{
		name:        "mapb",
		description: "Show previous location areas",
		callback:    commandMapb,
	}

	return cmds
}
