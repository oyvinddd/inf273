package operators

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// TwoExchange operator performs a 2-exchange on the given solution
func TwoExchange(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	// 1. copy existing solution
	newSolution := util.CopySolution(solution)

	// 2. chose two random vehicle routes
	r1, r2 := util.TwoRandomIndices(len(newSolution))

	if len(newSolution[r1]) == 0 || len(newSolution[r2]) == 0 {
		return nil
	}

	r3 := rand.Intn(len(newSolution[r1]))
	r4 := rand.Intn(len(newSolution[r2]))

	// 3. swap calls
	*newSolution[r1][r3], *newSolution[r2][r4] = *newSolution[r2][r4], *newSolution[r1][r3]

	// 4. align delivery alongside pickup
	p1, _ := alignPickupAndDelivery(newSolution[r1], newSolution[r1][r3])
	p2, _ := alignPickupAndDelivery(newSolution[r2], newSolution[r2][r4])

	// 5. find optimal position of delivery
	v1 := data.Vehicles[r1]
	v2 := data.Vehicles[r2]
	i1 := indexOfOptimalDelivery(data, v1, newSolution[r1], p1)
	i2 := indexOfOptimalDelivery(data, v2, newSolution[r2], p2)
	rightShiftAndInsert(newSolution[r1], i1)
	rightShiftAndInsert(newSolution[r2], i2)

	return newSolution
}

// -------- PRIVATE HELPERS FUNCTIONS --------

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
	var obj = a2.CalcVehicleObjective(data, vehicle, calls)
	var index = startingIndex + 1
	var optimalIndex = index
	if len(calls) > 2 {
		// find the optimal index
		for i := index; i < len(calls); i++ {
			calls[i-1], calls[i] = calls[i], calls[i-1]
			if newObj := a2.CalcVehicleObjective(data, vehicle, calls); newObj < obj {
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
