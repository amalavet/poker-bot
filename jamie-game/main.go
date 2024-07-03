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

	// Hex 2
	board[2].neighbors[0] = board[6]
	board[2].neighbors[1] = board[7]
	board[2].neighbors[2] = board[3]
	board[2].neighbors[3] = board[0]
	board[2].neighbors[4] = board[1]
	board[2].neighbors[5] = board[5]

	// Hex 3
	board[3].neighbors[0] = board[7]
	board[3].neighbors[1] = board[8]
	board[3].neighbors[4] = board[0]
	board[3].neighbors[5] = board[2]

	// Hex 4
	board[4].neighbors[0] = board[10]
	board[4].neighbors[1] = board[5]
	board[4].neighbors[2] = board[1]
	board[4].neighbors[5] = board[9]

	// Hex 5
	board[5].neighbors[0] = board[11]
	board[5].neighbors[1] = board[6]
	board[5].neighbors[2] = board[2]
	board[5].neighbors[3] = board[1]
	board[5].neighbors[4] = board[4]
	board[5].neighbors[5] = board[10]

	// Hex 6
	board[6].neighbors[0] = board[12]
	board[6].neighbors[1] = board[13]
	board[6].neighbors[2] = board[7]
	board[6].neighbors[3] = board[2]
	board[6].neighbors[4] = board[5]
	board[6].neighbors[5] = board[11]

	// Hex 7
	board[7].neighbors[0] = board[13]
	board[7].neighbors[1] = board[14]
	board[7].neighbors[2] = board[8]
	board[7].neighbors[3] = board[3]
	board[7].neighbors[4] = board[2]
	board[7].neighbors[5] = board[6]

	// Hex 8
	board[8].neighbors[0] = board[14]
	board[8].neighbors[1] = board[15]
	board[8].neighbors[4] = board[3]
	board[8].neighbors[5] = board[7]

	// Hex 9
	board[9].neighbors[0] = board[16]
	board[9].neighbors[1] = board[10]
	board[9].neighbors[2] = board[4]

	// Hex 10
	board[10].neighbors[0] = board[17]
	board[10].neighbors[1] = board[11]
	board[10].neighbors[2] = board[5]
	board[10].neighbors[3] = board[4]
	board[10].neighbors[4] = board[9]
	board[10].neighbors[5] = board[16]

	// Hex 11
	board[11].neighbors[0] = board[18]
	board[11].neighbors[1] = board[12]
	board[11].neighbors[2] = board[6]
	board[11].neighbors[3] = board[5]
	board[11].neighbors[4] = board[10]
	board[11].neighbors[5] = board[17]

	// Hex 12
	board[12].neighbors[0] = board[19]
	board[12].neighbors[1] = board[20]
	board[12].neighbors[2] = board[13]
	board[12].neighbors[3] = board[6]
	board[12].neighbors[4] = board[11]
	board[12].neighbors[5] = board[18]

	// Hex 13
	board[13].neighbors[0] = board[20]
	board[13].neighbors[1] = board[21]
	board[13].neighbors[2] = board[14]
	board[13].neighbors[3] = board[7]
	board[13].neighbors[4] = board[6]
	board[13].neighbors[5] = board[12]

	// Hex 14
	board[14].neighbors[0] = board[21]
	board[14].neighbors[1] = board[22]
	board[14].neighbors[2] = board[15]
	board[14].neighbors[3] = board[8]
	board[14].neighbors[4] = board[7]
	board[14].neighbors[5] = board[13]

	// Hex 15
	board[15].neighbors[0] = board[22]
	board[15].neighbors[4] = board[8]
	board[15].neighbors[5] = board[14]

	// Hex 16
    board[16].neighbors[0] = board[23]
    board[16].neighbors[1] = board[17]
    board[16].neighbors[2] = board[10]
    board[16].neighbors[3] = board[9]

    // Hex 17
    board[17].neighbors[0] = board[24]
    board[17].neighbors[1] = board[18]
    board[17].neighbors[2] = board[11]
    board[17].neighbors[3] = board[10]
    board[17].neighbors[4] = board[16]
    board[17].neighbors[5] = board[23]

    // Hex 18
    board[18].neighbors[0] = board[25]
    board[18].neighbors[1] = board[19]
    board[18].neighbors[2] = board[12]
    board[18].neighbors[3] = board[11]
    board[18].neighbors[4] = board[17]
    board[18].neighbors[5] = board[24]


    // Hex 19
    board[19].neighbors[0] = board[26]
    board[19].neighbors[1] = board[27]
    board[19].neighbors[2] = board[20]
    board[19].neighbors[3] = board[12]
    board[19].neighbors[4] = board[18]
    board[19].neighbors[5] = board[25]

    // Hex 20
    board[20].neighbors[0] = board[27]
    board[20].neighbors[1] = board[28]
    board[20].neighbors[2] = board[21]
    board[20].neighbors[3] = board[13]
    board[20].neighbors[4] = board[12]
    board[20].neighbors[5] = board[19]

    // Hex 21
    board[21].neighbors[0] = board[28]
    board[21].neighbors[1] = board[29]
    board[21].neighbors[2] = board[22]
    board[21].neighbors[3] = board[14]
    board[21].neighbors[4] = board[13]
    board[21].neighbors[5] = board[20]

    // Hex 22
    board[22].neighbors[0] = board[29]
    board[22].neighbors[3] = board[15]
    board[22].neighbors[4] = board[14]
    board[22].neighbors[5] = board[21]

    // Hex 23
    board[23].neighbors[0] = board[30]
    board[23].neighbors[1] = board[24]
    board[23].neighbors[2] = board[17]
    board[23].neighbors[3] = board[16]

    // Hex 24
    board[24].neighbors[0] = board[31]
    board[24].neighbors[1] = board[25]
    board[24].neighbors[2] = board[18]
    board[24].neighbors[3] = board[17]
    board[24].neighbors[4] = board[23]
    board[24].neighbors[5] = board[30]

    // Hex 25
    board[25].neighbors[0] = board[32]
    board[25].neighbors[1] = board[26]
    board[25].neighbors[2] = board[19]
    board[25].neighbors[3] = board[18]
    board[25].neighbors[4] = board[24]
    board[25].neighbors[5] = board[31]

    // Hex 26
    board[26].neighbors[0] = board[33]
    board[26].neighbors[1] = board[34]
    board[26].neighbors[2] = board[27]
    board[26].neighbors[3] = board[19]
    board[26].neighbors[4] = board[25]
    board[26].neighbors[5] = board[32]

    // Hex 27
    board[27].neighbors[0] = board[34]
    board[27].neighbors[1] = board[35]
    board[27].neighbors[2] = board[28]
    board[27].neighbors[3] = board[20]
    board[27].neighbors[4] = board[19]
    board[27].neighbors[5] = board[26]

    // Hex 28
    board[28].neighbors[0] = board[35]
    board[28].neighbors[1] = board[36]
    board[28].neighbors[2] = board[29]
    board[28].neighbors[3] = board[21]
    board[28].neighbors[4] = board[20]
    board[28].neighbors[5] = board[27]

    // Hex 29
    board[29].neighbors[0] = board[36]
    board[29].neighbors[3] = board[22]
    board[29].neighbors[4] = board[21]
    board[29].neighbors[5] = board[28]

    // Hex 30
    board[30].neighbors[1] = board[31]
    board[30].neighbors[2] = board[24]
    board[30].neighbors[3] = board[23]

    // Hex 31
    board[31].neighbors[1] = board[32]
    board[31].neighbors[2] = board[25]
    board[31].neighbors[3] = board[24]
    board[31].neighbors[4] = board[30]

    // Hex 32
    board[32].neighbors[1] = board[33]
    board[32].neighbors[2] = board[26]
    board[32].neighbors[3] = board[25]
    board[32].neighbors[4] = board[31]

    // Hex 33
    board[33].neighbors[2] = board[34]
    board[33].neighbors[3] = board[26]
    board[33].neighbors[4] = board[32]

    // Hex 34
    board[34].neighbors[2] = board[35]
    board[34].neighbors[3] = board[27]
    board[34].neighbors[4] = board[26]
    board[34].neighbors[5] = board[33]

    // Hex 35
    board[35].neighbors[2] = board[36]
    board[35].neighbors[3] = board[28]
    board[35].neighbors[4] = board[27]
    board[35].neighbors[5] = board[34]

    // Hex 36
    board[36].neighbors[3] = board[29]
    board[36].neighbors[4] = board[28]
    board[36].neighbors[5] = board[35]

	return board
}

func (b Board) String() string {
	f := boardStr
	guide := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "~",
	}
	for i, hex := range b {
		if hex.piece == nil {
			continue
		}
		fill := hex.piece.Fill()
		for _, k := range fill {
			f = strings.Replace(f, guide[i], k, 1)
		}
	}
	return f
}

func (p *Piece) Fill() []string {
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
