package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oyvinddd/inf273/heuristics"

	"github.com/oyvinddd/inf273/models"

	"github.com/oyvinddd/inf273/assignment2"
	"github.com/oyvinddd/inf273/util"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ------------- RUN THE PROGRAM -------------

func main() {

	// // benchmark program exection
	defer util.NewTimer().PrintElapsed()

	// // load data file and generate outsourced solution
	data := util.LoadDataFile(util.Call130Vehicle40)
	s0 := assignment2.GenerateOutsourcedSolution(data)
	//obj0 := assignment2.CalcTotalObjective(data, s0)
	s1, obj, _, _, _ := heuristics.SA(data, s0)

	util.PrintSolution(s1)
	util.PrintFlatSolution(s1)
	fmt.Printf("OBJ: %v\n", obj)
	// var x []float64 = nil
	// var y []float64 = nil
	// // var s1 [][]*models.Call = nil
	// r := make([]int, 10)
	// o := make([][][]*models.Call, 10)
	// for i := 0; i < 10; i++ {
	// 	s1, obj, xx, yy, _ := heuristics.SA(data, s0)
	// 	r[i] = obj
	// 	o[i] = s1
	// 	x = xx
	// 	y = yy
	// }
	// printResults(r, o, obj0)

	// g1 := chart.Chart{
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			XValues: x,
	// 			YValues: y,
	// 		},
	// 	},
	// }

	// f, _ := os.Create("temperature.png")
	// defer f.Close()
	// err := g1.Render(chart.PNG, f)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func printResults(r []int, o [][][]*models.Call, obj0 int) {
	fmt.Printf("INITIAL: %v\n", obj0)
	fmt.Println("ALL RESULTS:")
	best := r[0]
	var sol [][]*models.Call = nil
	sum := 0
	for i := range r {
		if r[i] < best {
			best = r[i]
			sol = o[i]
		}
		sum += r[i]
		fmt.Println(r[i])
	}
	fmt.Printf("AVG. OBJ: %v\n", sum/10.0)
	fmt.Printf("BEST: %v\n", best)
	fmt.Printf("IMPR.: %v\n", (100.0 * (float32(obj0) - float32(best)) / float32(obj0)))
	util.PrintSolution(sol)
	util.PrintFlatSolution(sol)
}
