package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type CLICommand struct {
	name        string
	description string
	callback    func() error
}

func (cmd CLICommand) getInformation() {
	fmt.Printf("%v: %v \n", cmd.name, cmd.description)
	fmt.Println("")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var cmds = make(map[string]CLICommand)

	cmds["exit"] = CLICommand{
		name:        "exit",
		description: "Using to stop application",
		callback: func() error {
			return errors.New("error");
		},
	}

	cmds["help"] = CLICommand{
		name:        "help",
		description: "Show all available commands",
		callback: func() error {
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage: ")
			fmt.Println("")
			for _, val := range cmds {
				val.getInformation()
			}
			return nil
		},
	}
	fmt.Println("Pokedex > ")
	for scanner.Scan() {
		cmdInput := scanner.Text()
		cmd, ok := cmds[cmdInput]
		if !ok {
			fmt.Printf("Command not found!")
		} else {
			shouldExit := cmd.callback();
			if shouldExit != nil{
				return
			}
		}
	}
}
