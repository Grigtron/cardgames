package main

import (
	"fmt"

	"github.com/Grigtron/cardgames/game"
)


func commandPlay(args ...string) error {
	if len(args) != 1 {
		fmt.Println("Please provide a game you wish to play\n\nCurrent games available to play:\n - war")
		return nil
	}
	switch args[0] {
	case "war":
		fmt.Println("Starting a game of War!")

		var err error
		currentGame, err = game.NewWarGame()
		if err != nil {
			return fmt.Errorf("could not start War game: %v", err)
		}

		err = currentGame.PlayTurn()
		if err != nil {
			fmt.Println("error during the game", err)
		}
		
		return nil
	default:
		fmt.Println("Sorry that game is not available\n\nCurrent games available to play:\n - war")
		return nil
	}
}

