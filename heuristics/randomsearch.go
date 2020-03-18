package heuristics

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

const maxIterations int = 10000

var noOfSolutions int = 0

// RandomSearch randomly searches for a better solution
func RandomSearch(data models.INF273Data, solution [][]*models.Call) (int, int) {
	obj := a2.CalculateObjective(data, solution)
	for i := 0; i < maxIterations; i++ {
		if newSolution := a2.GenerateSolution(data); isFeasible(data, newSolution) {
			util.PrintSolution(newSolution)
			obj = a2.CalculateObjective(data, newSolution)
			noOfSolutions++
		}
	}
	return obj, noOfSolutions
}

func isFeasible(data models.INF273Data, solution [][]*models.Call) bool {
	if err := a2.CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}
