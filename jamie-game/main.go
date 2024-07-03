package main

import (
	"fmt"
	"math"

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
		fmt.Println("3: Manual Move")
		fmt.Print("Choice: ")
		fmt.Scanln(&choice)
		if choice == 0 {
			state.RandomMove()
		} else if choice == 1 {
			state.Undo()
		} else if choice == 2 {
			for len(state.GenerateValidMoves()) > 0 && math.Abs(float64(state.Evaluate())) < game.WIN_SCORE/2 {
				state.RandomMove()
				fmt.Println(state)
			}
			for state.LenHistory() > 0 {
				state.Undo()
				fmt.Println(state)
			}
		} else if choice == 3 {
			fmt.Print("[0]Move or [1]Rotate?:")
			var action int
			fmt.Scanln(&action)
			fmt.Print("Hex [0-36]:")
			var hex int
			fmt.Scanln(&hex)
			fmt.Print("Direction [0-5]:")
			var direction int
			fmt.Scanln(&direction)

			var move *game.Move
			if action == 0 {
				move = game.NewMove(game.TRAVEL, hex, direction)
			} else {
				move = game.NewMove(game.ROTATE, hex, direction)
			}
			if !state.IsValidMove(move) {
				fmt.Println("Invalid Move")
				continue
			}
			state.Move(move)
		}

	}
}
