package main

import (
	"fmt"
	"math"

	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sammyoina/stewart-platform-ui/kinematics"
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
	//go pipeline.InitPipeline()
	plat := kinematics.NewStewartPlatform(11, 9.4, 0.226893, 0.226893, 2.8, 21, 0)
	pos := plat.Calculate(0, 0.0523599, 0, 0, 0, 0)
	fmt.Println(pos.Servo1, pos.Servo2, pos.Servo3, pos.Servo4, pos.Servo5, pos.Servo6)
	w.Run(ui.Loop)
}
