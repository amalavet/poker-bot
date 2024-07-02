package card

type (
	suit int
	rank int
)

const (
	Spades suit = iota + 1
	Hearts
	Diamonds
	Clubs
)

const (
	Ace rank = iota + 1
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
	suit suit
	rank rank
}

func NewCard(s suit, r rank) Card {
	return Card{s, r}
}
