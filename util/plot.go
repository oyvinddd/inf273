package util

import (
	"log"
	"os"

	"github.com/wcharczuk/go-chart"
)

// GenerateGraph from the given x and y coordinates
func GenerateGraph(x []float64, y []float64) {
	g1 := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}
	f, _ := os.Create("temperature.png")
	defer f.Close()
	err := g1.Render(chart.PNG, f)
	if err != nil {
		log.Fatal(err)
	}
}
