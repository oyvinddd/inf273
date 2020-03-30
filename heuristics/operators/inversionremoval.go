package operators

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// InversionRemoval removes inversions in a solution by exchanging order of calls
func InversionRemoval(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)

	index := rand.Intn(data.NoOfVehicles)
	vehicle := data.Vehicles[index]

	if !vehicle.IsDummy() && len(newSolution[index]) > 2 {
		//inversionIndex = indexOfInversion(data, newSolution[index])
	}
	return newSolution
}

// --------- PRIVATE HELPER FUNCTIONS ---------

func indexOfInversion(data models.INF273Data, solution []*models.Call, call *models.Call) int {

	// pickups := make(map[int]int)
	// for i := 0; i <

	return 0
}
