package main

import (
	"fmt"

	"github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {

	// time program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file and create initial outsourced solution
	data, _ := util.ParseFile("data/Call_7_Vehicle_3.txt", true)
	solution := assignment2.CreateOutsourcedSolution(data)

	for i := 0; i < 10000; i++ {
		heuristics.OneReinsert(solution)
	}

	util.PrintSolution(solution)

	obj := assignment2.CalcVehicleObjective(data, data.Vehicles[1], solution[1])
	fmt.Println(obj)
}
