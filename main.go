package main

import (
	"fmt"
	"math/rand"
	"time"

	a2 "github.com/oyvinddd/inf273/assignment2"
	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func runInstance(instance datafiles.INF273Instance, seconds float64) {
	fmt.Printf("\n----- START INSTANCE %v -----\n", string(instance))

	data := util.LoadDataFile(instance)
	s0 := a2.GenerateOutsourcedSolution(data)
	o0 := a2.TotalObjective(data, s0)

	opTimer := util.NewTimer()
	s1 := heuristics.Adaptive(data, s0, seconds)
	opTimer.End(fmt.Sprintf("Running time for instance %v", string(instance)))

	o1 := a2.TotalObjective(data, s1)

	util.PrintSolutionAndObj(s1, o1)
	fmt.Printf("Improvement: %v\n", (100.0 * (float32(o0) - float32(o1)) / float32(o0)))
	fmt.Printf("\n----- END INSTANCE %v -----\n", string(instance))
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// benchmark total program exection time
	defer util.NewTimer().PrintElapsed()

	runInstance(datafiles.Call7Vehicle3, 10)
	runInstance(datafiles.Call18Vehicle5, 20)
	runInstance(datafiles.Call35Vehicle7, 50)
	runInstance(datafiles.Call80Vehicle20, 120)
	runInstance(datafiles.Call130Vehicle40, 400)
}
