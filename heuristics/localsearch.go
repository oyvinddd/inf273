package heuristics

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
)

// LocalSearch does a local search on a given solution
func LocalSearch(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	current, best := solution, solution
	var p1 float32 = 1 / 3
	var p2 float32 = 1 / 3
	for i := 0; i < maxIterations; i++ {
		r := rand.Float32()
		if r < p1 {
			current = TwoExchange(best)
		} else if r < p1+p2 {
			current = ThreeExchange(best)
		} else {
			current = OneReinsert(best)
		}
		if isFeasible(data, current) && a2.CalcTotalObjective(data, current) < a2.CalcTotalObjective(data, best) {
			best = current
		}
	}
	return best
}
