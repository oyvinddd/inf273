package operators

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestSwapCalls(t *testing.T) {
	solution := util.FeasibleTestSolution()
	calls := solution[1]

	calls[0], calls[1] = calls[1], calls[0]
	calls[0], calls[1] = calls[1], calls[0]

	exp := 7
	if calls[0].Index != exp {
		t.Errorf("Wrong call at expected position: %v (wanted %v)", calls[0].Index, exp)
	}
}
