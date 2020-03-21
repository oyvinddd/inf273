package heuristics

import (
	"fmt"
	"os"
	"testing"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

var data models.INF273Data

func TestMain(m *testing.M) {
	data = util.TestData()
	os.Exit(m.Run())
}

func TestRemoveCall(t *testing.T) {

	solution := util.FeasibleSolution()
	callCount := noOfCallsInSolution(solution)
	removeCall(&solution[0])
	newCallCount := noOfCallsInSolution(solution)

	if newCallCount != callCount-2 {
		t.Errorf("Wrong number of calls in solution: %v (should be %v)", newCallCount, callCount-2)
	}

	if solutionHasNilCalls(solution) {
		t.Errorf("Solution has one or more nil calls")
	}
}

func TestInsertCall(t *testing.T) {

	solution := util.FeasibleSolution()
	callCount := noOfCallsInSolution(solution)
	insertCall(data, data.Vehicles[0], &solution[0], &models.Call{})
	newCallCount := noOfCallsInSolution(solution)

	if newCallCount != callCount+2 {
		t.Errorf("Wrong number of calls in solution: %v (should be %v)", newCallCount, callCount+2)
	}
}

func TestFindOptimalDelivery(t *testing.T) {
	solution := util.FeasibleSolution()
	vehicle := data.Vehicles[1]
	vehicleCalls := solution[1]

	findOptimalDelivery(data, vehicle, &vehicleCalls, 1)

	fmt.Println(vehicleCalls)
	// TODO:
}

// ------- HELPER FUNCTIONS -------

func noOfCallsInSolution(s [][]*models.Call) int {
	var noOfCalls int = 0
	for row := range s {
		for range s[row] {
			noOfCalls++
		}
	}
	return noOfCalls
}

func solutionHasNilCalls(s [][]*models.Call) bool {
	for row := range s {
		for col := range s[row] {
			if s[row][col] == nil {
				return true
			}
		}
	}
	return false
}
