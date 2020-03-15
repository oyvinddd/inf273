package heuristics

import (
	"github.com/oyvinddd/inf273/models"
)

// LocalSearch does a local search on a given solution
func LocalSearch(data models.INF273Data, solution [][]*models.Call) {
	// best, current := solution, solution
	// p1, p2, p3 := 1/3, 1/3, 1/3
	// for i := 0; i < maxIterations; i++ {
	// 	rand := rand.Intn(10)
	// 	if rand < p1 {
	// 		TwoExchange(best)
	// 	} else if rand < p1+p2 {
	// 		// 3-exchange
	// 	} else {
	// 		// 1 reinsert
	// 	}
	// 	currObj := CalculateObjective(data, current)
	// 	bestObj := CalculateObjective(data, best)
	// 	if isFeasible(current) && currObj < bestObj {
	// 		best = current
	// 	}
	// }
}
