package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/heuristics"

	"github.com/oyvinddd/inf273/assignment2"

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
	solution := assignment2.CreateOutsourcedSolution(data)
	s := heuristics.LocalSearch(data, solution)
	fmt.Printf("OBJ: %v\n", assignment2.CalcTotalObjective(data, s))
	util.PrintSolution(s)
}
