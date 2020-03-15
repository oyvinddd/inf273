package util

import (
	"fmt"

	"github.com/oyvinddd/inf273/models"
)

// --------------- MISC HELPER FUNCTIONS ---------------

// PrintSolution prints a given solution to standard output
func PrintSolution(solution [][]*models.Call) {
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
