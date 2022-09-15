package main

import (
	"math"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
	"github.com/sammyoina/stewart-platform-ui/ui"
)

func main() {
	for x := 0.0; x < 10; x += 0.1 {
		ui.Leg1 = append(ui.Leg1, math.Sin(x))
		ui.Leg2 = append(ui.Leg2, math.Cos(x))
		ui.Leg3 = append(ui.Leg3, math.Sin(x)+0.1)
		ui.Leg4 = append(ui.Leg4, math.Cos(x)+0.1)
		ui.Leg5 = append(ui.Leg5, math.Sin(x)+0.3)
		ui.Leg6 = append(ui.Leg6, math.Cos(x)+0.3)
	}
	for x := 0.0; x < 10; x += 0.1 {
		ui.AirVelocity = append(ui.AirVelocity, 1.5)
	}
	w := g.NewMasterWindow("Overview", 1300, 700, 0)
	//w.SetBgColor(color.White)
	imgui.StyleColorsLight()
	go ui.Rotate()
	go pipeline.InitPipeline()
	w.Run(ui.Loop)
}
