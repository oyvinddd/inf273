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

	if !data.VehicleAndCallIsCompatible(1, 3) {
		t.Errorf("Error in compatibility! Vehicle 1 and call 3 should be compatible")
	}
	if data.VehicleAndCallIsCompatible(1, 6) {
		t.Errorf("Error in compatibility! Vehicle 1 and call 6 should not be compatible")
	}
	if !data.VehicleAndCallIsCompatible(-1, 7) {
		t.Errorf("Error in compatibility! Dummy vehicle should be compatible with all calls")
	}
}
