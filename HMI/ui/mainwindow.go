package ui

import (
	"time"

	g "github.com/AllenDang/giu"
	"github.com/sammyoina/stewart-platform-ui/kinematics"
)

var (
	yaw         int32
	pitch       int32
	roll        int32
	transx      int32
	transy      int32
	transz      int32
	tabwidth    float32 = 350
	tabheight   float32 = 150
	Leg1        []float64
	Leg2        []float64
	Leg3        []float64
	Leg4        []float64
	Leg5        []float64
	Leg6        []float64
	lineTicks   []g.PlotTicker
	AirVelocity []float64
)

func Rotate() {
	ticker := time.NewTicker(time.Second * 1)
	var k int = 1
	for {
		Leg1 = append(Leg1[k:], Leg1[0:k]...)
		g.Update()
		<-ticker.C
	}
}

func Loop() {
	g.SingleWindowWithMenuBar().Layout(
		g.SplitLayout(g.DirectionVertical, tabheight,
			g.SplitLayout(g.DirectionHorizontal, tabwidth,
				g.Layout{
					g.Label("Rotation Angles"),
					g.Row(
						g.VSliderInt(&yaw, -20, 20).Label("Yaw").Size(40, 110),
						g.VSliderInt(&pitch, -22, 24).Label("Pitch").Size(40, 110),
						g.VSliderInt(&roll, -34, 34).Label("Roll").Size(40, 110),
					),
				},
				g.SplitLayout(g.DirectionHorizontal, tabwidth,
					g.Layout{
						g.Label("Translations (mm)"),
						g.Row(
							g.VSliderInt(&transx, -5, 5).Label("X").Size(40, 110),
							g.VSliderInt(&transy, -5, 5).Label("Y").Size(40, 110),
							g.VSliderInt(&transz, -5, 5).Label("Z").Size(40, 110),
						),
					},
					g.Layout{
						g.Row(
							g.Button("Home Platform").Size(120, 100).OnClick(func() {
								kinematics.SetOrientation(0, 0, 0, 0, 0, 0)
							}),
							g.Button("Write Position").Size(120, 100).OnClick(func() {
								kinematics.SetOrientation(float64(yaw), float64(pitch), float64(roll), float64(transx), float64(transy), float64(transz))
							}),
							g.Button("Record Data").Size(120, 100),
						),
					},
				),
			),
			g.Layout{
				g.Plot("Strain Data").AxisLimits(0, 100, -1.2, 1.2, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotLine("Leg 1", Leg1),
					g.PlotLine("Leg 2", Leg2),
					g.PlotLine("Leg 3", Leg3),
					g.PlotLine("Leg 4", Leg4),
					g.PlotLine("Leg 5", Leg5),
					g.PlotLine("Leg 6", Leg6),
				),
				g.Plot("Air velocity").AxisLimits(0, 100, -2, 2, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotScatter("velocity(m/s)", AirVelocity),
				),
			},
		),
	)
}
