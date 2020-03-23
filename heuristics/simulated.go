package heuristics

import (
	"math"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"

	"github.com/oyvinddd/inf273/heuristics/operators"

	"github.com/oyvinddd/inf273/models"
)

// SA (Simulated Annealing) iteratively searches for a better solution
func SA(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	incumbent, best := solution, solution
	var temp float64 = 100
	var a float64 = 0.85
	var p1 float32 = 1.0 / 5.0
	var p2 float32 = 1.0 / 5.0
	for i := 0; i < maxIterations; i++ {

		var random float32 = rand.Float32()
		var current [][]*models.Call = nil

		if random < p1 {
			current = operators.TwoExchange(data, incumbent)
		} else if random < p1+p2 {
			current = operators.ThreeExchange(data, incumbent)
		} else {
			current = operators.OneReinsert(data, solution)
		}

		fNewSolution := a2.CalcTotalObjective(data, current)
		fIncumbent := a2.CalcTotalObjective(data, incumbent)
		deltaE := float64(fNewSolution - fIncumbent)

		if current != nil && isFeasible(data, current) && deltaE < 0 {
			incumbent = current
			if fIncumbent < a2.CalcTotalObjective(data, best) {
				best = incumbent
			}
		} else if current != nil && isFeasible(data, current) && rand.Float64() < math.Exp(-deltaE/temp) {
			incumbent = current
		}
		temp = temp * a
	}
	return best
}
