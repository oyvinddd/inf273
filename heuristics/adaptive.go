package heuristics

import (
	"fmt"
	"math"
	"math/rand"

	a2 "github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/heuristics/operators"
	"github.com/oyvinddd/inf273/models"
	"github.com/oyvinddd/inf273/util"
)

const (
	highScore  float32 = 8
	smallScore float32 = 6
)

type operator func(models.INF273Data, [][]*models.Call) [][]*models.Call

// Adaptive is an adaptive meta-heuristic framework
func Adaptive(data models.INF273Data, s0 [][]*models.Call) [][]*models.Call {
	incumbent, best := s0, s0
	var ops []operator = ops()
	var seg = newSegment()

	var iterationsSinceBest = 0
	var deSum float64 = 0
	var deNum float64 = 0

	var T float64 = 1000     // temperature
	var a float64 = 0.998765 // cooling factor
	var p float64 = 0.8      // probability of accepting worse solution

	var u1 int = 0
	var u2 int = 0
	var u3 int = 0
	var u4 int = 0

	for i := 0; i < adMaxIterations; i++ {

		if iterationsSinceBest >= 500 {
			// for i := 0; i < 5; i++ {
			// 	incumbent = operators.WeightedReinsert(data, incumbent)
			// }
			iterationsSinceBest = 0
		}

		// this condition will pass at the start of each segment
		if i%100 == 0 && i > 0 {
			fmt.Println(seg)
			// calculate the new weights based on the previous weights
			seg.calculateWeights()
			// reset scores and # times used at each segment
			//seg.reset()
			seg.scores = []float32{0, 0, 0, 0}
			seg.usage = []float32{0, 0, 0, 0}
		}

		index := randomOperatorIndex(seg)
		if index == 0 {
			u1++
		} else if index == 1 {
			u2++
		} else if index == 2 {
			u3++
		} else if index == 3 {
			u4++
		}
		newSolution := ops[index](data, incumbent)
		seg.incrementTimesUsed(index)

		deltaE := float64(a2.TotalObjective(data, newSolution) - a2.TotalObjective(data, incumbent))

		if i < 100 && deltaE >= 0 {
			deSum += deltaE
			deNum++
		} else if i == 100 {
			avgDeltas := deSum / deNum
			T = -avgDeltas / math.Log(0.8)
		} else {
			p = math.Exp(-deltaE / T)
		}

		isFeasible := a2.IsFeasible(data, newSolution)

		if isFeasible && deltaE < 0 {
			incumbent = newSolution
			if a2.TotalObjective(data, incumbent) < a2.TotalObjective(data, best) {
				best = incumbent
				seg.addScore(index, highScore)
				iterationsSinceBest = 0
			}
		} else if isFeasible && rand.Float64() < p {
			incumbent = newSolution
			seg.addScore(index, smallScore)
		}
		T *= a
		iterationsSinceBest++
	}
	fmt.Println(u1, u2, u3, u4)
	return best
}

// -------- SEGMENT STRUCT W/ FUNCTIONS --------

type segment struct {
	scores  []float32
	weights []float32
	usage   []float32
	r       float32
}

func newSegment() segment {
	return segment{
		weights: []float32{
			100, 100, 100, 100,
		},
		scores: []float32{
			0, 0, 0, 0,
		},
		usage: []float32{
			0, 0, 0, 0,
		},
		r: 0.1,
	}
}

func (s *segment) addScore(heuristicIndex int, score float32) {
	s.scores[heuristicIndex] += score
}

func (s *segment) incrementTimesUsed(heuristicIndex int) {
	s.usage[heuristicIndex]++
}

func (s *segment) calculateWeights() {
	for i := 0; i < len(s.weights); i++ {
		if s.usage[i] > 0 {
			s.weights[i] = s.weights[i]*(1-s.r) + s.r*(s.scores[i]/s.usage[i])
			if s.weights[i] < 0.001 {
				s.weights[i] = 0.001
			}
		}
	}
}

func (s *segment) sumOfWeights() float32 {
	var w float32 = 0
	for _, wg := range s.weights {
		w += wg
	}
	return w
}

func (s *segment) reset() {
	w := s.weights
	*s = newSegment()
	(*s).weights = w
}

func (s segment) String() string {
	return fmt.Sprintf("Segment stats:\nW: %v\nS: %v\nU: %v\n", s.weights, s.scores, s.usage)
}

// -------- HELPER FUNCTIONS --------

func ops() []operator {
	return []operator{
		operators.TwoExchange,
		operators.OptOrdering,
		operators.OneReinsert,
		operators.WeightedReinsert,
	}
}

func randomOperatorIndex(s segment) int {
	// use roulette wheel principle to create an array of weights
	sum := s.sumOfWeights()
	weights := make([]float32, len(s.weights))
	for i, wg := range s.weights {
		weights[i] = wg / sum
	}
	// return a random weighted index
	return util.WeightedRandomNumber(weights)
}
