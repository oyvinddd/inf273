package heuristics

import (
	"math"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"

	"github.com/oyvinddd/inf273/heuristics/operators"

	"github.com/oyvinddd/inf273/models"
)

// SA (Simulated Annealing) iteratively searches for a better solution
func SA(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int) {
	incumbent := solution
	best := solution
	bestObj := a2.CalcTotalObjective(data, best)

	var temp float64 = 1000
	var a float64 = 0.998
	var p1 float32 = 0.2
	var p2 float32 = 0.3

	for i := 0; i < maxIterations; i++ {

		var random float32 = rand.Float32()
		var newSolution [][]*models.Call = nil

		if random < p1 {
			newSolution = operators.TwoExchange(data, incumbent)
		} else if random < p1+p2 {
			newSolution = operators.ThreeExchange(data, incumbent)
		} else {
			newSolution = operators.OneReinsert(data, incumbent)
		}

		fNewSolution := a2.CalcTotalObjective(data, newSolution)
		fIncumbent := a2.CalcTotalObjective(data, incumbent)
		deltaE := float64(fNewSolution - fIncumbent)
		p := math.Exp(-deltaE / temp)
		math.Exp(-(2200 - 2000) / 50)

		if newSolution != nil && isFeasible(data, newSolution) && deltaE < 0 {
			incumbent = newSolution
			if fIncumbent < a2.CalcTotalObjective(data, best) {
				best = incumbent
				bestObj = fIncumbent
			}
		} else if newSolution != nil && isFeasible(data, newSolution) && rand.Float64() < p {
			incumbent = newSolution
		}
		temp = temp * a
	}
	return best, bestObj
}
