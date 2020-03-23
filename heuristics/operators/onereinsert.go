package operators

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OneReinsert operator performs a 1-reinsert on the given solution
func OneReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	copiedSolution := util.CopySolution(solution)
	random := rand.Intn(len(copiedSolution))
	if removedCall := removeCall(&copiedSolution[random]); removedCall != nil {
		random := rand.Intn(len(copiedSolution))
		insertCall(data, data.Vehicles[random], &copiedSolution[random], removedCall)
	}
	return copiedSolution
}

// ------- PRIVATE HELPER FUNCTIONS -------

func removeCall(vehicleCalls *[]*models.Call) *models.Call {
	var removedCall *models.Call = nil
	noOfCalls := len(*vehicleCalls)
	if noOfCalls > 0 {
		index := rand.Intn(noOfCalls)
		removedCall = (*vehicleCalls)[index]
		var excluded []*models.Call
		for _, c := range *vehicleCalls {
			if c != removedCall {
				excluded = append(excluded, c)
			}
		}
		*vehicleCalls = excluded
	}
	return removedCall
}

func insertCall(data models.INF273Data, vehicle models.Vehicle, vehicleCalls *[]*models.Call, call *models.Call) {
	noOfCalls := len(*vehicleCalls)
	if noOfCalls == 0 {
		*vehicleCalls = append(*vehicleCalls, call, call)
	} else {
		index := rand.Intn(noOfCalls - 1)
		*vehicleCalls = append(*vehicleCalls, nil, nil)
		copy((*vehicleCalls)[index+2:], (*vehicleCalls)[index:])
		(*vehicleCalls)[index] = call
		(*vehicleCalls)[index+1] = call
		// don't bother finding the optimal destination if vehicle is dummy
		if !vehicle.IsDummy() {
			optIndex := indexOfOptimalDelivery(data, vehicle, *vehicleCalls, index)
			rightShiftAndInsert(*vehicleCalls, optIndex)
		}
	}
}
