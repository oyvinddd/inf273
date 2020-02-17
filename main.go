package main

import (
	"fmt"
	"time"

	"github.com/oyvinddd/inf273/util"
)

func main() {

	start := time.Now()

	util.ParseFile("Call_7_Vehicle_3.txt")

	elapsed := time.Now().Sub(start)

	fmt.Printf("Finished execution in %v", elapsed)
}
