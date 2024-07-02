package card

type (
	suit int
	rank int
)

const (
	Spades suit = iota
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

func (c Card) String() string {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"♠", "♥", "♦", "♣"}
	return ranks[c.rank-1] + suits[c.suit]
}
