package heuristics

import (
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/models"
)

type operator func(models.INF273Data, [][]*models.Call) [][]*models.Call

// Adaptive is an adaptive meta-heuristic framework
func Adaptive(data models.INF273Data, s0 [][]*models.Call) [][]*models.Call {
	current, best := s0, s0
	var ops []operator = ops()
	i := 0
	for i < adMaxIterations {
		// TODO: escape condition to get out of local optima
		// Select a heuristic based on selection parameters

		current = ops[rand.Intn(3)](data, current)
		// TODO: apply heuristic to best
		if a2.TotalObjective(data, current) < a2.TotalObjective(data, best) {
			best = current
		}
		i++
	}
	return best
}

func ops() []operator {
	return []operator{
		operators.OptExchange,
		operators.WeightedReinsert,
		operators.HomeClustering,
	}
}
