package operators

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestIndexOfOptimalVehicle(t *testing.T) {
	solution := util.FeasibleTestSolution()
	c1 := solution[1][1] // call 1
	index := indexOfOptimalVehicle(data, c1)
	if index != 0 {
		t.Errorf("Wrong optimal vehicle for call %v: %v (should be %v)", c1.Index, index, 0)
	}
	c4 := solution[3][2] // call 4
	index = indexOfOptimalVehicle(data, c4)
	if index != 2 {
		t.Errorf("Wrong optimal vehicle for call %v: %v (should be %v)", c4.Index, index, 2)
	}
}
