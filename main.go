package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

var solution [][]int // matrix for solution representation
var timer util.Timer // time the program execution

func main() {

	// start a timer
	timer.Start()
	defer timer.PrintElapsed()

	// 1. Parse input file

	data, err := util.ParseFile("Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	solution = make([][]int, data.NoOfVehicles) // empty solution matrix

	// 2. Generate a random solution

	generateSolution(data, &solution)

	printSolution(solution)

	// 3. Check feasibility of solution
}

// ---------------- ASSIGNMENT #2 ----------------

func generateSolution(data models.INF273Data, matrix *[][]int) {

	// Solution representation:
	//
	//	rows = vehicles, columns = calls
	//	[][][][][][][]
	//	[][][]
	//	[][][][]

	// fill rows with calls (every call appears two times)
	for _, call := range data.Calls {
		i := randomNumber(0, data.NoOfVehicles)
		(*matrix)[i] = append((*matrix)[i], call.Index, call.Index)
	}
	// for each row, randomize the order of the calls
	for _, calls := range *matrix {
		shuffleSlice(calls)
	}
}

func checkFeasability() {
	// TODO:
}

func calculateObj() {
	// TODO:
}

// ---------------- UTILITIES ----------------

func randomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min)
}

func shuffleSlice(a []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i int, j int) {
		a[i], a[j] = a[j], a[i]
	})
}

func printSolution(solution [][]int) {
	for i := range solution {
		for _, e := range solution[i] {
			fmt.Printf("[%v]", e)
		}
		fmt.Println()
	}
}
