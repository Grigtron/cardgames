package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists:= getCommands()[commandName]
		if exists {
			err := command.callback(args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} 
		if currentGame != nil {
			err := currentGame.HandleCommand(commandName, args...)
			if err != nil {
				fmt.Println("Game error:", err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)

	return words
}

type cliCommand struct {
	name string
	description string
	callback func(...string) error
}


func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays the help menu",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exits the program",
			callback: commandExit,
		},
		"play": {
			name: "play <game_name>",
			description: "Begin a new specified game",
			callback: commandPlay,
		},
	}
}
