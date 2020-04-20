package operators

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OptOrdering ...
func OptOrdering(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	for j := 0; j < data.NoOfVehicles; j++ {
		index := rand.Intn(data.NoOfVehicles - 1)
		vehicle := data.Vehicles[index]
		for i := 0; i < len(solution[index]); i++ {
			i1, i2 := util.TwoRandomIndices(len(newSolution[index]))
			swapCalls(data, vehicle, &newSolution[index], i1, i2)
		}
	}
	return newSolution
}

// -------- PRIVATE HELPER FUNCTIONS --------

func swapCalls(data models.INF273Data, vehicle models.Vehicle, calls *[]*models.Call, i1 int, i2 int) bool {
	oldObj := a2.VehicleObjective(data, vehicle, *calls)
	(*calls)[i1], (*calls)[i2] = (*calls)[i2], (*calls)[i1]

	err := a2.CheckTimeWindows(data, vehicle, *calls)
	err = a2.CheckCapacity(data, vehicle, *calls)
	err = a2.CheckCompatibility(data, vehicle, *calls)
	obj := a2.VehicleObjective(data, vehicle, *calls)

	if err != nil || obj > oldObj {
		// swap back to original position if solution is infeasible or obj is smaller than the old one
		(*calls)[i1], (*calls)[i2] = (*calls)[i2], (*calls)[i1]
		return false
	}
	return true
}
