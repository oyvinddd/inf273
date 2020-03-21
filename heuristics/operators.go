package heuristics

import (
	"fmt"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// OneReinsert operator performs a 1-reinsert on the given solution
func OneReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := copySolution(solution)
	random := rand.Intn(len(newSolution))
	if removedCall := removeCall(&newSolution[random]); removedCall != nil {
		random := rand.Intn(len(newSolution))
		insertCall(data, data.Vehicles[random], &newSolution[random], removedCall)
	}
	util.PrintSolution(solution)
	return newSolution
}

// ------- PRIVATE HELPER FUNCTIONS -------

func removeCall(vehicleCalls *[]*models.Call) *models.Call {
	var removedCall *models.Call = nil
	noOfCalls := len(*vehicleCalls)
	if noOfCalls > 0 {
		index := rand.Intn(noOfCalls)
		removedCall = (*vehicleCalls)[index]
		var b []*models.Call
		for _, c := range *vehicleCalls {
			if c != removedCall {
				b = append(b, c)
			}
		}
		*vehicleCalls = b
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
			fmt.Printf("Inserting call %v into vehicle %v\n", call.Index, vehicle.Index-1)
			findOptimalDelivery(data, vehicle, vehicleCalls, index+1)
		}
	}
}

func findOptimalDelivery(data models.INF273Data, vehicle models.Vehicle, vehicleCalls *[]*models.Call, index int) {
	noOfCalls, optIndex, obj := len(*vehicleCalls), index+1, 0
	// advance call one step forward and calculate objective
	for i := index; i < noOfCalls; i++ {
		swapCalls(vehicleCalls, index, index-1)
		if newObj := a2.CalcVehicleObjective(data, vehicle, *vehicleCalls); newObj > obj {
			optIndex = i
			obj = newObj
		}
	}
	swapCalls(vehicleCalls, noOfCalls-1, optIndex)
}

func copySolution(s [][]*models.Call) [][]*models.Call {
	duplicate := make([][]*models.Call, len(s))
	for i := range s {
		duplicate[i] = make([]*models.Call, len(s[i]))
		copy(duplicate[i], s[i])
	}
	return duplicate
}

func swapCalls(s *[]*models.Call, i1 int, i2 int) {
	(*s)[i1], (*s)[i2] = (*s)[i2], (*s)[i1]
}
