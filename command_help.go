package main

import "fmt"

func commandHelp(args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to Go Card Games!")
	fmt.Println("Commands:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
	
}