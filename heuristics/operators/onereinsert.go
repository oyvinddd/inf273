package operators

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

var totalIt int = 0

// OneReinsert operator performs a 1-reinsert on the given solution
func OneReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	totalIt++
	newSolution := util.CopySolution(solution)
	random := rand.Intn(len(newSolution))
	if removedCall := removeCall(&newSolution[random]); removedCall != nil {
		index := randomCompatibleIndex(data, removedCall)
		insertCall(data, data.Vehicles[random], &newSolution[index], removedCall)
	}
	return newSolution
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
		// don't bother finding the optimal delivery if vehicle is dummy
		if !vehicle.IsDummy() {
			optIndex := indexOfOptimalDelivery(data, vehicle, *vehicleCalls, index)
			rightShiftAndInsert(*vehicleCalls, optIndex)
		}
	}
}

func randomCompatibleIndex(data models.INF273Data, call *models.Call) int {
	index := rand.Intn(data.NoOfVehicles)
	for !data.VehicleAndCallIsCompatible(data.Vehicles[index].Index, call.Index) {
		index--
		if index < 0 {
			index = data.NoOfVehicles - 1
		}
	}
	return index
}

func weights(count int) []float32 {
	w := make([]float32, count)
	for i := 0; i < count; i++ {
		w[i] = 1.0
	}
	w[len(w)-1] = 0.2
	return w
}
