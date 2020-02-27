package solution

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/models"
)

// ---------------- ASSIGNMENT #2 ----------------

// GenerateSolution generates a random solution
func GenerateSolution(data models.INF273Data) [][]*models.Call {
	solution := make([][]*models.Call, data.NoOfVehicles)
	// fill rows with calls (every call appears two times; pickup/delivery)
	for _, call := range data.Calls {
		i := randomNumber(0, data.NoOfVehicles)
		ptr := new(models.Call)
		*ptr = call
		if i == len(solution)-1 {
			// dummy vehicles
			solution[i] = append(solution[i], ptr)
		} else {
			// regular vehicles (one for pickup and one for delivery)
			solution[i] = append(solution[i], ptr, ptr)
		}
	}
	// for each row, randomize the order of the calls
	for _, calls := range solution {
		shuffleSlice(calls)
	}
	return solution
}

// CheckFeasibility checks the feasability of a given solution
func CheckFeasibility(data models.INF273Data, solution [][]*models.Call) error {
	var err error = nil
	for row := range solution {
		vehicle, load := data.Vehicles[row], 0
		// skip feasibility checks for all dummy vehicles
		if vehicle.IsDummy() {
			continue
		}
		for _, call := range solution[row] {
			// vehicle capacity
			if load > vehicle.Capacity {
				err = errors.New("Infeasible solution: vehicle capacity")
			}
			// calls and vehicle compatibility
			if !data.VehicleAndCallIsCompatible(vehicle.Index, call.Index) {
				err = errors.New("Infeasible solution: compatibility")
			}
			if call.PickedUp {
				call.PickedUp = false
				load -= call.Size
			} else {
				call.PickedUp = true
				load += call.Size
			}
		}
	}
	return err
}

// CalculateObjective takes a solution as input and returns an objective value
func CalculateObjective(data models.INF273Data, solution [][]*models.Call) int {

	var obj int = 0
	for row := range solution {
		vehicle := data.Vehicles[row]
		for col, call := range solution[row] {

			// handle cost of not transporting
			if vehicle.IsDummy() && !call.PickedUp {
				obj += call.Penalty
				call.PickedUp = true
				continue
			}
			if col == 0 {
				// handle the cost of reaching the first customer from the home node
				obj += data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index).Cost
			} else if col > 0 {

				ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
				from, to := 0, 0
				previousCall := solution[row][col-1]

				if previousCall.Delivered {
					from = previousCall.Destination
				} else {
					previousCall.PickedUp = true
					from = previousCall.Origin
				}

				if call.PickedUp || call == previousCall {
					call.PickedUp = true
					call.Delivered = true
					to = call.Destination
					obj += ntac.OriginCost + ntac.DestinationCost
					fmt.Printf("CALL: %d\n", call.Index)
				} else {
					call.PickedUp = true
					to = call.Origin
				}

				obj += data.GetTravelTimeAndCost(from, to, vehicle.Index).Cost
			}
		}
	}
	return obj
}

// ---------------- HELPER FUNCTIONS ----------------

func randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min)
}

func shuffleSlice(a []*models.Call) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
}
