package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// // benchmark program exection
	defer util.NewTimer().PrintElapsed()

	// // load data file and generate outsourced solution
	data := util.LoadDataFile(util.Call18Vehicle5)
	s0 := assignment2.GenerateOutsourcedSolution(data)

	s1, obj := heuristics.LocalSearch(data, s0)

	util.PrintSolution(s1)
	fmt.Printf("Objective: %v", obj)
}
