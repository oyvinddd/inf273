package operators

import (
	"math/rand"

	"github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
)

// TwoExchange operator performs a 2-exchange on the given solution
func TwoExchange(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	// 1. copy existing solution
	newSolution := copySolution(solution)

	// 2. swap two random calls in different vehicles
	r1, r2 := randomIndices(len(newSolution))

	if len(newSolution[r1]) == 0 || len(newSolution[r2]) == 0 {
		return nil
	}

	r3 := rand.Intn(len(newSolution[r1]))
	r4 := rand.Intn(len(newSolution[r2]))

	// 3. swap calls
	*newSolution[r1][r3], *newSolution[r2][r4] = *newSolution[r2][r4], *newSolution[r1][r3]

	// 4. align delivery alongside pickup
	alignPickupAndDelivery(newSolution[r1], newSolution[r1][r3])
	alignPickupAndDelivery(newSolution[r2], newSolution[r2][r4])

	// 5. find optimal placement of delivery
	return newSolution
}

func randomIndices(max int) (int, int) {
	r1 := rand.Intn(max - 1)
	r2 := rand.Intn(max - 1)
	if r2 >= r1 {
		r2++
	}
	return r1, r2
}

func alignPickupAndDelivery(calls []*models.Call, call *models.Call) (int, int) {
	ca, cb := 0, len(calls)-1
	ia, ib := -1, -1
	for ca < cb {
		// pickup call
		if calls[ca] == call {
			ia = ca
		}
		// delivery call
		if calls[cb] == call {
			ib = cb
		}
		if ia < 0 {
			ca++
		}
		if ib < 0 {
			cb--
		}
		if ia >= 0 && ib >= 0 {
			break
		}
	}
	calls[ia+1], calls[ib] = calls[ib], calls[ia+1]
	return ia, ib
}

func indexOfOptimalDelivery(data models.INF273Data, vehicle models.Vehicle, calls []*models.Call, startingIndex int) int {
	var obj = assignment2.CalcVehicleObjective(data, vehicle, calls)
	var index = startingIndex + 1
	var optimalIndex = index
	if len(calls) > 2 {
		// find the optimal index
		for i := index; i < len(calls); i++ {
			calls[i-1], calls[i] = calls[i], calls[i-1]
			if newObj := assignment2.CalcVehicleObjective(data, vehicle, calls); newObj < obj {
				optimalIndex = i
				obj = newObj
			}
		}
	} else {
		return 1
	}
	return optimalIndex
}

func rightShiftAndInsert(s []*models.Call, index int) {
	length := len(s)
	if length > 0 && index > -1 && index < length {
		call := s[len(s)-1]
		copy(s[index+1:], s[index:])
		s[index] = call
	}
}
