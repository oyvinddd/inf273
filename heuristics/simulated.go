package heuristics

import (
	"fmt"
	"math"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"

	"github.com/oyvinddd/inf273/heuristics/operators"

	"github.com/oyvinddd/inf273/models"
)

// SA (Simulated Annealing) iteratively searches for a better solution
func SA(data models.INF273Data, s0 [][]*models.Call) [][]*models.Call {
	incumbent, best := s0, s0

	var deSum float64 = 0
	var deNum float64 = 0

	var T float64 = 1000     // temperature
	var a float64 = 0.998765 // cooling factor
	var p float64 = 0.8      // probability of accepting worse solution
	var p1 float32 = 0.14    // probability of using 2-exchange
	var p2 float32 = 0.15    // probability of using 3-exchange

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

		deltaE := float64(a2.TotalObjective(data, newSolution) - a2.TotalObjective(data, incumbent))

		if i < 100 && deltaE >= 0 {
			deSum += deltaE
			deNum++
		} else if i == 100 {
			avgDeltas := deSum / deNum
			T = -avgDeltas / math.Log(0.8)
			fmt.Printf("Avg delta: %v\n", avgDeltas)
		} else {
			p = math.Exp(-deltaE / T)
		}

		isFeasible := a2.IsFeasible(data, newSolution)

		if isFeasible && deltaE < 0 {
			incumbent = newSolution
			if a2.TotalObjective(data, incumbent) < a2.TotalObjective(data, best) {
				best = incumbent
			}
		} else if isFeasible && rand.Float64() < p {
			incumbent = newSolution
		}
		T *= a
	}
	return best
}
