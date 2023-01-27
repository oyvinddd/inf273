package main

import (
	"fmt"
	"log"
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

	// benchmark program execution
	defer util.NewTimer().PrintElapsed()

	data := util.LoadDataFile(datafiles.Call7Vehicle3)
	s0 := a2.GenerateOutsourcedSolution(data)
	o0 := a2.TotalObjective(data, s0)

	var solutions [][][]*models.Call = make([][][]*models.Call, 10)
	var result []int = make([]int, 10)
	for i := 0; i < 10; i++ {
		//s1 := heuristics.RandomSearch(data, s0)
		s1 := heuristics.LocalSearch(data, s0)
		// s1 := heuristics.SA(data, s0)
		solutions[i] = s1
		result[i] = a2.TotalObjective(data, s1)

	}
	obj, sss := util.PrintResult(result, solutions, o0)

	if a2.TotalObjective(data, sss) != obj || !a2.IsFeasible(data, sss) {
		log.Fatal("Not in sync!")
	} else {
		fmt.Printf("Double check obj: %v", obj)
	}
}
