package util

import "testing"

func TestCopySolution(t *testing.T) {

	s1 := FeasibleTestSolution()
	s2 := CopySolution(s1)

	if s1[1][0] == s2[1][0] {
		t.Errorf("Call %p exists in both the original and the copied solution", s1[1][0])
	}

	if s2[1][0] != s2[1][2] {
		t.Errorf("Call %p and %p should be the same value", s2[1][0], s2[1][2])
	}
}

func TestTwoRandomIndices(t *testing.T) {
	max := 10
	for i := 0; i < 1000; i++ {
		r1, r2 := TwoRandomIndices(max)
		if r1 == r2 {
			t.Errorf("Indices cannot be the same: %v and %v", r1, r2)
		}
		if r1 > 9 || r2 > 9 {
			t.Errorf("Indices cannot be %v or bigger: %v and %v", max, r1, r2)
		}
		if r1 < 0 || r2 < 0 {
			t.Errorf("Indices cannot be below 0: %v and %v", r1, r2)
		}
	}
}

func TestThreeRandomIndices(t *testing.T) {
	max := 10
	for i := 0; i < 1000; i++ {
		r1, r2, r3 := ThreeRandomIndices(max)
		if r1 == r2 || r1 == r3 || r2 == r3 {
			t.Errorf("Indices cannot be identical: %v, %v and %v", r1, r2, r3)
		}
		if r1 == max || r2 == max || r3 == max {
			t.Errorf("Indices cannot be %v or bigger: %v, %v and %v", max, r1, r2, r3)
		}
	}
}
