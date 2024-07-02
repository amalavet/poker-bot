package main

type (
	Suit int
	Rank int
)

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

const (
	Ace Rank = iota
	Two
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
)

type Card struct {
	Suit Suit
	Rank Rank
}

func main() {
	card := Card{Suit: Spades, Rank: Ace}
	println(card)
}
