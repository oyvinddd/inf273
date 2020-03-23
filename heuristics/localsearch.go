package heuristics

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/models"
)

// LocalSearch does a local search on a given solution
func LocalSearch(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int) {
	current, best := solution, solution
	obj := a2.CalcTotalObjective(data, best)
	var p1 float32 = 0.2
	var p2 float32 = 0.2
	for i := 0; i < maxIterations; i++ {
		random := rand.Float32()
		if random < p1 {
			current = operators.TwoExchange(data, best)
		} else if random < p1+p2 {
			current = operators.ThreeExchange(data, best)
		} else {
			current = operators.OneReinsert(data, best)
		}
		if current != nil && isFeasible(data, current) {
			if currentObj := a2.CalcTotalObjective(data, current); currentObj < obj {
				best = current
				obj = currentObj
			}
		}
	}
	return best, obj
}
