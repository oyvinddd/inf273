package operators

import (
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

	solution := util.FeasibleTestSolution()
	callCount := noOfCallsInSolution(solution)
	removedCall := removeCall(&solution[0])
	newCallCount := noOfCallsInSolution(solution)

	if newCallCount != callCount-2 {
		t.Errorf("Wrong number of calls in solution: %v (should be %v)", newCallCount, callCount-2)
	}

	if solutionHasNilCalls(solution) {
		t.Errorf("Solution has one or more nil calls")
	}

	if solutionContainsCall(solution[0], removedCall) {
		t.Errorf("Solution contains call %v after deletion", removedCall.Index)
	}
}

func TestInsertCall(t *testing.T) {

	solution := util.FeasibleTestSolution()
	callCount := noOfCallsInSolution(solution)
	insertCall(data, data.Vehicles[0], &solution[0], &models.Call{Index: 1})
	newCallCount := noOfCallsInSolution(solution)

	if newCallCount != callCount+2 {
		t.Errorf("Wrong number of calls in solution: %v (should be %v)", newCallCount, callCount+2)
	}
}

func TestRandomIndices(t *testing.T) {
	maxExc := 10
	for i := 0; i < 1000; i++ {
		r1, r2 := randomIndices(maxExc)
		if r1 == r2 {
			t.Errorf("Indices cannot be the same: %v and %v", r1, r2)
		}
		if r1 > 9 || r2 > 9 {
			t.Errorf("Indices cannot be %v or bigger: %v and %v", maxExc, r1, r2)
		}
		if r1 < 0 || r2 < 0 {
			t.Errorf("Indices cannot be below 0: %v and %v", r1, r2)
		}
	}
}

func TestIndexOfOptimalDelivery(t *testing.T) {
	solution := util.FeasibleTestSolution()
	vehicle := data.Vehicles[1]
	calls := solution[1]                    // 7 1 7 1
	calls[1], calls[2] = calls[2], calls[1] // 7 7 1 1
	index := indexOfOptimalDelivery(data, vehicle, calls, 0)
	if index != 2 {
		t.Errorf("Wrong optimal index for 7 7 1 1: %v (should be 2)", index)
	}

	index = indexOfOptimalDelivery(data, vehicle, calls, 2)
	if index != 3 {
		t.Errorf("Wrong optimal index for 7 7 1 1: %v (should be 3)", index)
	}

	vehicle = data.Vehicles[0]
	calls = solution[0] // 3 3
	index = indexOfOptimalDelivery(data, vehicle, calls, 1)
	if index != 1 {
		t.Errorf("Wrong optimal index for 3 3: %v (should be 1)", index)
	}

	solution = util.FeasibleTestSolution()
	vehicle = data.Vehicles[1]
	s1 := solution[1]
	s1[2], s1[3] = s1[3], s1[2] // 7 1 1 7

	index = indexOfOptimalDelivery(data, vehicle, s1, 1)
	if index != 3 {
		t.Errorf("Wrong optimal index for 7 1 1 7: %v (should be 3)", index)
	}
}

func TestRightShiftAndInsert(t *testing.T) {
	solution := util.FeasibleTestSolution()
	s1 := solution[1]
	if s1[1].Index != 1 && s1[2].Index != 7 {
		t.Errorf("Calls are misplaced")
	}

	rightShiftAndInsert(s1, 1)

	if s1[1].Index != 1 || s1[2].Index != 1 {
		t.Errorf("Call was not inserted at expected index (expected: 7 1 1 7)")
	}
}

func TestAlignPickupAndDelivery(t *testing.T) {

	solution := util.FeasibleTestSolution()
	calls := solution[1]
	call := calls[1]
	alignPickupAndDelivery(calls, call)
	if calls[1] != call || calls[2] != call {
		t.Errorf("Calls are not aligned")
	}

	calls = solution[3]
	call = calls[2]
	alignPickupAndDelivery(calls, call)
	if calls[2] != call || calls[3] != call {
		t.Errorf("Calls are not aligned")
	}

	call = calls[5] // call #6
	alignPickupAndDelivery(calls, call)
	if calls[4] != call || calls[5] != call {
		t.Errorf("Calls are not aligned")
	}
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

func solutionContainsCall(s []*models.Call, c *models.Call) bool {
	for _, call := range s {
		if call == c {
			return true
		}
	}
	return false
}
