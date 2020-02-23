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

	runProgram()
}

func runProgram() {

	// time the program execution
	defer util.NewTimer().PrintElapsed()

	// 1. Parse input file
	data, err := util.ParseFile("Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	// initialize empty solution matrix
	solution := make([][]*models.Call, data.NoOfVehicles)

	// 2. Generate a random solution
	generateSolution(data, solution)
	printSolution(solution)

	// 3. Check feasibility of solution
	checkFeasability(data, solution)
}

// ---------------- ASSIGNMENT #2 ----------------

func generateSolution(data models.INF273Data, matrix [][]*models.Call) {

	// fill rows with calls (every call appears two times)
	for _, call := range data.Calls {
		//fmt.Printf("Call %p: %v\n", &call, call)
		i := randomNumber(0, data.NoOfVehicles)
		ptr := new(models.Call)
		*ptr = call
		matrix[i] = append(matrix[i], ptr, ptr)
	}
	// for each row, randomize the order of the calls
	for _, calls := range matrix {
		shuffleSlice(calls)
	}
}

func checkFeasability(data models.INF273Data, solution [][]*models.Call) error {

	// vehicle capacity
	for row := range solution {
		vehicle, load := data.Vehicles[row], 0
		for _, call := range solution[row] {
			if load > vehicle.Capacity {
				return errors.New("Infeasible solution: vehicle capacity")
			}
			if call.PickedUp {
				load -= call.Size
			} else {
				load += call.Size
			}
		}
	}
	return nil
}

func calculateObj() {
	// TODO:
}

// ---------------- UTILITIES ----------------

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
			fmt.Print("[]")
		} else {
			for _, e := range solution[i] {
				fmt.Printf("[%v]", e.Index)
			}
		}
		fmt.Println()
	}
}
