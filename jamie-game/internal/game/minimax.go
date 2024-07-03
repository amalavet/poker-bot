package game

import "math"

func (s *State) ExecuteBestMove(depth int8) {
	bestScore := -10000.0
	var bestMove *Move

	for _, move := range s.GenerateValidMoves() {
		s.Move(move)
		score := s.alphaBeta(depth, -10000, 10000)
		s.Undo()
		if s.turn == BLACK {
			score = -score
		}
		if score > bestScore {
			bestMove = move
			bestScore = score
		}
	}
	s.Move(bestMove)
}

func (s *State) alphaBeta(depth int8, a, b float64) float64 {
	if depth == 0 {
		return float64(s.score)
	}

	if s.turn == RED {
		value := -10000.0
		for _, move := range s.GenerateValidMoves() {
			s.Move(move)
			value = math.Max(value, s.alphaBeta(depth-1, a, b))
			s.Undo()
			if value > b {
				return value
			}
			a = math.Max(a, value)
		}
		return value
	} else {
		value := 10000.0
		for _, move := range s.GenerateValidMoves() {
			s.Move(move)
			value = math.Min(value, s.alphaBeta(depth-1, a, b))
			s.Undo()
			if value < a {
				return value
			}
			b = math.Min(b, value)
		}
		return value
	}
}
