package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

func main() {

	// time the program execution
	defer util.NewTimer().PrintElapsed()

	// parse input file
	data, err := util.ParseFile("Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	// add dummy vehicle
	data.Vehicles = append(data.Vehicles, *models.NewDummyVehicle())
	data.NoOfVehicles++

	// generate a random solution
	solution := generateSolution(data)
	printSolution(solution)

	// check feasibility of solution
	if err := checkFeasibility(data, solution); err != nil {
		fmt.Println(err)
		//log.Fatal(err)
	}

	obj := calculateObj(data, solution)
	fmt.Printf("Objective function: %v\n", obj)
}

// ---------------- ASSIGNMENT #2 ----------------

func generateSolution(data models.INF273Data) [][]*models.Call {
	// initialize empty solution matrix
	solution := make([][]*models.Call, data.NoOfVehicles)
	// fill rows with calls (every call appears two times)
	for _, call := range data.Calls {
		i := randomNumber(0, data.NoOfVehicles)
		ptr := new(models.Call)
		*ptr = call
		solution[i] = append(solution[i], ptr, ptr)
	}
	// for each row, randomize the order of the calls
	for _, calls := range solution {
		shuffleSlice(calls)
	}
	return solution
}

func checkFeasibility(data models.INF273Data, solution [][]*models.Call) error {
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

func calculateObj(data models.INF273Data, solution [][]*models.Call) int {

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

				// cost of reaching the first customer from the home node
				tac := data.GetTravelTimeAndCost(vehicle.Home, call.Origin, vehicle.Index)
				obj += tac.Cost
			}
			if col < len(solution[row])-1 {

				// cost of transportation
				//ntac := 0
				from, to := 0, 0
				call2 := solution[row][col+1]
				if !call.PickedUp {
					call.PickedUp = true
					from = call.Origin
					ntac := data.GetNodeTimeAndCost(vehicle.Index, call.Index)
					fmt.Printf("Node cost (%v %v): O: %v D: %v\n", vehicle.Index, call.Index, ntac.OriginCost, ntac.DestinationCost)
				} else {
					from = call.Destination
				}
				if !call2.PickedUp {
					to = call2.Origin
				} else {
					to = call2.Destination
				}
				ttac := data.GetTravelTimeAndCost(from, to, vehicle.Index)
				obj += ttac.Cost

				// cost at origin and destination node
				// models.NodeTimeAndCost()
				// fmt.Printf("Origin node %v cost: %v", )
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
