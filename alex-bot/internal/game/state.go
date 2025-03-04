package game

import (
	"fmt"

	"github.com/amalavet/poker-bot/internal/card"
)

type State struct {
	Deck    *card.Deck
	Players []*Player
}

func (s *State) String() string {
	var str string
	for _, p := range s.Players {
		str += p.Name + ": "
		for _, c := range p.Hand {
			str += c.String() + " "
		}
		str += "\n"
	}
	str += "Deck: " + s.Deck.String()
	return str
}

type Player struct {
	Name   string
	Chips  int
	Hand   [2]*card.Card
	Folded bool
}

func NewState() *State {
	return &State{
		Deck:    card.NewDeck(),
		Players: []*Player{},
	}
}

func (s *State) AddPlayer(name string, chips int) {
	s.Players = append(s.Players, &Player{
		Name:   name,
		Chips:  chips,
		Folded: false,
	})
}

func (s *State) Deal() error {
	if !s.Deck.IsFull() {
		return fmt.Errorf("there are missing cards in the deck")
	}
	s.Deck.Shuffle()
	for i, p := range s.Players {
		p.Hand[0] = s.Deck[i*2]
		p.Hand[1] = s.Deck[i*2+1]
		s.Deck[i*2] = nil
		s.Deck[i*2+1] = nil
	}
	return nil
}
