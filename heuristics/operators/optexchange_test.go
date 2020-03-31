package operators

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestRemoveCall(t *testing.T) {
	solution := util.FeasibleTestSolution()
	calls1 := solution[0]
	noOfCalls1 := len(calls1)
	removeCall(&calls1, calls1[0])
	if len(calls1) != noOfCalls1-2 {
		t.Errorf("Wrong number of calls after removal %v (should be %v)", len(calls1), noOfCalls1-2)
	}
	calls2 := solution[1]
	noOfCalls2 := len(calls2)
	removeCall(&calls2, calls2[0])
	if len(calls2) != noOfCalls2-2 {
		t.Errorf("Wrong number of calls after removal %v (should be %v)", len(calls2), noOfCalls2-2)
	}
	if calls2[0].Index != 1 || calls2[1].Index != 1 {
		t.Errorf("Wrong calls remaining after removal %v (should be %v)", calls2[0].Index, 1)
	}
}

func TestCostForVehicleAndCall(t *testing.T) {
	solution := util.FeasibleTestSolution()
	expCost := 51810
	cost := costForVehicleAndCall(data, data.Vehicles[0], solution[0][0])
	if cost != expCost {
		t.Errorf("Wrong cost %v (should be %v)", cost, expCost)
	}
}
