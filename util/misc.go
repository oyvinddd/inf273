package util

import (
	"fmt"

	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/models"
)

// --------------- MISC HELPER FUNCTIONS ---------------

// PrintSolutionAndObj prints both the solution and the objective
func PrintSolutionAndObj(s [][]*models.Call, obj int) {
	PrintSolution(s)
	fmt.Printf("Objective: %v\n", obj)
}

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

// PrintFlatSolution prints solution on a single line (vehicles separated by 0)
func PrintFlatSolution(solution [][]*models.Call) {
	fmt.Println()
	fmt.Print("Solution: ")
	for row := range solution {
		for col := range solution[row] {
			callNum := solution[row][col].Index
			fmt.Print(callNum)
			fmt.Print(" ")
		}
		if row < len(solution)-1 {
			fmt.Print(0)
			fmt.Print(" ")
		}
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

// PrintResult prints the result to console
func PrintResult(r []int, s [][][]*models.Call, obj0 int) (int, [][]*models.Call) {
	fmt.Printf("INITIAL: %v\n", obj0)
	best := r[0]
	var s1 [][]*models.Call = s[0]
	sum := 0
	for i := range r {
		if r[i] < best {
			best = r[i]
			s1 = s[i]
		}
		sum += r[i]
	}
	fmt.Printf("AVG. OBJ: %v\n", sum/10.0)
	fmt.Printf("BEST: %v\n", best)
	fmt.Printf("IMPR.: %v\n", (100.0 * (float32(obj0) - float32(best)) / float32(obj0)))

	PrintSolution(s1)
	PrintFlatSolution(s1)
	return best, s1
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

// FeasibleTestSolution2 returns a feasible (but not optimal) solution (used for testing only)
func FeasibleTestSolution2() [][]*models.Call {
	// 1 1 0 7 7 3 3 0 5 5 6 6 0 4 4 2 2
	c1 := models.NewCall(1, 17, 37, 4601, 790000, 345, 417, 345, 1006)
	c2 := models.NewCall(2, 33, 36, 13444, 430790, 96, 168, 96, 529)
	c3 := models.NewCall(3, 17, 27, 6596, 200354, 715, 787, 715, 1089)
	c4 := models.NewCall(4, 6, 1, 11052, 275455, 0, 72, 0, 435)
	c5 := models.NewCall(5, 38, 33, 6643, 642740, 107, 179, 107, 593)
	c6 := models.NewCall(6, 10, 38, 14139, 587198, 300, 372, 300, 801)
	c7 := models.NewCall(7, 26, 6, 5310, 359885, 178, 250, 178, 567)
	return [][]*models.Call{
		{c1, c1},
		{c7, c7, c3, c3},
		{c5, c5, c6, c6},
		{c4, c4, c2, c2},
	}
}

// FeasibleTestSolution3 returns a feasible (but not optimal) solution (used for testing only)
func FeasibleTestSolution3() [][]*models.Call {
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
		{c5, c5, c6, c6},
		{c2, c2, c4, c4},
	}
}

// FeasibleTestSolution4 returns a feasible (but not optimal) solution (used for testing only)
func FeasibleTestSolution4() [][]*models.Call {
	//17 17 10 7 10 3 7 3 0
	// 15 15 11 6 1 11 6 1 0
	// 4 14 9 9 4 14 0
	// 12 12 8 8 0
	// 18 5 5 18 2 2 0
	// 16 16 13 13
	data := TestDataMedium()
	c1 := data.GetCall(1)
	c2 := data.GetCall(2)
	c3 := data.GetCall(3)
	c4 := data.GetCall(4)
	c5 := data.GetCall(5)
	c6 := data.GetCall(6)
	c7 := data.GetCall(7)
	c8 := data.GetCall(8)
	c9 := data.GetCall(9)
	c10 := data.GetCall(10)
	c11 := data.GetCall(11)
	c12 := data.GetCall(12)
	c13 := data.GetCall(13)
	c14 := data.GetCall(14)
	c15 := data.GetCall(15)
	c16 := data.GetCall(16)
	c17 := data.GetCall(17)
	c18 := data.GetCall(18)
	return [][]*models.Call{
		{c17, c17, c10, c7, c10, c3, c7, c3},
		{c15, c15, c11, c6, c1, c11, c6, c1},
		{c4, c14, c9, c9, c4, c14},
		{c12, c12, c8, c8},
		{c18, c5, c5, c18, c2, c2},
		{c16, c16, c13, c13},
	}
}

// TestData returns a test data set from file
func TestData() models.INF273Data {
	return LoadDataFile(datafiles.Call7Vehicle3)
}

// TestDataMedium returns a test data set from file
func TestDataMedium() models.INF273Data {
	return LoadDataFile(datafiles.Call18Vehicle5)
}

// TestDataBig returns a test data set from file
func TestDataBig() models.INF273Data {
	return LoadDataFile(datafiles.Call130Vehicle40)
}
