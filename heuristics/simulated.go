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
func SA(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int, []float64, []float64, []float64) {
	incumbent := solution
	best := solution
	bestObj := a2.CalcTotalObjective(data, best)

	var temp float64 = 1000 // temperature
	var a float64 = 0.9985  // cooling factor
	var p1 float32 = 0.1    // probability of using 2-exchange
	var p2 float32 = 0.05   // probability of using 3-exchange

	var x []float64 = make([]float64, 50000)
	var y []float64 = make([]float64, 50000)
	var pp []float64 = make([]float64, 50000)

	for i := 0; i < 50000; i++ {
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

		deltaE := float64(a2.CalcTotalObjective(data, newSolution) - a2.CalcTotalObjective(data, incumbent))
		feasible := isFeasible(data, newSolution)

		if feasible && deltaE < 0 {
			//fmt.Printf("LESS than 0, better solution found! %v\n", deltaE)
			incumbent = newSolution
			if newObj := a2.CalcTotalObjective(data, incumbent); newObj < a2.CalcTotalObjective(data, best) {
				best = incumbent
				bestObj = newObj
			}
		} else if feasible && rand.Float64() < math.Exp(-deltaE/temp) {
			incumbent = newSolution
			//p := math.Exp(-deltaE / temp)
			//fmt.Printf("Temp: %v, DeltaE: %v, Probability: %v\n", temp, deltaE, p)
		}
		temp = temp * a
		y[i] = temp
	}

	return best, bestObj, x, y, pp
}

func wrongstate(solution [][]*models.Call) bool {
	for row := range solution {
		for col := range solution[row] {
			if solution[row][col].PickedUp == true {
				fmt.Println(solution[row][col])
				return true
			}
		}
	}
	return false
}
