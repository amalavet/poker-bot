package main

import (
	"fmt"

	"github.com/amalavet/poker-bot/jamie-game/internal/game"
)

func main() {
	state := game.NewGame()
	fmt.Println(state)

	moves := state.GenerateMoves()
	fmt.Println(moves)
}
