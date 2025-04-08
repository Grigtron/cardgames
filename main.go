package main

import (
	"fmt"

	"github.com/Grigtron/cardgames/game"
)
var currentGame *game.WarGame
func main() {
	fmt.Println("Welcome to Go Card Games!")
	fmt.Println("Use 'help' to get a list of current commands")
	startRepl()
}