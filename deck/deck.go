package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	suits := []Suit{Clubs, Diamonds, Hearts, Spades}
	ranks := []Rank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
	cards := make([]Card, 0, len(suits) * len(ranks))
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, Card{suit, rank})
		}
	}
	return Deck{Cards: cards}
}

func ShuffleDeck(deck Deck) Deck{
	if len(deck.Cards) == 0 {
		fmt.Println("Deck is empty!")
		return deck
	}

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	rng.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
	return deck
}