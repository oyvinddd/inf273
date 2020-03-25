package a2

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestCheckTimeFeasibility(t *testing.T) {

	s1 := util.FeasibleTestSolution()

	for i, call := range s1 {
		vehicle := data.Vehicles[i]
		if err := CheckTimeFeasibility(data, vehicle, call); err != nil {
			t.Error(err)
		}
	}
}
