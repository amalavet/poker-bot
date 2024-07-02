package main

import (
	"fmt"

	"github.com/amalavet/poker-bot/internal/game"
)

func main() {
	state := game.NewState()
	state.AddPlayer("Alice", 1000)
	state.AddPlayer("Bob", 1000)
	state.Deal()
	fmt.Println(state)
}
