package operators

import (
	"math"
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// HomeClustering operator uses a clustering technique to segment
func HomeClustering(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)

	// 1. choose a random call to be reinserted
	if removedCall, _ := removeWorstCall(data, solution); removedCall != nil {

		// 2. find nearest centroid among all home nodes
		bestDist, bestIndex := math.MaxInt32, 0
		for index, vehicle := range data.Vehicles {
			if newDist := distance(data, vehicle, removedCall); newDist < bestDist {
				bestDist = newDist
				bestIndex = index
			}
		}

		// 3. insert call into the optimal vehicle route
		insertCall(data, data.Vehicles[bestIndex], &newSolution[bestIndex], removedCall)
	}
	return newSolution
}

// -------- PRIVATE HELPER FUNCTIONS --------

func distance(data models.INF273Data, vehicle models.Vehicle, call *models.Call) int {
	if data.VehicleAndCallIsCompatible(vehicle.Index, call.Index) {
		// travel time from home node to origin of given call
		t1 := data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index).Time
		// travel time from home node to destination of given call
		t2 := data.GetTravelTimeAndCost(vehicle.Home, call.Destination, vehicle.Index).Time
		return t1 + t2
	}
	return math.MaxInt32
}

func removeWorstCall(data models.INF273Data, solution [][]*models.Call) (*models.Call, int) {
	calls, index := randomNonEmptyRoute(data, solution)
	var vehicle models.Vehicle = data.Vehicles[index]
	var excluded []*models.Call
	var worstCall *models.Call = calls[0]
	// find index of the worst call
	for _, call := range calls {
		if distance(data, vehicle, call) > distance(data, vehicle, worstCall) {
			worstCall = call
		}
	}
	// remove the worst call from the slice
	for _, call := range calls {
		if call != worstCall {
			excluded = append(excluded, call)
		}
	}
	solution[index] = excluded
	return worstCall, index
}

func randomNonEmptyRoute(data models.INF273Data, solution [][]*models.Call) ([]*models.Call, int) {
	random := rand.Intn(data.NoOfVehicles)
	for len(solution[random]) == 0 {
		random++
		if random >= data.NoOfVehicles {
			random = 0
		}
	}
	return solution[random], random
}

// https://www.youtube.com/watch?v=_aWzGGNrcic&t=3s
