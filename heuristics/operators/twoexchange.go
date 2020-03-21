package operators

import (
	"fmt"
	"math/rand"

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

	fmt.Printf("Swapping (%v, %v) and (%v, %v)\n", r1, r3, r2, r4)

	// 3. swap calls
	*newSolution[r1][r3], *newSolution[r2][r4] = *newSolution[r2][r4], *newSolution[r1][r3]

	// 4. align delivery alongside pickup
	alignPickupAndDelivery(&newSolution[r1], newSolution[r1][r3])
	alignPickupAndDelivery(&newSolution[r2], newSolution[r2][r4])
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

func alignPickupAndDelivery(calls *[]*models.Call, call *models.Call) {
	ca, cb := 0, len(*calls)-1
	ia, ib := 0, 0
	for ca < cb {
		// pickup call
		if (*calls)[ca] == call {
			ia = ca
		}
		// delivery call
		if (*calls)[cb] == call {
			ib = cb
		}
		ca++
		cb--
	}
	(*calls)[ia+1], (*calls)[ib] = (*calls)[ib], (*calls)[ia+1]
}
