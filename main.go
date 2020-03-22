package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// load data from file
	//data := util.LoadDataFile(util.Call7Vehicle3)

	// benchmark program exection
	defer util.NewTimer().PrintElapsed()

	s1 := util.FeasibleTestSolution()
	s2 := util.CopySolution(s1)
	s3 := util.DeepCopy(s1)

	fmt.Printf("%p\n", s1[1][0])
	fmt.Printf("%p\n", s1[1][2])
	println()
	fmt.Printf("%p\n", s2[1][0])
	fmt.Printf("%p\n", s2[1][2])
	println()
	fmt.Printf("%p\n", s3[1][0])
	fmt.Printf("%p\n", s3[1][2])

	fmt.Println(s1[0][0] == s2[0][1])

	for i := 0; i < 10000; i++ {
		//util.DeepCopy(s1)
		//util.CopySolution(s1)
	}
}
