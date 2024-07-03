package main

import (
	"fmt"
	"strings"
)

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
	var board Board
	for i := range board {
		board[i] = &Hex{}
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
	f := boardStr
	guide := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "~",
	}
	for i, hex := range b {
		fill := []string{" ", " ", " ", " ", " ", " ", " "}
		if hex.piece != nil {
			fill = hex.piece.P()
		}
		for _, k := range fill {
			f = strings.Replace(f, guide[i], k, 1)
		}
	}
	return f
}

func (p *Piece) P() []string {
	out := []string{" ", " ", " ", " ", " ", " ", " "}

	if p.team == RED {
		out[3] = "R"
	} else {
		out[3] = "B"
	}

	if p.arrows[0] {
		out[1] = "^"
	}
	if p.arrows[1] {
		out[2] = ">"
	}
	if p.arrows[2] {
		out[6] = ">"
	}
	if p.arrows[3] {
		out[5] = "v"
	}
	if p.arrows[4] {
		out[4] = "<"
	}
	if p.arrows[5] {
		out[0] = "<"
	}

	return out
}

const boardStr = `
                          -------
                         / x x x \
                  -------    x    -------
                 / w w w \ x x x / y y y \
          -------    w    -------    y    -------
         / v v v \ w w w / q q q \ y y y / z z z \
  -------    v    -------    q    -------    z    -------  
 / u u u \ v v v / p p p \ q q q / r r r \ z z z / ~ ~ ~ \
/    u    -------    p    -------    r    -------    ~    \
 \ u u u / o o o \ p p p / j j j \ r r r / s s s \ ~ ~ ~ /
  -------    o    -------    j    -------    s    -------
 / n n n \ o o o / i i i \ j j j / k k k \ s s s / t t t \
/    n    -------    i    -------    k    -------    t    \
 \ n n n / h h h \ i i i / c c c \ k k k / l l l \ t t t /
  -------    h    -------    c    -------    l    -------
 / g g g \ h h h / b b b \ c c c / d d d \ l l l / m m m \
\    g    -------    b    -------    d    -------    m    /
 \ g g g / a a a \ b b b / 6 6 6 \ d d d / e e e \ m m m /
  -------    a    -------    6    -------    e    -------
 / 9 9 9 \ a a a / 5 5 5 \ 6 6 6 / 7 7 7 \ e e e / f f f \
\    9    -------    5    -------    7    -------    f    /
 \ 9 9 9 / 4 4 4 \ 5 5 5 / 2 2 2 \ 7 7 7 / 8 8 8 \ f f f /
  -------    4    -------    2    -------    8    ------- 
         \ 4 4 4 / 1 1 1 \ 2 2 2 / 3 3 3 \ 8 8 8 /
          -------    1    -------    3    -------
                 \ 1 1 1 / 0 0 0 \ 3 3 3 /
                  -------    0    -------
                         \ 0 0 0 /   
                          -------
`
