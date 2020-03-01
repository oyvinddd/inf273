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
			fmt.Print(" <-- Dummy vehicles")
		}
		fmt.Println()
	}
	fmt.Println()
}

func sol() [][]*models.Call {

	c1 := models.NewCall(1, 17, 37, 4601, 790000, 345, 417, 345, 1006)
	c2 := models.NewCall(2, 33, 36, 13444, 430790, 96, 168, 96, 529)
	c3 := models.NewCall(3, 17, 27, 6596, 200354, 715, 787, 715, 1089)
	c4 := models.NewCall(4, 6, 1, 11052, 275455, 0, 72, 0, 435)
	c5 := models.NewCall(5, 38, 33, 6643, 642740, 107, 179, 107, 593)
	c6 := models.NewCall(6, 10, 38, 14139, 587198, 300, 372, 300, 801)
	c7 := models.NewCall(7, 26, 6, 5310, 359885, 178, 250, 178, 567)

	return [][]*models.Call{
		{c3, c3},
		{c7, c1, c7, c1},
		{c5, c5},
		{c2, c4, c6}, // not transported
	}
}
