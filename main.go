package main

import (
	"fmt"
	"log"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

func main() {

	// 1. Parse input file

	data, err := util.ParseFile("Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Generate a random solution
	solution := generateSolution(data)
	fmt.Println(solution)

	// 3. Check feasibility of solution
}

func generateSolution(data *models.INF273Data) [][]*models.Call {
	var solution [][]*models.Call
	return solution
}

func checkFeasability() {
	// TODO:
}

func calculateObjectiveFunction() {
	// TODO:
}
