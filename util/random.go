package util

import "math/rand"

// WeightedRandomNumber returns a random number but takes into account the input weights
func WeightedRandomNumber(weights []float32) int {
	if len(weights) == 0 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}
	r := rand.Float32() * sum
	for i, w := range weights {
		r -= w
		if r < 0 {
			return i
		}
	}
	return len(weights) - 1
}

// TwoRandomIndices returns two unique random indices in range [0, max)
func TwoRandomIndices(max int) (int, int) {
	r1 := rand.Intn(max - 1)
	r2 := rand.Intn(max - 1)
	if r2 >= r1 {
		r2++
	}
	return r1, r2
}

func ThreeRandomIndices(max int) (int, int, int) {
	if max <= 0 {
		return 0, 0, 0
	}
	r1, r2 := TwoRandomIndices(max)
	r3 := rand.Intn(max)
	for r3 == r1 || r3 == r2 {
		r3++
		if r3 == max {
			r3 = 0
		}
	}
	return r1, r2, r3
}
