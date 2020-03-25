package a2

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestCheckTimeFeasibility(t *testing.T) {

	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()

	err := CheckTotalTimeFeasibility(data, s1)
	err = CheckTotalTimeFeasibility(data, s2)
	err = CheckTotalTimeFeasibility(data, s3)
	if err != nil {
		t.Error(err)
	}
}
