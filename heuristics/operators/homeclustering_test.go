package operators

import (
	"math"
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestDistance(t *testing.T) {
	s1 := util.FeasibleTestSolution()
	d := distance(data, data.Vehicles[0], s1[0][0])
	if d != 66 {
		t.Errorf("Wrong distance for vehicle %d and call %d: %v (should be 66)", data.Vehicles[0].Index, s1[0][0].Index, d)
	}
	d = distance(data, data.Vehicles[0], data.GetCall(4))
	if d != math.MaxInt32 {
		t.Errorf("Wrong distance for vehicle %d and call %d: %v (should be %d)", data.Vehicles[0], data.GetCall(4).Index, d, math.MaxInt32)
	}
}

func TestRandomNonEmptyRoute(t *testing.T) {
	s1 := util.FeasibleTestSolution()
	for i := 0; i < 100; i++ {
		calls, index := randomNonEmptyRoute(data, s1)
		if len(calls) == 0 {
			t.Errorf("Random route is empty")
		}
		if calls[0] != s1[index][0] {
			t.Errorf("Returned route doesn't correspond to the index: %v", index)
		}
	}
}

func TestRemoveWorstCall(t *testing.T) {
	s1 := util.FeasibleTestSolution()
	worstCall, index := removeWorstCall(data, s1)
	vehicle := data.Vehicles[index]
	calls := s1[index]
	if solutionContainsCall(calls, worstCall) {
		t.Errorf("Solution contains removed call %v", worstCall.Index)
	}
	for _, call := range calls {
		if distance(data, vehicle, call) > distance(data, vehicle, worstCall) {
			t.Errorf("Call %v is not the worst. Distance of call %v is worse", worstCall.Index, call.Index)
		}
	}
}
