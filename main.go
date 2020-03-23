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

	// benchmark program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file
	data := util.LoadDataFile(util.Call18Vehicle5)

	s := assignment2.CreateOutsourcedSolution(data)
	s2, obj := heuristics.RandomSearch(data, s)

	util.PrintSolution(s2)
	fmt.Println(obj)
}
