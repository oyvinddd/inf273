package assignment2

import (
	"os"
	"testing"

	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

var data models.INF273Data

func TestMain(m *testing.M) {
	data = util.TestData()
	os.Exit(m.Run())
}

// --------- UNIT TESTS FOR SOME OF THE FUNCTIONS USED IN THE PROGRAM ---------

func TestDataSet(t *testing.T) {

	v1 := data.Vehicles[0]
	v2 := data.Vehicles[1]
	v3 := data.Vehicles[2]
	c1 := data.Calls[0]
	c2 := data.Calls[1]
	c3 := data.Calls[2]
	c4 := data.Calls[3]
	c5 := data.Calls[4]
	c6 := data.Calls[5]
	c7 := data.Calls[6]

	// travel cost from home node to first call
	fromHome1 := data.GetTravelTimeAndCost(v1.Home, c3.Origin, v1.Index).Cost
	fromHome2 := data.GetTravelTimeAndCost(v2.Home, c7.Origin, v2.Index).Cost
	fromHome3 := data.GetTravelTimeAndCost(v3.Home, c5.Origin, v3.Index).Cost
	fromHomeSum := fromHome1 + fromHome2 + fromHome3
	expFromHome1 := 20854
	expFromHome2 := 25628
	expFromHome3 := 84492
	if fromHome1 != expFromHome1 {
		t.Errorf("Wrong travel cost from home for 1: %d (expected: %d)", fromHome1, expFromHome1)
	}
	if fromHome2 != expFromHome2 {
		t.Errorf("Wrong travel cost from home for 2: %d (expected: %d)", fromHome2, expFromHome2)
	}
	if fromHome3 != expFromHome3 {
		t.Errorf("Wrong travel cost from home for 3: %d (expected: %d)", fromHome3, expFromHome3)
	}

	// unhandled calls
	penalty := c2.Penalty + c4.Penalty + c6.Penalty
	expectedPenalty := 1293443

	if penalty != expectedPenalty {
		t.Errorf("Total penalty for unhandled calls is wrong: %d (expected: %d)", penalty, expectedPenalty)
	}

	// vehicle #1 travel cost
	v1tc := data.GetTravelTimeAndCost(c3.Origin, c3.Destination, 1).Cost
	expV1TravelCost := 5792
	if v1tc != expV1TravelCost {
		t.Errorf("Wrong transport cost for 1: %d (expected: %d)", v1tc, expV1TravelCost)
	}

	// vehicle #2 travel cost
	v2tc1 := data.GetTravelTimeAndCost(c7.Origin, c1.Origin, 2).Cost
	v2ct2 := data.GetTravelTimeAndCost(c1.Origin, c7.Destination, 2).Cost
	v2ct3 := data.GetTravelTimeAndCost(c7.Destination, c1.Destination, 2).Cost
	v2TravelCostSum := v2tc1 + v2ct2 + v2ct3
	expV2TravelCost := 221131
	if v2TravelCostSum != expV2TravelCost {
		t.Errorf("Wrong transport cost for 2: %d (expected: %d)", v2TravelCostSum, expV2TravelCost)
	}

	// vehicle #3 travel cost
	v3tc := data.GetTravelTimeAndCost(c5.Origin, c5.Destination, 3).Cost
	expV3TravelCost := 65817
	if v3tc != expV3TravelCost {
		t.Errorf("Wrong transport cost for 3: %d (expected: %d)", v3tc, expV3TravelCost)
	}

	// origin and destination node costs
	nc1 := data.GetNodeTimeAndCost(v1.Index, c3.Index).OriginCost
	nc2 := data.GetNodeTimeAndCost(v1.Index, c3.Index).DestinationCost
	nc3 := data.GetNodeTimeAndCost(v2.Index, c7.Index).OriginCost
	nc4 := data.GetNodeTimeAndCost(v2.Index, c7.Index).DestinationCost
	nc5 := data.GetNodeTimeAndCost(v2.Index, c1.Index).OriginCost
	nc6 := data.GetNodeTimeAndCost(v2.Index, c1.Index).DestinationCost
	nc7 := data.GetNodeTimeAndCost(v3.Index, c5.Index).OriginCost
	nc8 := data.GetNodeTimeAndCost(v3.Index, c5.Index).DestinationCost
	nodeCostSum := nc1 + nc2 + nc3 + nc4 + nc5 + nc6 + nc7 + nc8
	expNodeCostSum := 223313
	if nodeCostSum != expNodeCostSum {
		t.Errorf("Wrong origin/destination cost: %d (expected: %d)", nodeCostSum, expNodeCostSum)
	}

	expObj := 1940470
	sumCost := fromHomeSum + nodeCostSum + penalty + v1tc + v2TravelCostSum + v3tc
	if sumCost != expObj {
		t.Errorf("Wrong objective function: %d (expected: %d)", sumCost, expObj)
	}
}

func TestGenerateSolution(t *testing.T) {

	solution := GenerateSolution(data)

	for row := range solution {
		for col := range solution[row] {
			if solution[row][col].PickedUp == true {
				t.Errorf("Generated solution should have no picked up calls")
			}
		}
	}

	for _, call := range data.Calls {
		if !solutionContainsCall(solution, &call) {
			t.Errorf("Not the correct amount of calls in solution")
		}
	}
}

func TestCheckFeasibility(t *testing.T) {
	err := CheckFeasibility(data, util.FeasibleTestSolution())
	if err != nil {
		t.Errorf("Infeasible solution: %v", err)
	}
}

func TestCalculateObjective(t *testing.T) {

	solution := util.FeasibleTestSolution()
	expObj := 1940470

	objective := CalcTotalObjective(data, solution)
	if objective != expObj {
		t.Errorf("Objective function is wrong: %v (should be %v)", objective, expObj)
	}
}

func TestOutsourcedSolution(t *testing.T) {
	solution := CreateOutsourcedSolution(data)
	obj := CalcTotalObjective(data, solution)
	expObj := 3286422
	if obj != expObj {
		t.Errorf("Objective function is wrong: %v (should be %v)", obj, expObj)
	}
}

func solutionContainsCall(solution [][]*models.Call, call *models.Call) bool {
	var count int = 0
	for row := range solution {
		for col := range solution[row] {
			if solution[row][col].Index == call.Index {
				count++
			}
		}
	}
	return count == 2
}
