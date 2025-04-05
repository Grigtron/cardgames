package deck

import "fmt"

type Suit int
const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Rank int
const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (s Suit) String() string {
	switch {
	case s == 0:
		return "\u2663"
	case s == 1:
		return "\u2666"
	case s == 2:
		return  "\u2665"
	case s == 3:
		return "\u2660"
	default:
		return "Unknown Suit(Error)"
	}

}

func (r Rank) String() string {
	
	switch {
	case r == 11:
		return "Jack"
	case r == 12:
		return "Queen"
	case r == 13:
		return "King"
	case r == 14:
		return "Ace"
	default:
		return fmt.Sprintf("%d", r)
	}
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}


