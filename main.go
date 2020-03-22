package main

import (
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {
	rand.Seed(time.Now().UnixNano())
	data := util.LoadDataFile(util.Call7Vehicle3)

	// time program exection
	defer util.NewTimer().PrintElapsed()

	solution := util.FeasibleSolution()
	util.PrintSolution(solution)

	for i := 0; i < 10000; i++ {
		operators.TwoExchange(data, solution)
	}
	// util.PrintSolution(newSol)
}
