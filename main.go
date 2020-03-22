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

	// benchmark program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file
	data := util.LoadDataFile(util.Call7Vehicle3)
	solution := util.FeasibleTestSolution()

	ns := operators.TwoExchange(data, solution)

	util.PrintSolution(ns)
}
