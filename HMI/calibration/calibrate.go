package calibration

import (
	"fmt"

	g "github.com/AllenDang/giu"
)

var CalibrationFactor float64
var CalibrationLoads [100]float64
var CalibrationRef float32

func init() {
	CalibrationFactor = 1
}

func CalibratePlatform() {
	var sum float64

	for _, val := range CalibrationLoads {
		sum += val
	}
	average := sum / float64(len(CalibrationLoads))
	CalibrationFactor = average / float64(CalibrationRef)
	fmt.Println("New Calibration Factor: ", CalibrationFactor)
	g.CloseCurrentPopup()
}
