package a2

import (
	"fmt"
	"math/rand"

	"github.com/oyvinddd/inf273/models"
)

// ---------------- ASSIGNMENT #2 ----------------

// GenerateSolution generates a random solution
func GenerateSolution(data models.INF273Data) [][]*models.Call {
	solution := make([][]*models.Call, data.NoOfVehicles)
	// fill rows with pointers to calls (every call appears two times; pickup/delivery)
	for _, call := range data.Calls {
		i := rand.Intn(data.NoOfVehicles)
		ptr := new(models.Call)
		*ptr = call
		solution[i] = append(solution[i], ptr, ptr)
	}
	// for each vehicle, randomize the order of the calls
	for _, calls := range solution {
		shuffleSlice(calls)
	}
	return solution
}

var t int = 0
var a int = 0
var b int = 0
var c int = 0

// CheckFeasibility checks the feasibility of a given solution
func CheckFeasibility(data models.INF273Data, solution [][]*models.Call) error {
	t++
	for row := range solution {

		vehicle := data.Vehicles[row]
		calls := solution[row]

		if e1 := CheckTimeWindows(data, vehicle, calls); e1 != nil {
			a++
			return e1
		}

		if e2 := CheckCapacity(data, vehicle, calls); e2 != nil {
			b++
			return e2
		}

		if e3 := CheckCompatibility(data, vehicle, calls); e3 != nil {
			c++
			return e3
		}
	}
	fmt.Printf("Total times checked: %v, time window fails: %v, cap fails: %v, comp fails: %v\n", t, a, b, c)
	return nil
}

// IsFeasible returns true if a given solution is feasible, otherwise false
func IsFeasible(data models.INF273Data, solution [][]*models.Call) bool {
	if solution == nil {
		return false
	}
	if err := CheckFeasibility(data, solution); err != nil {
		return false
	}
	return true
}

// TotalObjective takes a solution as input and returns an objective value
func TotalObjective(data models.INF273Data, solution [][]*models.Call) int {
	var obj int = 0
	for row := range solution {
		obj += VehicleObjective(data, data.Vehicles[row], solution[row])
	}
	return obj
}

// VehicleObjective calculates objective function for a specific vehicle
func VehicleObjective(data models.INF273Data, vehicle models.Vehicle, calls []*models.Call) int {
	var obj = 0
	for col, call := range calls {

		// handle cost of not transporting
		if vehicle.IsDummy() {
			if !call.PickedUp {
				call.PickedUp = true
				obj += call.Penalty
			}
			continue
		}
		// handle the cost of reaching the first customer from the home node
		if col == 0 {
			obj += data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index).Cost
		}
		// handle transportation cost
		if col < len(calls)-1 {

			ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)

			nextCall := calls[col+1]
			from, to := 0, 0

			if !call.PickedUp {
				obj += ntac.OriginCost + ntac.DestinationCost
				call.PickedUp = true
				from = call.Origin
			} else {
				call.PickedUp = false
				from = call.Destination
			}

			if !nextCall.PickedUp {
				to = nextCall.Origin
			} else {
				to = nextCall.Destination
			}

			obj += data.GetTravelTimeAndCost(from, to, vehicle.Index).Cost
		}
	}
	for col := range calls {
		calls[col].PickedUp = false
	}
	return obj
}

// GenerateOutsourcedSolution creates a solution where all calls are handled by the dummy vehicle
func GenerateOutsourcedSolution(data models.INF273Data) [][]*models.Call {
	solution := make([][]*models.Call, data.NoOfVehicles)
	dummyVehicle := solution[data.NoOfVehicles-1]
	for _, call := range data.Calls {
		ptr := new(models.Call)
		*ptr = call
		dummyVehicle = append(dummyVehicle, ptr, ptr)
	}
	solution[data.NoOfVehicles-1] = dummyVehicle
	return solution
}

// ---------------- HELPER FUNCTIONS ----------------

func shuffleSlice(a []*models.Call) {
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

func resetPickedUpState(solution [][]*models.Call) {
	for row := range solution {
		for col := range solution[row] {
			solution[row][col].PickedUp = false
		}
	}
}
