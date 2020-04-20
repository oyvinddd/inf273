package operators

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OptReinsert operator performs an optimal 1-reinsert on the given solution
func OptReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	if routeIndex := indexOfWorstRoute(data, solution); routeIndex > -1 {
		if removedCall := removeRandomCall(&newSolution[routeIndex]); removedCall != nil {
			index := randomCompatibleIndex(data, removedCall)
			insertCall(data, data.Vehicles[routeIndex], &newSolution[index], removedCall)
		}
	}
	return newSolution
}

// ------- PRIVATE HELPER FUNCTIONS -------

func indexOfWorstRoute(data models.INF273Data, solution [][]*models.Call) int {
	var index int = -1
	var maxObj int = 0
	for i, calls := range solution {
		vehicle := data.Vehicles[i]
		if vehicle.IsDummy() {
			continue
		}
		obj := a2.VehicleObjective(data, vehicle, calls)
		if obj > maxObj {
			maxObj = obj
			index = i
		}
	}
	return index
}
