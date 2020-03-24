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
	data := util.LoadDataFile(util.Call130Vehicle40)
	s0 := assignment2.GenerateOutsourcedSolution(data)
	obj0 := assignment2.CalcTotalObjective(data, s0)

	results := make([]int, 10)
	for i := 0; i < 10; i++ {
		// _, obj := heuristics.RandomSearch(data, s0)
		//_, obj := heuristics.LocalSearch(data, s0)
		_, obj := heuristics.SA(data, s0)
		//util.PrintSolution(s1)
		results[i] = obj
	}
	printResults(results, obj0)
	s1, obj := heuristics.SA(data, s0)
	util.PrintSolution(s1)
	fmt.Println(obj)
}

func printResults(r []int, obj0 int) {
	fmt.Printf("INITIAL: %v\n", obj0)
	fmt.Println("ALL RESULTS:")
	best := r[0]
	sum := 0
	for i := range r {
		if r[i] < best {
			best = r[i]
		}
		sum += r[i]
		fmt.Println(r[i])
	}
	fmt.Printf("AVG. OBJ: %v\n", sum/10.0)
	fmt.Printf("BEST: %v\n", best)
	fmt.Printf("IMPR.: %v\n", (100.0 * (float32(obj0) - float32(best)) / float32(obj0)))
}
