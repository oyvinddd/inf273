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

func TestRandomIndices(t *testing.T) {
	maxExc := 10
	for i := 0; i < 1000; i++ {
		r1, r2 := randomIndices(maxExc)
		if r1 == r2 {
			t.Errorf("Number cannot be the same: %v and %v", r1, r2)
		}
		if r1 > 9 || r2 > 9 {
			t.Errorf("Number cannot be %v or bigger: %v and %v", maxExc, r1, r2)
		}
		if r1 < 0 || r2 < 0 {
			t.Errorf("Number cannot be below 0: %v and %v", r1, r2)
		}
	}
}

func TestAlignPickupAndDelivery(t *testing.T) {

	solution := util.FeasibleSolution()
	calls := solution[1]
	call := calls[1]
	alignPickupAndDelivery(&calls, call)
	if calls[1] != call || calls[2] != call {
		t.Errorf("Calls are not aligned")
	}

	calls = solution[3]
	call = calls[2]
	alignPickupAndDelivery(&calls, call)
	if calls[2] != call || calls[3] != call {
		t.Errorf("Calls are not aligned")
	}

	call = calls[5] // call #6
	alignPickupAndDelivery(&calls, call)
	if calls[4] != call || calls[5] != call {
		t.Errorf("Calls are not aligned")
	}
}
