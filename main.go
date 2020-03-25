package main

import (
	"math/rand"
	"time"

	a2 "github.com/oyvinddd/inf273/assignment2"
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
	defer util.NewTimer().PrintElapsed()

	// // load data file and generate outsourced solution
	data := util.LoadDataFile(datafiles.Call130Vehicle40)
	s0 := a2.GenerateOutsourcedSolution(data)

	// heuristics.SA(data, s0)
	// heuristics.RandomSearch(data, s0)
	s, o, _, _, _ := heuristics.SA(data, s0)
	util.PrintSolutionAndObj(s, o)
	util.PrintFlatSolution(s)
	// g1 := chart.Chart{
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			XValues: x,
	// 			YValues: y,
	// 		},
	// 	},
	// }
	// f, _ := os.Create("temperature.png")
	// defer f.Close()
	// err := g1.Render(chart.PNG, f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// util.PrintSolutionAndObj(s1, obj)
}
