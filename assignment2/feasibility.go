package a2

import (
	"fmt"

	"github.com/oyvinddd/inf273/models"
)

// CheckTimeFeasibility checks if time windows for a given call is feasible
func CheckTimeFeasibility(d models.INF273Data, v models.Vehicle, s []*models.Call) error {
	if noOfCalls := len(s); noOfCalls > 1 && !v.IsDummy() {

		visited := make(map[int]bool)

		// 1. travel time from home node to first call (+ starting time)
		time := d.GetTravelTimeAndCost(v.Home, s[0].Origin, v.Index).Time + v.StartTime

		// 2. calculate travel cost between each node
		for i := 0; i < noOfCalls-1; i++ {
			c1, c2 := s[i], s[i+1]
			n1, n2 := 0, 0
			if !visited[c1.Index] {
				visited[c1.Index] = true
				n1 = c1.Origin
				if time < c1.LowerPW {
					time += (c1.LowerPW - time)
				} else if time > c1.UpperPW {
					return fmt.Errorf("Infeasible! Arrival time at %v (origin of %v) is %v, but time window is [%v, %v]", c1.Origin, c1.Index, time, c1.LowerPW, c1.UpperPW)
				}
				time += d.GetNodeTimeAndCost(v.Index, c1.Index).OriginTime
			} else {
				n1 = c1.Destination
				if time < c1.LowerDW {
					time += (c1.LowerDW - time)
				} else if time > c1.UpperDW {
					return fmt.Errorf("Infeasible! Arrival time at %v (destination of %v) is %v, but time window is [%v, %v]", c1.Destination, c1.Index, time, c1.LowerDW, c1.UpperDW)
				}
				time += +d.GetNodeTimeAndCost(v.Index, c1.Index).DestinationTime
			}
			if !visited[c2.Index] {
				n2 = c2.Origin
			} else {
				n2 = c2.Destination
			}
			time += d.GetTravelTimeAndCost(n1, n2, v.Index).Time
		}

		// edge case: last call
		c1 := s[noOfCalls-1]
		if time > c1.UpperDW {
			return fmt.Errorf("Infeasible! Arrival time at %v (destination of %v) is %v, but time window is [%v, %v]", c1.Destination, c1.Index, time, c1.LowerDW, c1.UpperDW)
		}
	}
	return nil
}
