package operators

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OneReinsert operator performs a 1-reinsert on the given solution
func OneReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	random := rand.Intn(len(newSolution))
	if removedCall := removeRandomCall(&newSolution[random]); removedCall != nil {
		index := randomCompatibleIndex(data, removedCall)
		insertCall(data, data.Vehicles[random], &newSolution[index], removedCall)
	}
	return newSolution
}

// ------- PRIVATE HELPER FUNCTIONS -------

func indexOfMostExpensiveRoute(data models.INF273Data, solution [][]*models.Call) int {
	var idx int = -1
	var obj int = 0
	for i := 0; i < data.NoOfVehicles-1; i++ {
		vehicle := data.Vehicles[i]
		currObj := a2.VehicleObjective(data, vehicle, solution[i])
		if currObj > obj {
			obj = currObj
			idx = i
		}
	}
	return idx
}

func removeRandomCall(vehicleCalls *[]*models.Call) *models.Call {
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

// WIP
func removeSuboptimalCall(data models.INF273Data, vehicle models.Vehicle, calls *[]*models.Call) *models.Call {
	var callToRemove *models.Call = nil
	noOfCalls := len(*calls)
	if noOfCalls > 0 {
		var excluded []*models.Call
		// removing incompiatible calls is no 1 priority
		//callToRemove = incompatibleCall(data, vehicle, calls)
		callToRemove = mostExpensiveCall(data, vehicle, *calls)
		if callToRemove == nil {
			randomIndex := rand.Intn(noOfCalls)
			callToRemove = (*calls)[randomIndex]
		}
		for _, c := range *calls {
			if c != callToRemove {
				excluded = append(excluded, c)
			}
		}
		*calls = excluded
	}
	return callToRemove
}

// WIP (untested)
func mostExpensiveCall(data models.INF273Data, vehicle models.Vehicle, calls []*models.Call) *models.Call {
	if !vehicle.IsDummy() && len(calls) != 0 {

		pickup := make(map[int]int)
		cost := make(map[int]int)

		for _, call := range calls {
			if index, ok := pickup[call.Index]; !ok {
				pickup[call.Index] = index
			}
		}

		c := data.GetTravelTimeAndCost(vehicle.Home, calls[0].Origin, vehicle.Index).Cost
		cost[calls[0].Index] += c
		for i := 0; i < len(calls)-1; i++ {

			c1 := calls[i]
			c2 := calls[i+1]

			if c1 == c2 {
				c := data.GetTravelTimeAndCost(c1.Origin, c2.Destination, vehicle.Index).Cost
				cost[c1.Index] += c
			} else {
				start, end := 0, 0
				if pickup[c1.Index] == i {
					start = c1.Origin
				} else {
					start = c1.Destination
				}
				if pickup[c2.Index] == i+1 {
					end = c2.Origin
				} else {
					end = c2.Destination
				}
				c := data.GetTravelTimeAndCost(start, end, vehicle.Index).Cost
				cost[c1.Index] += c
				cost[c2.Index] += c
			}
		}

		// find the most expensive call by looking at the costs
		idx, cst := 0, 0
		for k, v := range cost {
			if v > cst {
				cst = v
				idx = k
			}
		}
		return data.GetCall(idx)
	}
	return nil
}

func incompatibleCall(data models.INF273Data, vehicle models.Vehicle, calls *[]*models.Call) *models.Call {
	for i := 0; i < len(*calls); i++ {
		if !data.VehicleAndCallIsCompatible(vehicle.Index, (*calls)[i].Index) {
			return (*calls)[i]
		}
	}
	return nil
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
	index := util.WeightedRandomNumber(weights(data.NoOfVehicles))
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
