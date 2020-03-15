package heuristics

import (
	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
)

const maxIterations int = 10000

// RandomSearch randomly searches for a better solution
func RandomSearch(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int) {
	bestObjective := CalculateObjective(data, solution)
	bestSolution := solution
	for i := 0; i < maxIterations; i++ {

		currentSolution := GenerateSolution(data)
		currentObjective := CalculateObjective(data, currentSolution)

		if isFeasible(data, solution) && currentObjective < bestObjective {
			bestObjective = currentObjective
			bestSolution = currentSolution
		}
	}
	return bestSolution, bestObjective
}

func isFeasible(data models.INF273Data, solution [][]*models.Call) bool {
	if err := CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}
