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
		

		g, err := game.NewWarGame()
		if err != nil {
			return fmt.Errorf("could not start War game: %v", err)
		}
		currentGame = g

		fmt.Println("War: Use 'playturn' to play a round")
		
		return nil
	case "blackjack":
		fmt.Println("Starting a game of Blackjack!")

		g, err := game.NewBlackjackGame()
		if err != nil {
			return fmt.Errorf("could not start Blackjack game: %v", err)
		}
		currentGame = g

		fmt.Println("Blackjack: Use 'hit' or 'stand' to play! Use 'deal' to deal a new game from the same deck")
		return nil
	default:
		fmt.Println("Sorry that game is not available\n\nCurrent games available to play:\n - war\n - blackjack")
		return nil
	}
}

