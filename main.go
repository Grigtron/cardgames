package main

import(
	"fmt"
	"github.com/Grigtron/cardgames/deck"
)

func main() {
	newDeck := deck.ShuffleDeck(deck.NewDeck())
	for _, card := range newDeck.Cards {
		fmt.Printf("%v\n", card)
	}
	fmt.Println(len(newDeck.Cards))

	fmt.Println("Welcome to Go Card Games!")
	fmt.Println("Use 'help' to get a list of current commands")
	startRepl()
}