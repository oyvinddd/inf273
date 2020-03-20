package assignment2

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/models"
)

func init() {
	// we only want to run this once for every program execution
	rand.Seed(time.Now().UnixNano())
}

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

// CheckFeasibility checks the feasibility of a given solution
func CheckFeasibility(data models.INF273Data, solution [][]*models.Call) error {
	for row := range solution {
		vehicle, vehicleLoad, currentTime := data.Vehicles[row], 0, 0
		// skip feasibility checks for dummy vehicle
		if vehicle.IsDummy() {
			continue
		}
		for col, call := range solution[row] {

			if col == 0 {
				// make sure to add travel time from home node to first call
				currentTime += data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index).Time
			}

			ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
			// add loading (or unloading) time to current time
			time, err := calculateLoadingOrUnloadingTime(currentTime, vehicle.Index, ntac, call)
			if err != nil {
				return err
			}
			currentTime += time

			// vehicle capacity
			if !call.PickedUp {
				vehicleLoad += call.Size
				if vehicleLoad > vehicle.Capacity {
					return fmt.Errorf("Infeasible: vehicle capacity is %d, load is %d", vehicle.Capacity, vehicleLoad)
				}
			} else {
				vehicleLoad -= call.Size
			}

			// calls and vehicle compatibility
			if !data.VehicleAndCallIsCompatible(vehicle.Index, call.Index) {
				return fmt.Errorf("Infeasible: vehicle %d not compatible with call %d", vehicle.Index, call.Index)
			}

			if col < len(solution[row])-1 {

				nextCall := solution[row][col+1]
				from, to := 0, 0

				if !call.PickedUp {
					call.PickedUp = true
					from = call.Origin
				} else {
					from = call.Destination
				}

				if !nextCall.PickedUp {
					to = nextCall.Origin
				} else {
					to = nextCall.Destination
				}

				// add travel time from current node to next node
				currentTime += data.GetTravelTimeAndCost(from, to, vehicle.Index).Time
			}
		}
	}
	resetPickedUpState(solution)
	return nil
}

// CalcTotalObjective takes a solution as input and returns an objective value
func CalcTotalObjective(data models.INF273Data, solution [][]*models.Call) int {
	var obj int = 0
	for row := range solution {
		obj += CalcVehicleObjective(data, data.Vehicles[row], solution[row])
	}
	return obj
}

// CalcVehicleObjective calculates objective function for a specific vehicle
func CalcVehicleObjective(data models.INF273Data, vehicle models.Vehicle, calls []*models.Call) int {
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

// CreateOutsourcedSolution creates a solution where all calls are handled by the dummy vehicle
func CreateOutsourcedSolution(data models.INF273Data) [][]*models.Call {
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

func calculateLoadingOrUnloadingTime(currentTime int, vehicleIndex int, ntac models.NodeTimeAndCost, call *models.Call) (int, error) {

	var time int = 0

	if !call.PickedUp {
		if currentTime < call.LowerPW {
			// vehicle arrived early at pickup, add waiting time
			time += call.LowerPW - currentTime
		} else if currentTime > call.UpperPW {
			// vehicle arrived too late at pickup, infeasible
			return 0, fmt.Errorf("Infeasible: vehicle %d arrived at pickup node %d too late", vehicleIndex, call.Origin)
		}
		time += ntac.OriginTime
	} else {
		if currentTime < call.LowerDW {
			// vehicle arrived early at destination, add waiting time
			time += call.LowerDW - currentTime
		} else if currentTime > call.UpperDW {
			// vehicle arrived too late at destination, infeasible
			return 0, fmt.Errorf("Infeasible: vehicle %d arrived at destination node %d too late", vehicleIndex, call.Destination)
		}
		time += ntac.DestinationTime
	}
	return time, nil
}

func resetPickedUpState(solution [][]*models.Call) {
	for row := range solution {
		for col := range solution[row] {
			solution[row][col].PickedUp = false
		}
	}
}

// t1 := solution[1][0]
// t2 := solution[1][1]
// *t1, *t2 = *t2, *t1
