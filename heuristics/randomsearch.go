package heuristics

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
)

const maxIterations int = 10000

// RandomSearch randomly searches for a better solution
func RandomSearch(data models.INF273Data, solution [][]*models.Call) ([][]*models.Call, int) {
	bestObjective := a2.CalculateObjective(data, solution)
	bestSolution := solution
	for i := 0; i < maxIterations; i++ {

		currentSolution := a2.GenerateSolution(data)
		currentObjective := a2.CalculateObjective(data, currentSolution)

		if isFeasible(data, solution) && currentObjective < bestObjective {
			bestObjective = currentObjective
			bestSolution = currentSolution
		}
	}
	return bestSolution, bestObjective
}

func isFeasible(data models.INF273Data, solution [][]*models.Call) bool {
	if err := a2.CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}
