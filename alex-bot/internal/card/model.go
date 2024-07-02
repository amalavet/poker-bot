package card

import (
	"math/rand"
)

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

type Deck [52]*Card

func NewDeck() *Deck {
	d := Deck{}
	for s := Spades; s <= Clubs; s++ {
		for r := Ace; r <= King; r++ {
			d[int(s)*13+int(r)-1] = NewCard(s, r)
		}
	}
	return &d
}

func (d *Deck) String() string {
	var s string
	for i, c := range d {
		if i%13 == 0 {
			s += "\n"
		}
		if c == nil {
			s += "XX "
		} else {
			s += c.String() + " "
		}
	}
	return s
}

func (d *Deck) Shuffle() {
	for i := range d {
		j := i + rand.Intn(len(d)-i)
		d[i], d[j] = d[j], d[i]
	}
}

func (d *Deck) IsFull() bool {
	for _, c := range d {
		if c == nil {
			return false
		}
	}
	return true
}

type Hand [5]*Card

type Card struct {
	suit suit
	rank rank
}

func NewCard(s suit, r rank) *Card {
	return &Card{s, r}
}

func (c Card) String() string {
	ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	suits := []string{"♠", "♥", "♦", "♣"}
	return ranks[c.rank-1] + suits[c.suit]
}

func (c Card) Rank() rank {
	return c.rank
}

func (c Card) Suit() suit {
	return c.suit
}
