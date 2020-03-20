package heuristics

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
)

func TwoExchange(solution [][]*models.Call) [][]*models.Call {
	return nil
}

func ThreeExchange(solution [][]*models.Call) [][]*models.Call {
	return nil
}

func OneReinsert(solution [][]*models.Call) [][]*models.Call {
	random := rand.Intn(len(solution))
	if removedCall := removeCall(&solution[random]); removedCall != nil {
		random := rand.Intn(len(solution))
		insertCall(&solution[random], removedCall)
	}
	return nil
}

// ------- PRIVATE HELPERS -------

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

func insertCall(vehicleCalls *[]*models.Call, call *models.Call) {
	noOfCalls := len(*vehicleCalls)
	if noOfCalls == 0 {
		*vehicleCalls = append(*vehicleCalls, call, call)
	} else {
		index := rand.Intn(noOfCalls - 1)
		*vehicleCalls = append(*vehicleCalls, nil, nil)
		copy((*vehicleCalls)[index+2:], (*vehicleCalls)[index:])
		(*vehicleCalls)[index] = call
		(*vehicleCalls)[index+1] = call
	}
}
