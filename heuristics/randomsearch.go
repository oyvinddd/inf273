package heuristics

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
)

const maxIterations int = 10000

// RandomSearch randomly searches for a better solution
func RandomSearch(data models.INF273Data, s0 [][]*models.Call) [][]*models.Call {
	best := s0
	for i := 0; i < maxIterations; i++ {
		current := a2.GenerateSolution(data)
		if a2.IsFeasible(data, current) && a2.TotalObjective(data, current) < a2.TotalObjective(data, best) {
			best = current
		}
	}
	return best
}
