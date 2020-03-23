package operators

import (
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// ThreeExchange operator performs a 3-exchange on the given solution
func ThreeExchange(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	// 1. copy existing solution
	newSolution := util.CopySolution(solution)

	// 2. generate random indices
	r1, r2, r3 := util.ThreeRandomIndices(len(newSolution))
	if len(newSolution[r1]) == 0 || len(newSolution[r2]) == 0 || len(newSolution[r3]) == 0 {
		return nil
	}
	r4 := rand.Intn(len(newSolution[r1]))
	r5 := rand.Intn(len(newSolution[r2]))
	r6 := rand.Intn(len(newSolution[r3]))

	// 3. swap calls
	*newSolution[r1][r4], *newSolution[r2][r5] = *newSolution[r2][r5], *newSolution[r1][r4]
	*newSolution[r3][r6], *newSolution[r1][r4] = *newSolution[r1][r4], *newSolution[r3][r6]

	// 4. align delivery alongside pickup
	p1, _ := alignPickupAndDelivery(newSolution[r1], newSolution[r1][r4])
	p2, _ := alignPickupAndDelivery(newSolution[r2], newSolution[r2][r5])
	p3, _ := alignPickupAndDelivery(newSolution[r3], newSolution[r3][r6])

	// 5. find optimal position of delivery and insert it
	v1 := data.Vehicles[r1]
	v2 := data.Vehicles[r2]
	v3 := data.Vehicles[r3]
	i1 := indexOfOptimalDelivery(data, v1, newSolution[r1], p1)
	i2 := indexOfOptimalDelivery(data, v2, newSolution[r2], p2)
	i3 := indexOfOptimalDelivery(data, v3, newSolution[r3], p3)
	rightShiftAndInsert(newSolution[r1], i1)
	rightShiftAndInsert(newSolution[r2], i2)
	rightShiftAndInsert(newSolution[r3], i3)

	return newSolution
}
