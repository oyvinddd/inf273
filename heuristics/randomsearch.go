package heuristics

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
)

const maxIterations int = 10000

// RandomSearch randomly searches for a better solution
func RandomSearch(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int) {
	best := solution
	obj := a2.CalcTotalObjective(data, best)
	for i := 0; i < maxIterations; i++ {
		if current := a2.GenerateSolution(data); isFeasible(data, current) {
			if currentObj := a2.CalcTotalObjective(data, current); currentObj < obj {
				best = current
				obj = currentObj
			}
		}
	}
	return best, obj
}

func isFeasible(data models.INF273Data, solution [][]*models.Call) bool {
	if err := a2.CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}
