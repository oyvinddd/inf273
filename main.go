package main

import (
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// load data from file
	data := util.LoadDataFile(util.Call7Vehicle3)

	// benchmark program exection
	defer util.NewTimer().PrintElapsed()

	solution := util.FeasibleTestSolution()
	util.PrintSolution(solution)

	// s1 := solution[1]
	newSolution := operators.TwoExchange(data, solution)

	util.PrintSolution(newSolution)
}
