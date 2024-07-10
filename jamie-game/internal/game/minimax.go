package game

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var (
	seen        = make(map[string]float64)
	cacheHits   = 0
	cacheMisses = 0
	evals       int
)

func (s *State) ExecuteBestMove(maxDepth int8) {
	bestScore := math.Inf(-1)
	var bestMoves []*Move
	now := time.Now()

	for _, move := range s.GenerateValidMoves() {
		s.Move(move)
		score := s.alphaBeta(0, maxDepth, math.Inf(-1), math.Inf(1))
		s.Undo()
		if s.turn == BLACK {
			score = -score
		}
		if score == bestScore {
			bestMoves = append(bestMoves, move)
		} else if score > bestScore || len(bestMoves) == 0 {
			bestMoves = []*Move{move}
			bestScore = score
		}
	}
	fmt.Printf("Evals: %d, Time Spent: %v", evals, time.Since(now))
	fmt.Printf("  %.2f evals per second\n", float64(evals)/time.Since(now).Seconds())
	fmt.Printf(
		"Cached moves: %d, Hits: %d, Misses: %d, Hit Rate: %.2f\n",
		len(seen),
		cacheHits,
		cacheMisses,
		100*float64(cacheHits)/float64(cacheHits+cacheMisses),
	)
	if len(bestMoves) > 0 {
		s.Move(bestMoves[rand.Intn(len(bestMoves))])
	}
}

func (h *Hex) String() string {
	out := fmt.Sprintf("_%d_", h.id)
	if h.piece != nil {
		out += h.piece.team.String() + "_" + fmt.Sprintf("%v", h.piece.arrows)
	}
	return out
}

func (s *State) Hash(depth int8) string {
	return fmt.Sprintf("%s_%s_depth:%d", s.turn, s.board, depth)
}

func (s *State) Evaluate() float64 {
	score := 0.0
	for _, hex := range s.board {
		if hex.piece == nil {
			continue
		}
		score += hex.value
		if hex.piece.team == RED {
			score += float64(hex.piece.value)
		} else {
			score -= float64(hex.piece.value)
		}
	}
	return float64(score)
}

func (s *State) alphaBeta(depth int8, maxDepth int8, a, b float64) float64 {
	if depth == maxDepth {
		evals++
		return s.Evaluate()
	}

	hash := s.Hash(maxDepth - depth)
	if value, ok := seen[hash]; ok {
		cacheHits++
		return value
	}
	cacheMisses++

	moves := s.GenerateValidMoves()
	if s.turn == RED {
		value := math.Inf(-1)
		for _, move := range moves {
			s.Move(move)
			value = math.Max(value, s.alphaBeta(depth+1, maxDepth, a, b))
			s.Undo()
			if value > b {
				break
			}
			a = math.Max(a, value)
		}
		seen[hash] = value
		return value
	} else {
		value := math.Inf(1)
		for _, move := range moves {
			s.Move(move)
			value = math.Min(value, s.alphaBeta(depth+1, maxDepth, a, b))
			s.Undo()
			if value < a {
				break
			}
			b = math.Min(b, value)
		}
		seen[hash] = value
		return value
	}
}
