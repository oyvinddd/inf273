package heuristics

import (
	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
	"github.com/oyvinddd/inf273/util"
)

var data models.INF273Data

func init() {
	// load data from file
	data, _ = util.ParseFile("data/Call_7_Vehicle_3.txt", true)
}

// RandomSearch randomly searches for a better solution
func RandomSearch(solution [][]*models.Call) ([][]*models.Call, int) {
	bestObjective := CalculateObjective(data, solution)
	bestSolution := solution
	for i := 0; i < 10000; i++ {
		currentSolution := GenerateSolution(data)
		currentObjective := CalculateObjective(data, currentSolution)
		if isFeasible(currentSolution) && currentObjective < bestObjective {
			bestObjective = currentObjective
			bestSolution = currentSolution
		}
	}
	return bestSolution, bestObjective
}

// ----------- HELPER FUNCTIONS -----------

func isFeasible(solution [][]*models.Call) bool {
	if err := CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}
