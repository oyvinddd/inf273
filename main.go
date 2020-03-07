package main

import (
	"github.com/oyvinddd/inf273/heuristics"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {

	// time program exection
	defer util.NewTimer().PrintElapsed()

	heuristics.RandomSearch(initialSolution())
}

func initialSolution() [][]*models.Call {
	c1 := models.NewCall(1, 17, 37, 4601, 790000, 345, 417, 345, 1006)
	c2 := models.NewCall(2, 33, 36, 13444, 430790, 96, 168, 96, 529)
	c3 := models.NewCall(3, 17, 27, 6596, 200354, 715, 787, 715, 1089)
	c4 := models.NewCall(4, 6, 1, 11052, 275455, 0, 72, 0, 435)
	c5 := models.NewCall(5, 38, 33, 6643, 642740, 107, 179, 107, 593)
	c6 := models.NewCall(6, 10, 38, 14139, 587198, 300, 372, 300, 801)
	c7 := models.NewCall(7, 26, 6, 5310, 359885, 178, 250, 178, 567)

	return [][]*models.Call{
		{c3, c3},
		{c7, c1, c7, c1},
		{c5, c5},
		{c2, c4, c6}, // not transported
	}
}
