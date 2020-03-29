package operators

import (
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// WeightedReinsert tries to even out the average calls in each route
func WeightedReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	w1, w2 := generateWeights(newSolution, data.NoOfVehicles, data.NoOfCalls)
	removeIndex := util.WeightedRandomNumber(w1)
	if removedCall := removeCall(&newSolution[removeIndex]); removedCall != nil {
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
	return removeWeights, insertWeights
}
