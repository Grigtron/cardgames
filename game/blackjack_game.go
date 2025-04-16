package game

import (
	"fmt"

	"github.com/Grigtron/cardgames/deck"
)

type BlackjackGame struct {
	Shoe deck.Deck
	Discard []deck.Card
	PlayerHand []deck.Card
	DealerHand []deck.Card
	PlayerScore int
	DealerScore int
}

var _ Game = (*BlackjackGame)(nil)

func NewBlackjackGame() (*BlackjackGame, error) {
	d := deck.NewDeck()
	d = deck.ShuffleDeck(d)

	if len(d.Cards) == 0 {
		return nil, fmt.Errorf("error creating blackjack deck: %d", len(d.Cards))
	}
	playerHand := []deck.Card{d.Cards[0], d.Cards[2]}
	dealerHand := []deck.Card{d.Cards[1], d.Cards[3]}
	d.Cards = d.Cards[4:]

	gameInstance := &BlackjackGame{
        Shoe:       d,
        Discard:    []deck.Card{},
        PlayerHand: playerHand,
        DealerHand: dealerHand,
    }

    gameInstance.DisplayHands()
    return gameInstance, nil
}

func (game *BlackjackGame) dealBlackjackCards() error {
	game.maybeReshuffle()

	playerHand := []deck.Card{}
	dealerHand := []deck.Card{}

	for i := 0; i < 2; i++ {
		playerCard, _ := game.drawCard()
		dealerCard, _ := game.drawCard()
		playerHand = append(playerHand, playerCard)
		dealerHand = append(dealerHand, dealerCard)

	}
	game.PlayerHand = playerHand
	game.DealerHand = dealerHand

	game.DisplayHands()
	return nil
} 

func (game *BlackjackGame) DisplayHands() {
	fmt.Println("Player Hand: ")
	for _, card := range game.PlayerHand {
		fmt.Printf("%s ", card.String())
	}
	fmt.Println("\nDealer Hand: ")
	fmt.Printf("%s [hidden]\n", game.DealerHand[0].String())
}

func (game *BlackjackGame) drawCard() (deck.Card, error) {
	game.maybeReshuffle()
	if len(game.Shoe.Cards) == 0 {
		return deck.Card{}, fmt.Errorf("cannot draw card: shoe is empty")
	}
	card := game.Shoe.Cards[0]
	game.Shoe.Cards = game.Shoe.Cards[1:]
	fmt.Printf("Draws a %v\n", card)
	return card, nil
}

func (game *BlackjackGame) Hit() error {
	hitCard, err := game.drawCard()
	if err != nil {
		fmt.Println("Error drawing card:", err)
		return nil
	}

	game.PlayerHand = append(game.PlayerHand, hitCard)
	game.PlayerScore = game.CalculateScore(game.PlayerHand)
	game.DisplayHands()

	if game.PlayerScore > 21 {
		fmt.Printf("Player busts! %d\n", game.PlayerScore)
		game.whoWins()
	}
	return nil
}

func (game *BlackjackGame) Stand() error {
	game.DisplayHands()
	game.PlayerScore = game.CalculateScore(game.PlayerHand)
	fmt.Printf("Player stands at %d\nDealer's turn\n", game.PlayerScore)
	game.dealerTurn()
	return nil
}

func (game *BlackjackGame) dealerTurn() {
	game.maybeReshuffle()
	game.DealerScore = game.CalculateScore(game.DealerHand)
	if game.DealerScore >= 17 {
		fmt.Printf("Dealer stands at %d\n", game.DealerScore)
		game.whoWins()
	} else {
		for game.DealerScore < 17 {
			dealerCard, _ := game.drawCard()
			game.DealerHand = append(game.DealerHand, dealerCard)
			game.DealerScore = game.CalculateScore(game.DealerHand)
			game.DisplayHands()
		}
		if game.DealerScore > 21 {
			game.whoWins()
		}
	}
	game.whoWins()
}

func (game *BlackjackGame) whoWins() {
	game.DisplayHands()
	game.PlayerScore = game.CalculateScore(game.PlayerHand)
	game.DealerScore = game.CalculateScore(game.DealerHand)
	if game.PlayerScore > 21 {
		fmt.Printf("\nPlayer busts! Dealer wins!\nPlayer Score: %d\nDealer Score: %d\n", game.PlayerScore, game.DealerScore)
	} else if game.DealerScore > 21 {
		fmt.Printf("\nDealer busts! Player wins!\nPlayer Score: %d\nDealer Score: %d\n", game.PlayerScore, game.DealerScore)
	} else if game.PlayerScore > game.DealerScore {
		fmt.Printf("\nPlayer wins!\nPlayer Score: %d\nDealer Score: %d\n", game.PlayerScore, game.DealerScore)
	} else if game.DealerScore > game.PlayerScore {
		fmt.Printf("\nDealer wins!\nPlayer Score: %d\nDealer Score: %d\n", game.PlayerScore, game.DealerScore)
	} else if game.PlayerScore == game.DealerScore {
		fmt.Printf("\nIt's a tie!\nPlayer Score: %d\nDealer Score: %d\n", game.PlayerScore, game.DealerScore)
	}
}

func (game *BlackjackGame) CalculateScore(hand []deck.Card) int {
	total := 0
	aceCount := 0

	for _, card := range hand {
		switch {
		case card.Rank >= deck.Two && card.Rank <= deck.Ten:
			total += int(card.Rank)
		case card.Rank >= deck.Jack && card.Rank <= deck.King:
			total += 10
		case card.Rank == deck.Ace:
			total += 11
			aceCount++
		}
	}

	for total > 21 && aceCount > 0 {
		total -= 10
		aceCount--
	}

	return total
}

func (game *BlackjackGame) maybeReshuffle() {
	if len(game.Shoe.Cards) == 1 {
		tempDeck := deck.Deck{Cards: game.Discard}
		lastCard := game.Shoe.Cards[0]
		shuffled := deck.ShuffleDeck(tempDeck)
		game.Discard = []deck.Card{lastCard}
		game.Shoe = shuffled
	}
	fmt.Println("Draw pile depleted! Discarding last Shoe card and reshuffling the discard deck...")
}

func (game *BlackjackGame) HandleCommand(cmd string, args ...string) error{
	switch cmd {
	case "hit":
		return game.Hit()
	case "stand":
		return game.Stand()
	case "deal":
		return game.dealBlackjackCards()
	default:
		return fmt.Errorf("unknown command for Blackjack: %s", cmd)
	}
}

func (game *BlackjackGame) Description() string {
	return "Blackjack: Use 'hit' or 'stand' to play"
}