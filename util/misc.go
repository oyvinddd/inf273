package util

import (
	"fmt"

	"github.com/oyvinddd/inf273/models"
)

// --------------- MISC HELPER FUNCTIONS ---------------

// PrintSolution prints a given solution to standard output
func PrintSolution(solution [][]*models.Call) {
	fmt.Println()
	fmt.Println("-------- SOLUTION REPRESENTATION --------")
	for i := range solution {
		row := solution[i]
		if len(row) == 0 {
			fmt.Println("[-]")
			continue
		}
		for _, e := range solution[i] {
			if e == nil {
				fmt.Printf("[X]")
			} else {
				fmt.Printf("[%v]", e.Index)
			}
		}
		if i == len(solution)-1 {
			fmt.Print(" <-- Dummy vehicle (unhandled calls)")
		}
		fmt.Println()
	}
	fmt.Println()
}

// PrintRowInSolution prints a given row in the solution
func PrintRowInSolution(solution [][]*models.Call, row int) {
	for _, call := range solution[row] {
		fmt.Printf("[%v]", call.Index)
	}
	fmt.Println()
}

// CopySolution copies a given solution (new pointers to calls are also created)
// not the prettiest solution but required for keeping the order of pointers
func CopySolution(solution [][]*models.Call) [][]*models.Call {
	copied := make([][]*models.Call, len(solution))
	for row := range solution {
		copied[row] = make([]*models.Call, len(solution[row]))
		dict := make(map[int]*models.Call)
		for col, c := range solution[row] {
			ptr := dict[c.Index]
			if ptr == nil {
				ptr = new(models.Call)
				*ptr = *c
				dict[c.Index] = ptr
			}
			copied[row][col] = ptr
		}
	}
	return copied
}

// FeasibleTestSolution returns a feasible (but not optimal) solution (used for testing only)
func FeasibleTestSolution() [][]*models.Call {

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
		{c2, c2, c4, c4, c6, c6}, // not transported
	}
}

// TestData returns a test data set from file
func TestData() models.INF273Data {
	return LoadDataFile(Call7Vehicle3)
}
