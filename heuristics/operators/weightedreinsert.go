package operators

import (
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// WeightedReinsert tries to even out the average # of calls in each vehicle route
func WeightedReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	w1, w2 := generateWeights(newSolution, data.NoOfVehicles, data.NoOfCalls)
	removeIndex := util.WeightedRandomNumber(w1)
	// if # of calls in dummy vehicle is more than 1/4 of all calls, force remove from dummy
	if len(newSolution[data.NoOfVehicles-1]) > data.NoOfCalls/5 {
		removeIndex = data.NoOfVehicles - 1
	}
	if removedCall := removeRandomCall(&newSolution[removeIndex]); removedCall != nil {
		insertIndex := util.WeightedRandomNumber(w2)
		if !data.VehicleAndCallIsCompatible(data.Vehicles[insertIndex].Index, removedCall.Index) {
			insertIndex = randomCompatibleIndex(data, removedCall)
		}
		insertCall(data, data.Vehicles[insertIndex], &newSolution[insertIndex], removedCall)
	}
	return newSolution
}

// --------- PRIVATE HELPER FUNCTIONS ---------

func generateWeights(solution [][]*models.Call, noOfVehicles int, noOfCalls int) ([]float32, []float32) {
	removeWeights := make([]float32, noOfVehicles)
	insertWeights := make([]float32, noOfVehicles)
	for index, calls := range solution {
		removeWeights[index] = float32(len(calls)/2) / float32(noOfCalls)
		insertWeights[index] = 1 - float32(len(calls)/2)/float32(noOfCalls)
	}
	// setting weight for insertion to a fixed small percentage
	insertWeights[noOfVehicles-1] = 0
	return removeWeights, insertWeights
}
