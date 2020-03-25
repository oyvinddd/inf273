package a2

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestCheckTimeWindows(t *testing.T) {

	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()

	err := CheckTotalTimeWindows(data, s1)
	err = CheckTotalTimeWindows(data, s2)
	err = CheckTotalTimeWindows(data, s3)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckCapacity(t *testing.T) {
	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()
	for index, calls := range s1 {
		vehicle := data.Vehicles[index]
		if err := CheckCapacity(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s2 {
		vehicle := data.Vehicles[index]
		if err := CheckCapacity(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s3 {
		vehicle := data.Vehicles[index]
		if err := CheckCapacity(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
}

func TestCheckCompatibility(t *testing.T) {
	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()
	for index, calls := range s1 {
		vehicle := data.Vehicles[index]
		if err := CheckCompatibility(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s2 {
		vehicle := data.Vehicles[index]
		if err := CheckCompatibility(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s3 {
		vehicle := data.Vehicles[index]
		if err := CheckCompatibility(data, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
}
