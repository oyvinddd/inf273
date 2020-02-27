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

// TODO: remove
var fromHome = 0
var notTrans = 0
var travelCost = 0
var nodeCost = 0

// CalculateObjective takes a solution as input and returns an objective value
func CalculateObjective(data models.INF273Data, solution [][]*models.Call) int {

	var obj int = 0
	for row := range solution {
		vehicle := data.Vehicles[row]
		for col, call := range solution[row] {

			// handle cost of not transporting
			if vehicle.IsDummy() && !call.PickedUp {
				obj += call.Penalty
				notTrans += call.Penalty
				call.PickedUp = true
				continue
			}
			if col == 0 {
				// handle the cost of reaching the first customer from the home node
				ttac := data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index)
				obj += ttac.Cost
				fromHome = ttac.Cost
				//fmt.Printf("REACHING FIRST NODE FROM HOME: %d\n", ttac.Cost)
			} else if col > 0 {

				// handle travel and node costs

				ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
				from, to, prevNodeCost, currNodeCost := 0, 0, 0, 0
				previousCall := solution[row][col-1]

				previousCall.PickedUp = true
				if previousCall.Delivered {
					from = previousCall.Destination
					prevNodeCost = ntac.DestinationCost
				} else {
					from = previousCall.Origin
					prevNodeCost = ntac.OriginCost
				}

				if call == previousCall {
					call.PickedUp = true
					call.Delivered = true
					to = call.Destination
					currNodeCost = ntac.DestinationCost
				} else if !call.PickedUp {
					call.PickedUp = true
					to = call.Origin
					currNodeCost = ntac.OriginCost
				} else {
					call.Delivered = true
					to = call.Destination
					currNodeCost = ntac.DestinationCost
				}

				ttac := data.GetTravelTimeAndCost(from, to, vehicle.Index)
				obj += (prevNodeCost + currNodeCost + ttac.Cost)

				fmt.Printf("\nVehicle #%d: Going from %d (%d) to %d (%d) costs %v\n", vehicle.Index, from, previousCall.Index, to, call.Index, ttac.Cost)
			}
		}
	}
	fmt.Printf("## COST OF REACHING FIRST NODE FROM HOME: %v ##\n", fromHome)
	fmt.Printf("## COST OF NOT TRANSPORTING: %v ##\n", notTrans)
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
