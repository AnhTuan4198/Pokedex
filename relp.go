package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AnhTuan4198/Pokedex/pokeapi"
	"github.com/AnhTuan4198/Pokedex/pokecache"
)

type config struct {
	pokeApiClient   pokeapi.Client
	cache           *pokecache.Cache
	nextLocationUrl *string
	prevLocationUrl *string
	exploreLocation *string
}

type CLICommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		terminalInput := scanner.Text()
		if terminalInput == "" {
			continue
		}
		commandSlices := strings.Fields(terminalInput)
		if commandSlices[0] == "" {
			continue
		}
		arg := ""
		if len(commandSlices) > 1 {
			arg = commandSlices[1]
		}
		cmds := getCommands()
		cmd, ok := cmds[commandSlices[0]]
		if !ok {
			fmt.Printf("Command not found!\n")
			continue
		} else {
			shouldExit := cmd.callback(cfg, arg)
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

	cmds["explore"] = CLICommand{
		name:        "explore",
		description: "Explore all pokemon in given area",
		callback:    commandExplore,
	}

	return cmds
}
