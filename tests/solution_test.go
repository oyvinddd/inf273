package tests

import (
	"log"
	"testing"

	"github.com/oyvinddd/inf273/models"
	. "github.com/oyvinddd/inf273/solution"
	"github.com/oyvinddd/inf273/util"
)

func TestDataSet(t *testing.T) {

	data := getData()
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

func TestCheckFeasibility(t *testing.T) {
	// TODO: implement
}

func TestCalculateObjective(t *testing.T) {

	data := getData()
	solution := getFeasibleSolution()
	expObj := 1940470

	objective := CalculateObjective(data, solution)
	if objective != expObj {
		t.Errorf("Objective function is wrong: %v (should be %v)", objective, expObj)
	}
}

func getData() models.INF273Data {
	data, err := util.ParseFile("../data/Call_7_Vehicle_3.txt")
	if err != nil {
		log.Fatal(err)
	}
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
