package game

import (
	"fmt"
	"github.com/Grigtron/cardgames/deck"
	"math/rand"
	"time"
)

type WarGame struct {
	PlayerDeck deck.Deck
	ComputerDeck deck.Deck
}

func NewWarGame() (*WarGame, error) {
	d := deck.NewDeck()
	d = deck.ShuffleDeck(d)
	
	playerDeck, computerDeck := DealWarCards(d)
	if len(playerDeck.Cards) == 0 || len(computerDeck.Cards) == 0 {
		return nil, fmt.Errorf("error creating decks. playerdeck = %d. computerdeck = %d", len(playerDeck.Cards), len(computerDeck.Cards))
	}
	
	return &WarGame{
		PlayerDeck: playerDeck,
		ComputerDeck: computerDeck,
	}, nil

}

func DealWarCards(d deck.Deck) (deck.Deck, deck.Deck) {
	playerDeck := deck.Deck{Cards: []deck.Card{}}
	computerDeck := deck.Deck{Cards: []deck.Card{}}

	for i := 0; i < len(d.Cards); i += 2 {
		playerDeck.Cards = append(playerDeck.Cards, d.Cards[i])
	}
	for i := 1; i < len(d.Cards); i += 2 {
		computerDeck.Cards = append(computerDeck.Cards, d.Cards[i])
	}

	return playerDeck, computerDeck
}

func (game *WarGame) PlayTurn() error {
	playerDeck := &game.PlayerDeck
	computerDeck := &game.ComputerDeck

	playerCard := playerDeck.Cards[0]
	computerCard := computerDeck.Cards[0]

	playerDeck.Cards = playerDeck.Cards[1:]
	computerDeck.Cards = computerDeck.Cards[1:]

	if playerCard.Rank == computerCard.Rank {
		fmt.Printf("Player card: %v\nComputer card: %v\n\n", playerCard, computerCard)
		fmt.Println("Prepare for War!")

		warCount := min(min(len(playerDeck.Cards), len(computerDeck.Cards)), 3)

		playerWarHand := playerDeck.Cards[:warCount]
		computerWarHand := computerDeck.Cards[:warCount]

		playerDeck.Cards = playerDeck.Cards[warCount:]
		computerDeck.Cards = computerDeck.Cards[warCount:]

		playerWarCard := playerWarHand[len(playerWarHand)-1]
		computerWarCard := computerWarHand[len(computerWarHand)-1]

		if playerWarCard.Rank > computerWarCard.Rank {
			fmt.Printf("Player card: %v\nComputer card: %v\nPlayer wins!\n", playerWarCard, computerWarCard)
			playerDeck.Cards = append(playerDeck.Cards, computerWarHand...)
			playerDeck.Cards = append(playerDeck.Cards, playerWarHand...)
			
		} else if playerWarCard.Rank < computerWarCard.Rank {
			fmt.Printf("Player card: %v\nComputer card: %v\nComputer wins!\n", playerWarCard, computerWarCard)
			computerDeck.Cards = append(computerDeck.Cards, playerWarHand...)
			computerDeck.Cards = append(computerDeck.Cards, computerWarHand...)
		} else {
			rng := rand.New(rand.NewSource(time.Now().UnixNano()))
			if rng.Intn(2) == 0 {
				fmt.Println("Player wins the war by luck!")
				playerDeck.Cards = append(playerDeck.Cards, computerWarHand...)
				playerDeck.Cards = append(playerDeck.Cards, playerWarHand...)
			} else {
				fmt.Println("Computer wins the war by luck!")
				computerDeck.Cards = append(computerDeck.Cards, playerWarHand...)
				computerDeck.Cards = append(computerDeck.Cards, computerWarHand...)
			}

		}
	} else if playerCard.Rank > computerCard.Rank {
		fmt.Printf("Player card: %v\nComputer card: %v\n", playerCard, computerCard)
		fmt.Println("Player wins!")
		playerDeck.Cards = append(playerDeck.Cards, playerCard, computerCard)
	} else {
		fmt.Printf("Player card: %v\nComputer card: %v\n", playerCard, computerCard)
		fmt.Println("Computer wins!")
		computerDeck.Cards = append(computerDeck.Cards, computerCard, playerCard)
	}

	if len(playerDeck.Cards) == 0 {
		fmt.Println("Computer is the winner!")
	}
	if len(computerDeck.Cards) == 0 {
		fmt.Println("Player is the winner!")
	}

	return nil
}

