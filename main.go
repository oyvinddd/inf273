package main

import (
	"fmt"

	"github.com/oyvinddd/inf273/assignment2"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

// ------------- RUN THE PROGRAM -------------

func main() {

	// time program exection
	defer util.NewTimer().PrintElapsed()

	// load data from file and create initial outsourced solution
	data, _ := util.ParseFile("data/Call_7_Vehicle_3.txt", true)
	solution := assignment2.CreateOutsourcedSolution(data)

}

func printSlice(s []*models.Call) {
	for i := range s {
		if s[i] != nil {
			fmt.Printf("[ %v ]", s[i].Index)
		} else {
			fmt.Println("[ X ]")
		}
	}
	fmt.Println()
}

var testSol []*models.Call = []*models.Call{
	&models.Call{Index: 1},
	&models.Call{Index: 2},
	&models.Call{Index: 3},
	&models.Call{Index: 4},
	&models.Call{Index: 5},
	&models.Call{Index: 6},
	&models.Call{Index: 7},
	&models.Call{Index: 8},
	&models.Call{Index: 9},
}
