package game

import (
	//"fmt"
	"github.com/Grigtron/cardgames/deck"
)

type WarGame struct {
	PlayerDeck deck.Deck
	ComputerDeck deck.Deck
}

func NewWarGame() *WarGame {
	d := deck.NewDeck()
	deck.ShuffleDeck(d)
	
	return nil

}