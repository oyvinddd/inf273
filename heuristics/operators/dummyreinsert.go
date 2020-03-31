package operators

import (
	"math"
	"math/rand"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// DummyReinsert removes a call from the dummy vehicle and inserts it into the optimal vehicle route
func DummyReinsert(data models.INF273Data, solution [][]*models.Call) [][]*models.Call {
	newSolution := util.CopySolution(solution)
	// if # of unhandled calls is larger than the avg. call per vehicle, then proceed
	if len(newSolution[data.NoOfVehicles-1]) > data.NoOfCalls*2/data.NoOfVehicles {
		if removedCall := removeRandomCall(&newSolution[data.NoOfVehicles-1]); removedCall != nil {
			index := indexOfOptimalVehicle(data, removedCall)
			vehicle := data.Vehicles[index]
			insertCall(data, vehicle, &newSolution[index], removedCall)
		}
	}
	return newSolution
}

// ----------- PRIVATE HELPER FUNCTIONS -----------

func indexOfOptimalVehicle(data models.INF273Data, call *models.Call) int {
	optIndex, cost := rand.Intn(data.NoOfVehicles-1), math.MaxInt32
	for i := 0; i < data.NoOfVehicles-1; i++ {
		vehicle := data.Vehicles[i]
		ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
		currCost := ntac.OriginCost + ntac.DestinationCost
		if data.VehicleAndCallIsCompatible(vehicle.Index, call.Index) && currCost < cost {
			optIndex = i
			cost = currCost
		}
	}
	return optIndex
}
