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

	var s1 [][]*models.Call = nil
	var result []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		//_, obj, _, _, _ := heuristics.SA(data, s0)
		s1 = heuristics.LocalSearch(data, s0)
		result[i] = a2.TotalObjective(data, s1)
	}
	util.PrintSolution(s1)
	util.PrintResult(result, o0)

}
