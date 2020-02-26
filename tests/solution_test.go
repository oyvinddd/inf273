package tests

import (
	"testing"

	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
	"github.com/oyvinddd/inf273/util"
)

func TestGenerateSolution(t *testing.T) {
	solution := GenerateSolution(getData())

	if solution == nil {
		t.Errorf("Solution is nil")
	}

	validCallNum, callNum := 11, 0
	for i := range solution {
		callNum += len(solution[i])
	}

	if callNum != validCallNum {
		t.Errorf("Number of calls is incorrect: %d (should be %d)", callNum, validCallNum)
	}
}

func TestCheckFeasibility(t *testing.T) {
	// TODO:
}

func TestCalculateObjective(t *testing.T) {

	data := getData()
	solution := getFeasibleSolution()

	solutionLength := len(solution)
	if solutionLength != 4 {
		t.Errorf("Length of solution is incorrect: %d", solutionLength)
	}

	objective := CalculateObjective(data, solution)
	if objective != 1940470 {
		t.Errorf("Objective function is wrong: %v (should be %v)", objective, 1940470)
	}
}

func getData() models.INF273Data {
	data, _ := util.ParseFile("Call_7_Vehicle_3.txt")
	data.Vehicles = append(data.Vehicles, *models.NewDummyVehicle())
	data.NoOfVehicles++
	return data
}

func getFeasibleSolution() [][]*models.Call {

	c1 := models.NewCall(1, 17, 37, 4601, 790000, 345, 417, 345, 1006)
	c2 := models.NewCall(2, 33, 36, 13444, 430790, 96, 168, 96, 529)
	c3 := models.NewCall(3, 17, 27, 6596, 200354, 715, 787, 715, 1089)
	c4 := models.NewCall(4, 6, 1, 11052, 275455, 0, 72, 0, 435)
	c5 := models.NewCall(5, 38, 33, 6643, 642740, 107, 179, 107, 593)
	c6 := models.NewCall(6, 10, 38, 14139, 587198, 300, 372, 300, 801)
	c7 := models.NewCall(7, 26, 6, 5310, 359885, 178, 250, 178, 567)

	return [][]*models.Call{
		{c3, c3},
		{c7, c1, c7, c1},
		{c5, c5},
		{c2, c4, c6}, // not transported
	}
}
