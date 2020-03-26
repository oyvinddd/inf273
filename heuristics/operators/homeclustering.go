package operators

import (
	"fmt"
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// HomeClustering operator uses a clustering technique to segment
func HomeClustering(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)

	// 1. choose a random call to be reinserted
	if call := removeRandomCall(solution); call != nil {

		// 2. find nearest centroid among all home nodes
		optIndex, minDistance := 999999999, -1
		for index, vehicle := range data.Vehicles {
			dist := distance(data, vehicle, call)
			if dist < minDistance {
				minDistance = dist
				optIndex = index
			}
		}
		fmt.Println(optIndex, minDistance)

		// 3. insert call into the optimal vehicle route
		// insertCall(data, )
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
	return 999999999999 // TODO: add constant here.. very big number
}

func removeRandomCall(solution [][]*models.Call) *models.Call {
	var call *models.Call = nil
	var excluded []*models.Call
	if noOfCalls := len(solution); noOfCalls > 0 {
		r1 := rand.Intn(noOfCalls)
		calls := solution[r1]
		if len(calls) > 0 {
			r2 := rand.Intn(len(calls))
			call = calls[r2]
			for _, c := range calls {
				if c != call {
					excluded = append(excluded, c)
				}
			}
			solution[r1] = excluded
		}
	}
	return call
}

// https://www.youtube.com/watch?v=_aWzGGNrcic&t=3s
