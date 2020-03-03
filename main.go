package main

import (
	"fmt"
	"log"

	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {

	// time program exection
	defer util.NewTimer().PrintElapsed()

	// parse input file
	data, err := util.ParseFile("data/Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	// add dummy vehicle to list of parsed vehicles
	data.Vehicles = append(data.Vehicles, *models.NewDummyVehicle())
	data.NoOfVehicles++

	// generate a random solution
	solution := GenerateSolution(data)
	printSolution(solution)

	// check feasibility of solution
	if err := CheckFeasibility(data, solution); err != nil {
		fmt.Println(err) // print error but don't terminate execution
	}

	// calculate objective
	obj := CalculateObjective(data, solution)
	fmt.Printf("\nObjective function: %v\n", obj)
}

// --------------- HELPER FUNCTIONS ---------------

func printSolution(solution [][]*models.Call) {
	fmt.Println("-------- SOLUTION REPRESENTATION --------")
	for i := range solution {
		row := solution[i]
		if len(row) == 0 {
			fmt.Println("[-]")
			continue
		}
		for _, e := range solution[i] {
			fmt.Printf("[%v]", e.Index)
		}
		if i == len(solution)-1 {
			fmt.Print(" <-- Dummy vehicle (unhandled calls)")
		}
		fmt.Println()
	}
	fmt.Println()
}

// Used if we want to make a separate copy of the solution
func deepCopy(solution [][]*models.Call) [][]*models.Call {
	cp := make([][]*models.Call, len(solution))
	for i := range solution {
		cp[i] = make([]*models.Call, len(solution[i]))
		for j := range solution[i] {
			ptr := new(models.Call)
			*ptr = *solution[i][j]
			cp[i][j] = ptr
		}
	}
	return cp
}
