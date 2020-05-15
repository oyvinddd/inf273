package util

import (
	"fmt"

	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/models"
)

// --------------- MISC HELPER FUNCTIONS ---------------

// PrintSolutionAndObj prints both the solution and the objective
func PrintSolutionAndObj(s [][]*models.Call, obj int) {
	PrintFlatSolution(s) //PrintSolution(s)
	fmt.Println()
	fmt.Printf("Objective: %v\n", obj)
}

// PrintSolution prints a given solution to standard output
func PrintSolution(solution [][]*models.Call) {
	fmt.Println()

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
	fmt.Println("ALL RESULTS:")
	best := r[0]
	var s1 [][]*models.Call = s[0]
	sum := 0
	for i := range r {
		if r[i] < best {
			best = r[i]
			s1 = s[i]
		}
		sum += r[i]
		fmt.Println(r[i])
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
	// 18 6 1 6 1 18 0 15 5 15 5 14 9 9 14 0 11 17 17 11 10 10 3 3 0 12 12 0 8 7 8 7 2 2 0 13 13 4 4 16 16
	// Obj.: 3278316
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
		{c18, c6, c1, c6, c1, c18},
		{c15, c5, c15, c5, c14, c9, c9, c14},
		{c11, c17, c17, c11, c10, c10, c3, c3},
		{c12, c12},
		{c8, c7, c8, c7, c2, c2},
		{c13, c13, c4, c4, c16, c16},
	}
}

// FeasibleTestSolution5 returns a feasible (but not optimal) solution (used for testing only)
func FeasibleTestSolution5() [][]*models.Call {
	data := TestDataBig()
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
	c19 := data.GetCall(19)
	c20 := data.GetCall(20)
	c21 := data.GetCall(21)
	c22 := data.GetCall(22)
	c23 := data.GetCall(23)
	c24 := data.GetCall(24)
	c25 := data.GetCall(25)
	c26 := data.GetCall(26)
	c27 := data.GetCall(27)
	c28 := data.GetCall(28)
	c29 := data.GetCall(29)
	c30 := data.GetCall(30)
	c31 := data.GetCall(31)
	c32 := data.GetCall(32)
	c33 := data.GetCall(33)
	c34 := data.GetCall(34)
	c35 := data.GetCall(35)
	c36 := data.GetCall(36)
	c37 := data.GetCall(37)
	c38 := data.GetCall(38)
	c39 := data.GetCall(39)
	c40 := data.GetCall(40)
	c41 := data.GetCall(41)
	c42 := data.GetCall(42)
	c43 := data.GetCall(43)
	c44 := data.GetCall(44)
	c45 := data.GetCall(45)
	c46 := data.GetCall(46)
	c47 := data.GetCall(47)
	c48 := data.GetCall(48)
	c49 := data.GetCall(49)
	c50 := data.GetCall(50)
	c51 := data.GetCall(51)
	c52 := data.GetCall(52)
	c53 := data.GetCall(53)
	c54 := data.GetCall(54)
	c55 := data.GetCall(55)
	c56 := data.GetCall(56)
	c57 := data.GetCall(57)
	c58 := data.GetCall(58)
	c59 := data.GetCall(59)
	c60 := data.GetCall(60)
	c61 := data.GetCall(61)
	c62 := data.GetCall(62)
	c63 := data.GetCall(63)
	c64 := data.GetCall(64)
	c65 := data.GetCall(65)
	c66 := data.GetCall(66)
	c67 := data.GetCall(67)
	c68 := data.GetCall(68)
	c69 := data.GetCall(69)
	c70 := data.GetCall(70)
	c71 := data.GetCall(71)
	c72 := data.GetCall(72)
	c73 := data.GetCall(73)
	c74 := data.GetCall(74)
	c75 := data.GetCall(75)
	c76 := data.GetCall(76)
	c77 := data.GetCall(77)
	c78 := data.GetCall(78)
	c79 := data.GetCall(79)
	c80 := data.GetCall(80)
	return [][]*models.Call{
		{c72, c60, c72, c42, c60, c40, c40, c42},
		{c57, c37, c13, c13, c57, c37},
		{c66, c36, c66, c36, c19, c8, c8, c19},
		{c33, c33, c14, c22, c14, c22},
		{c61, c7, c7, c61},
		{c4, c4, c54, c54, c27, c27, c64, c64, c43, c43},
		{c68, c68, c28, c28, c11, c11, c17, c17, c10, c10},
		{c59, c67, c59, c67},
		{c53, c29, c53, c29},
		{c52, c52, c31, c18, c38, c35, c35, c18, c38, c31},
		{c50, c50, c73, c73},
		{c47, c47, c51, c51, c75, c75},
		{c39, c39, c70, c9, c9, c70, c6, c76, c6, c76},
		{c69, c69, c24, c71, c71, c24},
		{c48, c55, c48, c55},
		{c74, c30, c74, c30, c65, c65},
		{c3, c3, c5, c56, c56, c5},
		{c63, c63, c23, c23, c45, c1, c45, c1},
		{c32, c12, c32, c21, c12, c41, c21, c25, c41, c25},
		{c20, c20, c77, c77, c44, c44, c15, c15},
		{c62, c62, c49, c49, c2, c34, c34, c2, c79, c16, c16, c80, c80, c46, c46, c26, c26, c79, c78, c78, c58, c58},
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
	return LoadDataFile(datafiles.Call80Vehicle20)
}
