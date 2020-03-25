package main

import (
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics"

	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// // benchmark program exection
	//defer util.NewTimer().PrintElapsed()

	// // load data file and generate outsourced solution
	data := util.LoadDataFile(datafiles.Call130Vehicle40)
	s0 := assignment2.GenerateOutsourcedSolution(data)

	s1, obj, _, _, _ := heuristics.SA(data, s0)

	
}
