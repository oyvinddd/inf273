package a2

import (
	"testing"

	"github.com/oyvinddd/inf273/util"
)

func TestCheckTimeWindows(t *testing.T) {

	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()
	s4 := util.FeasibleTestSolution4()
	s5 := util.FeasibleTestSolution5()

	err := CheckTotalTimeWindows(data, s1)
	err = CheckTotalTimeWindows(data, s2)
	err = CheckTotalTimeWindows(data, s3)
	err = CheckTotalTimeWindows(data2, s4)
	err = CheckTotalTimeWindows(data3, s5)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckCapacity(t *testing.T) {

	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()
	s4 := util.FeasibleTestSolution4()
	s5 := util.FeasibleTestSolution5()

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
	for index, calls := range s4 {
		vehicle := data2.Vehicles[index]
		if err := CheckCapacity(data2, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s5 {
		vehicle := data3.Vehicles[index]
		if err := CheckCapacity(data3, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
}

func TestCheckCompatibility(t *testing.T) {

	s1 := util.FeasibleTestSolution()
	s2 := util.FeasibleTestSolution2()
	s3 := util.FeasibleTestSolution3()
	s4 := util.FeasibleTestSolution4()
	s5 := util.FeasibleTestSolution5()

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
	for index, calls := range s4 {
		vehicle := data2.Vehicles[index]
		if err := CheckCompatibility(data2, vehicle, calls); err != nil {
			t.Error(err)
		}
	}
	for index, calls := range s5 {
		vehicle := data3.Vehicles[index]
		if err := CheckCompatibility(data3, vehicle, calls); err != nil {
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
	if !data2.VehicleAndCallIsCompatible(3, 11) {
		t.Errorf("Error in compatibility! Vehicle 3 and call 11 should be compatible")
	}
	if data2.VehicleAndCallIsCompatible(3, 12) {
		t.Errorf("Error in compatibility! Vehicle 3 and call 12 should not be compatible")
	}
}
