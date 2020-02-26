package main

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
		// add dummy vehicle to the last index in the array
		if i == len(solution)-1 {
			solution[i] = append(solution[i], ptr)
		} else {
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
		// skip feasibility checks if vehicle is dummy
		if vehicle.IsDummy() {
			continue
		}
		for _, call := range solution[row] {
			if !data.VehicleAndCallIsCompatible(vehicle.Index, call.Index) {
				err = errors.New("Infeasible solution: compatibility")
			} else if load > vehicle.Capacity {
				err = errors.New("Infeasible solution: vehicle capacity")
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
			// handle the cost of reaching the first customer from the home node
			if col == 0 {
				ttac := data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index)
				obj += ttac.Cost
			}
			// handle travel costs and node costs
			if col > 0 {

				previousCall := solution[row][col-1]

				from := previousCall.Location()
				if !previousCall.PickedUp {
					call.PickedUp = true
				}
				to := call.Location()
				if !call.PickedUp {
					previousCall.PickedUp = true
				}

				ttac := data.GetTravelTimeAndCost(from, to, vehicle.Index)
				obj += ttac.Cost

				//ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)

				fmt.Printf("\nGoing from %d (%d) to %d (%d)\n", from, previousCall.Index, to, call.Index)

				// if !previousCall.PickedUp {
				// 	previousCall.PickedUp = true
				// }
				// if !call.PickedUp {
				// 	call.PickedUp = true
				// }

				// cost of transportation

				// from, to := 0, 0
				// ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
				// call2 := solution[row][col+1]
				// if !call.PickedUp {
				// 	call.PickedUp = true
				// 	from = call.Origin
				// 	obj += ntac.OriginCost
				// } else {
				// 	from = call.Destination
				// 	obj += ntac.DestinationCost
				// }
				// if !call2.PickedUp {
				// 	to = call2.Origin
				// } else {
				// 	to = call2.Destination
				// }
				// ttac := data.GetTravelTimeAndCost(from, to, vehicle.Index)
				// obj += ttac.Cost
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

func printSolution(solution [][]*models.Call) {
	for i := range solution {
		row := solution[i]
		if len(row) == 0 {
			fmt.Println("[-]")
			continue
		}
		for _, e := range solution[i] {
			fmt.Printf("[%v]", e.Index)
		}
		fmt.Println()
	}
}
