package operators

import (
	"fmt"
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// InversionRemoval removes inversions in a solution by exchanging order of calls
func InversionRemoval(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)

	for i := 0; i < 10; i++ {
		index := rand.Intn(data.NoOfVehicles - 1)
		vehicle := data.Vehicles[index]

		if !vehicle.IsDummy() && len(newSolution[index]) > 2 {
			if i1, i2, found := invertedIndices(data, newSolution[index]); found {
				fmt.Printf("before: %v %v\n", newSolution[index][i1].Index, newSolution[index][i2].Index)
				newSolution[index][i1], newSolution[index][i2] = newSolution[index][i2], newSolution[index][i1]
				fmt.Printf("after: %v %v\n", newSolution[index][i1].Index, newSolution[index][i2].Index)
			}
		}
	}
	return newSolution
}

// --------- PRIVATE HELPER FUNCTIONS ---------

func invertedIndices(data models.INF273Data, calls []*models.Call) (int, int, bool) {
	if noOfCalls := len(calls); noOfCalls > 2 {
		pickups := findPickups(calls)
		start := rand.Intn(len(calls))
		c1 := calls[start]
		for i := start + 1; i < len(calls)-1; i++ {
			c2 := calls[i]
			ltwC1 := lowerTimeWindow(c1, start, pickups)
			ltwC2 := lowerTimeWindow(c2, i, pickups)
			if ltwC2 < ltwC1 {
				return start, i, true
			}
		}
	}
	return -1, -1, false
}

func findPickups(calls []*models.Call) map[int]int {
	pickups := make(map[int]int)
	for index, call := range calls {
		if _, ok := pickups[call.Index]; !ok {
			pickups[call.Index] = index
		}
	}
	return pickups
}

func lowerTimeWindow(call *models.Call, index int, pickups map[int]int) int {
	if pickups[call.Index] == index { // call is pickup
		return call.LowerPW
	}
	// call is destination
	return call.LowerDW
}
