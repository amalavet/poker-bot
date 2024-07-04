package game

import (
	"fmt"
	"math"
	"math/rand"
)

var seen map[string]struct{}

func (s *State) ExecuteBestMove(depth int8) {
	bestScore := -100000.0
	var bestMoves []*Move
	seen = make(map[string]struct{})

	for _, move := range s.GenerateValidMoves() {
		s.Move(move)
		score := s.alphaBeta(depth, -100000, 100000)
		s.Undo()
		if s.turn == BLACK {
			score = -score
		}
		if score == bestScore {
			bestMoves = append(bestMoves, move)
		} else if score > bestScore {
			bestMoves = []*Move{move}
			bestScore = score
		}
	}
	s.Move(bestMoves[rand.Intn(len(bestMoves))])
}

func (h *Hex) String() string {
	out := fmt.Sprintf("_%d_", h.id)
	if h.piece != nil {
		out += h.piece.team.String() + "_" + fmt.Sprintf("%v", h.piece.arrows)
	}
	return out
}

func (s *State) Hash() string {
	hash := s.turn.String()
	for _, hex := range s.board {
		hash += hex.String()
	}
	return hash
}

func (s *State) Evaluate() float64 {
	score := 0
	for _, hex := range s.board {
		if hex.piece == nil {
			continue
		}
		score += hex.value
		if hex.piece.team == RED {
			score += hex.piece.value
		} else {
			score -= hex.piece.value
		}
	}
	return float64(score)
}

func (s *State) alphaBeta(depth int8, a, b float64) float64 {
	if depth == 0 {
		return s.Evaluate()
	}

	hash := s.Hash()
	if _, ok := seen[hash]; ok {
		return s.Evaluate()
	}
	seen[hash] = struct{}{}

	moves := s.GenerateValidMoves()
	if s.turn == RED {
		value := -100000.0
		for _, move := range moves {
			s.Move(move)
			value = math.Max(value, s.alphaBeta(depth-1, a, b))
			s.Undo()
			if value > b {
				break
			}
			a = math.Max(a, value)
		}
		return value
	} else {
		value := 100000.0
		for _, move := range moves {
			s.Move(move)
			value = math.Min(value, s.alphaBeta(depth-1, a, b))
			s.Undo()
			if value < a {
				break
			}
			b = math.Min(b, value)
		}
		return value
	}
}
