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
	data := util.LoadDataFile(util.Call80Vehicle20)
	solution := util.FeasibleTestSolution() //assignment2.CreateOutsourcedSolution(data)
	// s := heuristics.SA(data, solution)
	// for i := 0; i < 10; i++ {
	// 	// _, obj := heuristics.RandomSearch(data, solution)
	// 	_, obj := heuristics.LocalSearch(data, solution)
	// 	fmt.Println(obj)
	// }
	s := operators.ThreeExchange(data, solution)
	util.PrintSolution(solution)
	util.PrintSolution(s)
}
