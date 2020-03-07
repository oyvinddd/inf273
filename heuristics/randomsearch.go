package heuristics

import (
	"fmt"

	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
	"github.com/oyvinddd/inf273/util"
)

var data models.INF273Data

func init() {
	// load data from file
	data, _ = util.ParseFile("data/Call_7_Vehicle_3.txt", true)
}

// RandomSearch ....
func RandomSearch(solution [][]*models.Call) {
	bestObjective := CalculateObjective(&data, solution)
	for i := 0; i < 10000; i++ {
		currentSolution := GenerateSolution(data)
		currentObjective := CalculateObjective(data, currentSolution)
		if isFeasible(currentSolution) && currentObjective < bestObjective {
			bestObjective = currentObjective
			fmt.Printf("CURRENT BEST OBJ: %v\n", currentObjective)
		}
	}
}

// ----------- HELPER FUNCTIONS -----------

func isFeasible(solution [][]*models.Call) bool {
	if err := CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}

// evaluation function: Improvement (%) = 100 * (Objective of Initial solution - Best objective) / Objective of Initial solution
func f(solution [][]*models.Call) int {
	return CalculateObjective(data, solution)
}
