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
