package heuristics

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/models"
)

// LocalSearch does a local search on a given solution
func LocalSearch(data models.INF273Data, s0 [][]*models.Call) [][]*models.Call {
	current, best := s0, s0
	var p1 float32 = 0.25
	var p2 float32 = 0.1666
	for i := 0; i < maxIterations; i++ {
		random := rand.Float32()
		if random < p1 {
			current = operators.TwoExchange(data, best)
		} else if random < p1+p2 {
			current = operators.ThreeExchange(data, best)
		} else {
			current = operators.OneReinsert(data, best)
		}
		if a2.IsFeasible(data, current) && a2.TotalObjective(data, current) < a2.TotalObjective(data, best) {
			best = current
		}
	}
	return best
}
