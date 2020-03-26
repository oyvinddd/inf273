package heuristics

import (
	"math"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"

	"github.com/oyvinddd/inf273/heuristics/operators"

	"github.com/oyvinddd/inf273/models"
)

// SA (Simulated Annealing) iteratively searches for a better solution
func SA(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, []float64, []float64, []float64) {
	incumbent := solution
	best := solution

	var temp float64 = 1000 // temperature
	var a float64 = 0.99985 // cooling factor
	var p1 float32 = 0.1    // probability of using 2-exchange
	var p2 float32 = 0.05   // probability of using 3-exchange

	var x []float64 = make([]float64, 50000)
	var y []float64 = make([]float64, 50000)
	var pp []float64 = make([]float64, 50000)

	for i := 0; i < maxIterations; i++ {
		x[i] = float64(i)

		var random float32 = rand.Float32()
		var newSolution [][]*models.Call = nil

		if random < p1 {
			newSolution = operators.TwoExchange(data, incumbent)
		} else if random < p1+p2 {
			newSolution = operators.ThreeExchange(data, incumbent)
		} else {
			newSolution = operators.OneReinsert(data, incumbent)
		}

		deltaE := float64(a2.TotalObjective(data, newSolution) - a2.TotalObjective(data, incumbent))
		feasible := a2.IsFeasible(data, newSolution)

		if feasible && deltaE < 0 {
			incumbent = newSolution
			if newObj := a2.TotalObjective(data, incumbent); newObj < a2.TotalObjective(data, best) {
				best = incumbent
			}
		} else if feasible && rand.Float64() < math.Exp(-deltaE/temp) {
			incumbent = newSolution
		}
		temp = temp * a
		y[i] = temp
	}

	return best, x, y, pp
}
