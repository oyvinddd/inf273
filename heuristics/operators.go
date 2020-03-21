package heuristics

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
)

// TwoExchange operator performs a 2-exchange on the given solution
func TwoExchange(solution [][]*models.Call) [][]*models.Call {
	newSolution := copySolution(solution)
	return newSolution
}

// ThreeExchange operator performs a 3-exchange on the given solution
func ThreeExchange(solution [][]*models.Call) [][]*models.Call {
	newSolution := copySolution(solution)
	return newSolution
}

// OneReinsert operator performs a 1-reinsert on the given solution
func OneReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := copySolution(solution)
	random := rand.Intn(len(newSolution))
	if removedCall := removeCall(&newSolution[random]); removedCall != nil {
		random := rand.Intn(len(newSolution))
		insertCall(data, &newSolution[random], removedCall)
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

func insertCall(data models.INF273Data, vehicleCalls *[]*models.Call, call *models.Call) {
	noOfCalls := len(*vehicleCalls)
	if noOfCalls == 0 {
		*vehicleCalls = append(*vehicleCalls, call, call)
	} else {
		index := rand.Intn(noOfCalls - 1)
		*vehicleCalls = append(*vehicleCalls, nil, nil)
		copy((*vehicleCalls)[index+2:], (*vehicleCalls)[index:])
		(*vehicleCalls)[index] = call
		(*vehicleCalls)[index+1] = call
		findBestDestination(data, vehicleCalls, index)
	}
}

func findBestDestination(data models.INF273Data, s *[]*models.Call, index int) {
	// advance call one step forward and calculate objective
	for i := index; i < len(*s)-1; i++ {
		(*s)[i], (*s)[i+1] = (*s)[i+1], (*s)[i]

	}
}

func copySolution(s [][]*models.Call) [][]*models.Call {
	duplicate := make([][]*models.Call, len(s))
	for i := range s {
		duplicate[i] = make([]*models.Call, len(s[i]))
		copy(duplicate[i], s[i])
	}
	return duplicate
}
