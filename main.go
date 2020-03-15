package main

import (
	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {

	// time program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file and create initial outsourced solution
	data, _ := util.ParseFile("data/Call_7_Vehicle_3.txt", true)
	outsourcedSolution := a2.CreateOutsourcedSolution(data)

	solution, _ := heuristics.RandomSearch(data, outsourcedSolution)

	util.PrintSolution(solution)
}
