package main

import (
	"math/rand"
	"time"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/models"

	datafiles "github.com/oyvinddd/inf273/data"
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
	data := util.LoadDataFile(datafiles.Call18Vehicle5)
	s0 := a2.GenerateOutsourcedSolution(data)
	o0 := a2.TotalObjective(data, s0)

	var x []float64
	var y []float64
	var z []float64
	var s1 [][]*models.Call = nil
	var result []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		// s1 := heuristics.RandomSearch(data, s0)
		s1, xx, yy, zz := heuristics.SA(data, s0)
		x = xx
		y = yy
		z = zz
		//s1 = heuristics.LocalSearch(data, s0)
		result[i] = a2.TotalObjective(data, s1)
	}
	util.PrintSolution(s1)
	util.PrintResult(result, o0)
	util.GenerateGraph(x, y, "temp.png")
	util.GenerateGraph(x, z, "p.png")
}
