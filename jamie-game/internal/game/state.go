package game

const (
	NUM_DIRECTIONS = 6
	NUM_HEXES      = 37
)

type Team string

const (
	RED   Team = "Red"
	BLACK Team = "Black"
)

type Board [NUM_HEXES]*Hex

type State struct {
	board Board
	turn  Team
}

type Piece struct {
	team   Team
	arrows [NUM_DIRECTIONS]bool
}

func (p *Piece) Rotate(n int) {
	var temp [NUM_DIRECTIONS]bool
	for i := 0; i < NUM_DIRECTIONS; i++ {
		temp[(i+n)%NUM_DIRECTIONS] = p.arrows[i]
	}

	p.arrows = temp
}

type Hex struct {
	piece     *Piece
	neighbors [NUM_DIRECTIONS]*Hex
}

func (s *State) GenerateMoves() []Move {
	var moves []Move
	for i, hex := range s.board {
		if hex.piece == nil {
			continue
		}
		if hex.piece.team != s.turn {
			continue
		}

		for j, validDirection := range hex.piece.arrows {
			if !validDirection {
				continue
			}

			n := hex.neighbors[j]
			if n == nil {
				continue
			}

			if n.piece == nil {
				moves = append(moves, Move{TRAVEL, i, j})
			} else if n.piece.team != s.turn {
				moves = append(moves, Move{TRAVEL, i, j})
			}
		}

		// // Add all possible rotations
		// moves = append(moves, Move{ROTATE, i, 1})
		// moves = append(moves, Move{ROTATE, i, 2})
		// moves = append(moves, Move{ROTATE, i, 3})
		// moves = append(moves, Move{ROTATE, i, 4})
		// moves = append(moves, Move{ROTATE, i, 5})

	}
	return moves
}

type Move struct {
	moveType  MoveType
	hex       int
	direction int
}

type MoveType string

const (
	TRAVEL MoveType = "TRAVEL"
	ROTATE MoveType = "ROTATE"
)
