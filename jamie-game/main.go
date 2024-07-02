package main

import "fmt"

func main() {
	game := NewGame()
	fmt.Println(game)
}

const (
	NUM_ROTATIONS = 6
	NUM_HEXES     = 37
)

type Team bool

const (
	RED   Team = true
	BLACK Team = false
)

type Board [NUM_HEXES]*Hex

type Piece struct {
	team   Team
	arrows [NUM_ROTATIONS]bool
}

func (p *Piece) Rotate(n int) {
	var temp [NUM_ROTATIONS]bool
	for i := 0; i < NUM_ROTATIONS; i++ {
		temp[(i+n)%NUM_ROTATIONS] = p.arrows[i]
	}

	p.arrows = temp
}

type Hex struct {
	piece     *Piece
	neighbors [NUM_ROTATIONS]*Hex
}

func NewGame() Board {
	board := [NUM_HEXES]*Hex{
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
		{
			piece: nil,
		},
	}

	// Hex 0
	board[0].neighbors[0] = board[2]
	board[0].neighbors[1] = board[3]
	board[0].neighbors[5] = board[1]

	// Hex 1
	board[1].neighbors[0] = board[5]
	board[1].neighbors[1] = board[2]
	board[1].neighbors[2] = board[0]
	board[1].neighbors[5] = board[4]
	board[1].piece = &Piece{
		team:   RED,
		arrows: [NUM_ROTATIONS]bool{true, false, false, true, false, false},
	}

	// Hex 2
	board[2].neighbors[0] = board[6]
	board[2].neighbors[1] = board[7]
	board[2].neighbors[2] = board[3]
	board[2].neighbors[3] = board[0]
	board[2].neighbors[4] = board[1]
	board[2].neighbors[5] = board[5]
	board[2].piece = &Piece{
		team:   RED,
		arrows: [NUM_ROTATIONS]bool{true, false, false, false, false, false},
	}

	// Hex 3
	board[3].neighbors[0] = board[7]
	board[3].neighbors[1] = board[8]
	board[3].neighbors[4] = board[0]
	board[3].neighbors[5] = board[2]
	board[3].piece = &Piece{
		team:   RED,
		arrows: [NUM_ROTATIONS]bool{true, false, false, true, false, false},
	}

	return board
}

func (b Board) String() string {
	boardString := ""
	for i := range b {
		fmt.Printf(pieceString, i)
	}

	return boardString
}

const pieceString = `
         /%s\
      /%s\  /%s\
   /%s\  /%s\  /%s\
/%s\  /%s\  /%s\  /%s\
\  /%s\  /%s\  /%s\  /
/%s\  /%s\  /%s\  /%s\    
\  /%s\  /%s\  /%s\  /
/%s\  /%s\  /%s\  /%s\
\  /%s\  /%s\  /%s\  /
/%s\  /%s\  /%s\  /%s\    
\  /%s\  /%s\  /%s\  /
   \  /%s\  /%s\  /
      \  /%s\  /
         \  /
`
