package operators

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestLowerTimeWindow(t *testing.T) {
	solution := util.FeasibleTestSolution()
	calls := solution[1]
	pickups := findPickups(calls)
	ltw1 := lowerTimeWindow(calls[0], 0, pickups)
	if ltw1 != 178 {
		t.Errorf("Wrong lower time window for call %v: %v (should be %v)", calls[0].Index, ltw1, 178)
	}
	ltw2 := lowerTimeWindow(calls[3], 3, pickups)
	if ltw2 != 345 {
		t.Errorf("Wrong lower time window for call %v: %v (should be %v)", calls[3].Index, ltw2, 178)
	}
}
