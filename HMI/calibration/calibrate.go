package calibration

import (
	"log"

	g "github.com/AllenDang/giu"
)

var (
	//CalibrationFactor  float64
	//CalibrationLoads   []float64
	CalibrationFactors []float64
	CalibrationRef     float32
	CalibLoad1         []float64
	CalibLoad2         []float64
	CalibLoad3         []float64
	CalibLoad4         []float64
	CalibLoad5         []float64
	CalibLoad6         []float64
)

func init() {
	//CalibrationFactor = 1
	for i := 0; i < 6; i++ {
		CalibrationFactors[i] = 1
	}
}

/*func CalibratePlatform() {
	var sum float64

	for _, val := range CalibrationLoads {
		sum += val
	}
	average := sum / float64(len(CalibrationLoads))
	CalibrationFactor = average / (float64(CalibrationRef) * 0.00981)
	fmt.Println("New Calibration Factor: ", CalibrationFactor)
	g.CloseCurrentPopup()
}*/

func CalibratePlatform() {
	var sums []float64
	for _, v := range CalibLoad1 {
		sums[0] += v
	}
	for _, v := range CalibLoad2 {
		sums[1] += v
	}
	for _, v := range CalibLoad3 {
		sums[2] += v
	}
	for _, v := range CalibLoad4 {
		sums[3] += v
	}
	for _, v := range CalibLoad5 {
		sums[4] += v
	}
	for _, v := range CalibLoad6 {
		sums[5] += v
	}
	for v, sum := range sums {
		CalibrationFactors[v] = (sum / float64(len(CalibLoad1))) / float64(CalibrationRef/6)
		log.Println("Calibration factor ", v, " : ", CalibrationFactors[v])
	}
	g.CloseCurrentPopup()
}
