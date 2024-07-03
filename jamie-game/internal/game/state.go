package game

import (
	"fmt"
	"math/rand"
)

const (
	NUM_DIRECTIONS = 6
	NUM_HEXES      = 37
)

type Team string

const (
	NONE  Team = ""
	RED   Team = "Red"
	BLACK Team = "Black"
)

type Board [NUM_HEXES]*Hex

type State struct {
	board   Board
	turn    Team
	score   int
	history []*History
}

func (s *State) LenHistory() int {
	return len(s.history)
}

type History struct {
	move    *Move
	capture *Piece
	win     Team
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
	id        int
}

func NewMove(action action, hex int, target int) *Move {
	return &Move{action, hex, target}
}

func (s *State) IsValidMove(move *Move) bool {
	for _, validMove := range s.GenerateValidMoves() {
		if move.hex == validMove.hex &&
			move.target == validMove.target &&
			move.action == validMove.action {
			return true
		}
	}
	return false
}

const WIN_SCORE = 1000

func (s *State) Move(m *Move) {
	hex := s.board[m.hex]
	history := &History{m, nil, NONE}
	switch m.action {
	case TRAVEL:
		target := hex.neighbors[m.target]
		if target.piece != nil {
			history.capture = target.piece
			if target.piece.team == RED {
				s.score--
			} else {
				s.score++
			}
		}
		target.piece = hex.piece
		hex.piece = nil
		if target.id == RED_BASE && s.turn == BLACK {
			s.score -= WIN_SCORE
			history.win = BLACK
		}
		if target.id == BLACK_BASE && s.turn == RED {
			s.score += WIN_SCORE
			history.win = RED
		}
	case ROTATE:
		hex.piece.Rotate(m.target)
	}
	s.history = append(s.history, history)
	if s.turn == RED {
		s.turn = BLACK
	} else {
		s.turn = RED
	}
}

func (s *State) Undo() {
	if len(s.history) == 0 {
		return
	}
	history := s.history[len(s.history)-1]
	s.history = s.history[:len(s.history)-1]
	hex := s.board[history.move.hex]
	switch history.move.action {
	case TRAVEL:
		target := hex.neighbors[history.move.target]
		hex.piece = target.piece
		target.piece = history.capture
	case ROTATE:
		hex.piece.Rotate(NUM_DIRECTIONS - history.move.target)
	}

	if history.capture != nil {
		if history.capture.team == RED {
			s.score++
		} else {
			s.score--
		}
	}

	if history.win == RED {
		s.score -= WIN_SCORE
	} else if history.win == BLACK {
		s.score += WIN_SCORE
	}

	if s.turn == RED {
		s.turn = BLACK
	} else {
		s.turn = RED
	}
}

func (s *State) Evaluate() int {
	return s.score
}

func (s *State) RandomMove() {
	moves := s.GenerateValidMoves()
	fmt.Println(moves)
	move := moves[rand.Intn(len(moves))]
	s.Move(move)
}

func (s *State) GenerateValidMoves() []*Move {
	var moves []*Move
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

			if n.id == RED_BASE && s.turn == RED {
				continue
			}

			if n.id == BLACK_BASE && s.turn == BLACK {
				continue
			}

			if n.piece == nil {
				moves = append(moves, &Move{TRAVEL, i, j})
			} else if n.piece.team != s.turn &&
				!n.piece.arrows[(j+3)%NUM_DIRECTIONS] {
				moves = append(moves, &Move{TRAVEL, i, j})
			}
		}

		// Add all possible rotations
		moves = append(moves, &Move{ROTATE, i, 1})
		moves = append(moves, &Move{ROTATE, i, 2})
		moves = append(moves, &Move{ROTATE, i, 3})
		moves = append(moves, &Move{ROTATE, i, 4})
		moves = append(moves, &Move{ROTATE, i, 5})

	}
	return moves
}

type Move struct {
	action action
	hex    int
	target int
}

type action string

const (
	TRAVEL action = "TRAVEL"
	ROTATE action = "ROTATE"
)
