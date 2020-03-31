package operators

import (
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OptExchange takes two non-optimal calls in different routes, swaps them and tries to insert them into optimal positions
func OptExchange(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)

	i1, i2 := twoRandomVehicleIndices(solution)
	if i1 > -1 && i2 > -1 {
		vehicle1 := data.Vehicles[i1]
		vehicle2 := data.Vehicles[i2]

		i3 := indexOfSuboptimalCall(data, vehicle1, vehicle2, newSolution[i1])
		i4 := indexOfSuboptimalCall(data, vehicle2, vehicle1, newSolution[i2])
		if i3 != -1 && i4 != -1 {

			removedCall1 := removeCall(&newSolution[i1], newSolution[i1][i3])
			removedCall2 := removeCall(&newSolution[i2], newSolution[i2][i4])

			insertCall(data, vehicle1, &newSolution[i1], removedCall2)
			insertCall(data, vehicle2, &newSolution[i2], removedCall1)

			// insertCallAtBestTime(data, vehicle1, &newSolution[i1], removedCall2)
			// insertCallAtBestTime(data, vehicle2, &newSolution[i2], removedCall1)
		}
	}
	return newSolution
}

// --------- PRIVATE HELPER FUNCTIONS ---------

func indexOfSuboptimalCall(data models.INF273Data, oldVehicle models.Vehicle, newVehicle models.Vehicle, calls []*models.Call) int {
	var suboptimalIndex int = -1
	for index, call := range calls {
		// first criteria for optimality is that the new vehicle can transport the call
		if data.VehicleAndCallIsCompatible(newVehicle.Index, call.Index) {
			suboptimalIndex = index
			// second criteria is that the call gets cheaper when handled by the new vehicle rather than the old
			// we can break out of the loop immediately if this condition passes because we found our best option
			if costForVehicleAndCall(data, newVehicle, call) < costForVehicleAndCall(data, oldVehicle, call) {
				suboptimalIndex = index
				break
			}
		}
	}
	return suboptimalIndex
}

func indexOfOptimalTimeWindow(data models.INF273Data, calls []*models.Call, call *models.Call) int {
	if noOfCalls := len(calls); noOfCalls > 0 {
		pickups := findPickups(calls)
		for index, call := range calls {
			window := lowerTimeWindow(call, index, pickups)
			if window > call.LowerPW {
				return index
			}
		}
		return noOfCalls - 1
	}
	return 0
}

func costForVehicleAndCall(data models.INF273Data, vehicle models.Vehicle, call *models.Call) int {
	if vehicle.IsDummy() {
		return call.Penalty
	}
	originCost := data.GetNodeTimeAndCost(vehicle.Index, call.Index).OriginCost
	destinationCost := data.GetNodeTimeAndCost(vehicle.Index, call.Index).DestinationCost
	return originCost + destinationCost
}

func removeCall(calls *[]*models.Call, call *models.Call) *models.Call {
	var excluded []*models.Call
	for _, c := range *calls {
		if c != call {
			excluded = append(excluded, c)
		}
	}
	*calls = excluded
	return call
}
