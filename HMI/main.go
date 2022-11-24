package main

import (
	g "github.com/AllenDang/giu"
	"github.com/AllenDang/imgui-go"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
	"github.com/sammyoina/stewart-platform-ui/ui"
)

func main() {
	for x := 0.0; x < 10; x += 0.1 {
		pipeline.IntakeVelocity = append(pipeline.IntakeVelocity, 0)
		pipeline.DiffuserVelocity = append(pipeline.DiffuserVelocity, 0)
		pipeline.TestSectionVelocity = append(pipeline.TestSectionVelocity, 0)
		pipeline.Fx = append(pipeline.Fx, 0)
		pipeline.Fy = append(pipeline.Fy, 0)
		pipeline.Fz = append(pipeline.Fz, 0)
		pipeline.Mx = append(pipeline.Mx, 0)
		pipeline.My = append(pipeline.My, 0)
		pipeline.Mz = append(pipeline.Mz, 0)
		pipeline.Strain1 = append(pipeline.Strain1, 0)
		pipeline.Strain2 = append(pipeline.Strain2, 0)
		pipeline.Strain3 = append(pipeline.Strain3, 0)
		pipeline.Strain4 = append(pipeline.Strain4, 0)
		pipeline.Strain5 = append(pipeline.Strain5, 0)
		pipeline.Strain6 = append(pipeline.Strain6, 0)
		pipeline.Yaw = append(pipeline.Yaw, 0)
		pipeline.Roll = append(pipeline.Roll, 0)
		pipeline.Pitch = append(pipeline.Pitch, 0)
	}
	w := g.NewMasterWindow("Overview", 1300, 700, 0)
	//w.SetBgColor(color.White)
	imgui.StyleColorsLight()
	go ui.UpdateCharts()
	go pipeline.InitPipeline()
	w.Run(ui.Loop)
}
