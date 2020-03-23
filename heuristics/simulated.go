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
	incumbent, bestSolution := solution, solution
	var temp float64 = 100
	var a float64 = 0.85
	var p1 float32 = 1.0 / 4.0
	var p2 float32 = 1.0 / 4.0
	for i := 0; i < maxIterations; i++ {

		var random float32 = rand.Float32()
		var newSolution [][]*models.Call = nil

		if random < p1 {
			newSolution = operators.TwoExchange(data, incumbent)
		} else if random < p1+p2 {
			newSolution = operators.ThreeExchange(data, incumbent)
		} else {
			newSolution = operators.OneReinsert(data, solution)
		}

		fNewSolution := a2.CalcTotalObjective(data, newSolution)
		fIncumbent := a2.CalcTotalObjective(data, incumbent)
		deltaE := float64(fNewSolution - fIncumbent)

		if newSolution != nil && isFeasible(data, newSolution) && deltaE < 0 {
			incumbent = newSolution
			if fIncumbent < a2.CalcTotalObjective(data, bestSolution) {
				bestSolution = incumbent
			}
		} else if newSolution != nil && isFeasible(data, newSolution) && rand.Float64() < math.Exp(-deltaE/temp) {
			incumbent = newSolution
		}
		temp = temp * a
	}
	return bestSolution
}
