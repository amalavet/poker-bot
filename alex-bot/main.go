package main

import "github.com/amalavet/poker-bot/internal/card"

func main() {
	c := card.NewCard(card.Spades, card.Three)
	println(c)
}
