package main

import (
	"fmt"
)

func commandQuit (args ...string) error {
	if currentGame == nil {
		fmt.Println("No game is currently being played.")
		return nil
	}

	currentGame = nil
	fmt.Println("You have quit the current game. Feel free to 'play' a new one!")
	return nil
}