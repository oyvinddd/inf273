package main

import (
	"fmt"
	"math/rand"
	"time"

	a2 "github.com/oyvinddd/inf273/assignment2"
	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// benchmark program exection
	defer util.NewTimer().PrintElapsed()

	// load data file and generate outsourced solution
	// data := util.LoadDataFile(datafiles.Call7Vehicle3)
	// s0 := a2.GenerateOutsourcedSolution(data)
	// fmt.Printf("Total Distance of solution: %v", operators.TotalDistance(data, s0))
	// var s1 [][]*models.Call
	// for i := 0; i < 1000; i++ {
	// 	s1 = operators.HomeClustering(data, s0)
	// 	s0 = util.CopySolution(s1)
	// 	//util.PrintSolution(s1)
	// }
	// util.PrintSolution(s0)

	// fmt.Println(a2.CheckFeasibility(data, s0))

	data := util.LoadDataFile(datafiles.Call35Vehicle7)
	s0 := a2.GenerateOutsourcedSolution(data)
	o0 := a2.TotalObjective(data, s0)
	s1 := heuristics.SA(data, s0)
	util.PrintSolution(s1)
	util.PrintFlatSolution(s1)
	fmt.Printf("Obj.: %v", a2.TotalObjective(data, s1))

	var solutions [][][]*models.Call = make([][][]*models.Call, 10)
	var result []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		// s1 = heuristics.RandomSearch(data, s0)
		// s1 = heuristics.LocalSearch(data, s0)
		s1 := heuristics.SA(data, s0)
		solutions[i] = s1
		result[i] = a2.TotalObjective(data, s1)
	}
	util.PrintResult(result, solutions, o0)
}
