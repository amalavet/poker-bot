package main

import (
	"fmt"

	"github.com/amalavet/poker-bot/jamie-game/internal/game"
)

func main() {
	state := game.NewGame()
	for {
		fmt.Println(state)
		var choice int
		fmt.Println("0: Random Move")
		fmt.Println("1: Undo")
		fmt.Println("2: Random Game")
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)
		if choice == 0 {
			state.RandomMove()
		} else if choice == 1 {
			state.Undo()
		} else if choice == 2 {
			for len(state.GenerateMoves()) > 0 {
				state.RandomMove()
				fmt.Println(state)
			}
			for state.LenHistory() > 0 {
				state.Undo()
				fmt.Println(state)
			}
		}
	}
}
