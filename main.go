package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {
	rand.Seed(time.Now().UnixNano())

	// time program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file and create initial outsourced solution
	data, err := util.ParseFile("Call_7_Vehicle_3.txt", true)
	if err != nil {
		log.Fatal(err)
	}
	solution := util.FeasibleSolution()
	util.PrintSolution(solution)

	for i := 0; i < 10000; i++ {
		operators.TwoExchange(data, solution)
	}
	// util.PrintSolution(newSol)
}
