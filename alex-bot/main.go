package main

import (
	"fmt"

	"github.com/amalavet/poker-bot/internal/card"
)

func main() {
	c := card.NewCard(card.Spades, card.Three)
	fmt.Println(c)
}
