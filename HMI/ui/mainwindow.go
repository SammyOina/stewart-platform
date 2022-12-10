package ui

import (
	"fmt"
	"time"

	g "github.com/AllenDang/giu"
	"github.com/sammyoina/stewart-platform-ui/calibration"
	"github.com/sammyoina/stewart-platform-ui/fileWriter"
	"github.com/sammyoina/stewart-platform-ui/pipeline"
)

var (
	Yaw       int32
	Pitch     int32
	Roll      int32
	Transx    int32
	Transy    int32
	Transz    int32
	tabwidth  float32 = 350
	tabheight float32 = 150
	lineTicks []g.PlotTicker
)

func UpdateCharts() {
	ticker := time.NewTicker(time.Second * 1)
	for {
		g.Update()
		<-ticker.C
	}
}

func Loop() {
	mainWindow()
}

func mainWindow() *g.WindowWidget {
	w := g.SingleWindowWithMenuBar()
	w.Layout(
		g.MenuBar().Layout(
			g.Menu("Run Tests").Layout(
				g.MenuItem("Roll Test").OnClick(pipeline.TestPlatformRoll),
				g.MenuItem("Yaw Test").OnClick(pipeline.TestPlatformYaw),
				g.MenuItem("Pitch Test").OnClick(pipeline.TestPlatformPitch),
				g.MenuItem("X Translation Test").OnClick(pipeline.TestPlatformXTrans),
				g.MenuItem("Y Translation Test").OnClick(pipeline.TestPlatformYTrans),
				g.MenuItem("Z Translation Test").OnClick(pipeline.TestPlatformZTrans),
			),
		),
		g.PrepareMsgbox(),
		g.SplitLayout(g.DirectionVertical, tabheight,
			g.SplitLayout(g.DirectionHorizontal, tabwidth,
				g.Layout{
					g.Label("Rotation Angles (deg)"),
					g.Row(
						g.VSliderInt(&Yaw, -20, 20).Label("Yaw").Size(40, 110),
						g.VSliderInt(&Pitch, -22, 24).Label("Pitch").Size(40, 110),
						g.VSliderInt(&Roll, -34, 34).Label("Roll").Size(40, 110),
					),
				},
				g.SplitLayout(g.DirectionHorizontal, tabwidth,
					g.Layout{
						g.Label("Translations (mm)"),
						g.Row(
							g.VSliderInt(&Transx, -5, 5).Label("X").Size(40, 110),
							g.VSliderInt(&Transy, -5, 5).Label("Y").Size(40, 110),
							g.VSliderInt(&Transz, -5, 5).Label("Z").Size(40, 110),
						),
					},
					g.Layout{
						g.Row(
							g.Button("Home Platform").Size(120, 100).OnClick(func() {
								pipeline.SetOrientation(0, 0, 0, 0, 0, 0)
							}),
							g.Button("Write Position").Size(120, 100).OnClick(func() {
								pipeline.SetOrientation(float64(Yaw), float64(Pitch), float64(Roll), float64(Transx), float64(Transy), float64(Transz))
							}),
							g.Condition(pipeline.RecordData,
								g.Layout{
									g.Button("Stop Recording").Size(120, 100).OnClick(func() {
										pipeline.RecordData = false
										if pipeline.PitotWriter != nil {
											pipeline.PitotWriter.QuitChannel <- true
										}
										pipeline.PitotWriter = nil
										if pipeline.IMUWriter != nil {
											pipeline.IMUWriter.QuitChannel <- true
										}
										pipeline.IMUWriter = nil
										if pipeline.StrainWriter != nil {
											pipeline.StrainWriter.QuitChannel <- true
										}
										pipeline.StrainWriter = nil
										if pipeline.ForceMomentsWriter != nil {
											pipeline.ForceMomentsWriter.QuitChannel <- true
										}
										pipeline.ForceMomentsWriter = nil
									}),
								}, g.Layout{
									g.Button("Record Data").Size(120, 100).OnClick(func() {
										pipeline.RecordData = true
										currTime := time.Now()
										timeS := fmt.Sprintf("%d-%d-%d_%d-%d", currTime.Day(), currTime.Month(), currTime.Year(), currTime.Hour(), currTime.Minute())
										pipeline.PitotWriter, _ = fileWriter.NewWriter(timeS + "/" + "pitot.csv")
										pipeline.IMUWriter, _ = fileWriter.NewWriter(timeS + "/" + "imu.csv")
										pipeline.StrainWriter, _ = fileWriter.NewWriter(timeS + "/" + "strain_raw.csv")
										pipeline.ForceMomentsWriter, _ = fileWriter.NewWriter(timeS + "/" + "forces-moments.csv")
										go pipeline.PitotWriter.Record()
										go pipeline.IMUWriter.Record()
										go pipeline.StrainWriter.Record()
										go pipeline.ForceMomentsWriter.Record()
									}),
								}),
							g.Button("Calibrate").OnClick(func() {
								g.OpenPopup("Calibration")
							}).Size(120, 100),
							g.PopupModal("Calibration").Layout(
								g.Label("Set Reference Weight in grams"),
								g.Row(
									g.InputFloat(&calibration.CalibrationRef).Size(200),
								),
								g.Row(
									g.Button("Set").OnClick(calibration.CalibratePlatform).Size(100, 35),
									g.Button("Close").OnClick(func() { g.CloseCurrentPopup() }).Size(100, 35),
								),
							),
						),
					},
				),
			),
			g.Layout{
				g.Plot("Aerodynamic Loads (N)").AxisLimits(0, 100, -1.2, 1.2, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotLine("Force X", pipeline.Fx),
					g.PlotLine("Force Y", pipeline.Fy),
					g.PlotLine("Force Z", pipeline.Fz),
					g.PlotLine("Moment X", pipeline.Mx),
					g.PlotLine("Moment Y", pipeline.My),
					g.PlotLine("Moment Z", pipeline.Mz),
				),
				g.Plot("Air velocity (m/s)").AxisLimits(0, 100, -2, 10, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotScatter("Intake", pipeline.IntakeVelocity),
					g.PlotScatter("Test Section", pipeline.TestSectionVelocity),
					g.PlotScatter("Diffuser", pipeline.DiffuserVelocity),
				),
				g.Plot("Raw Strain Loads (g)").AxisLimits(0, 100, -1000, 1000, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotLine("Strain 1", pipeline.Strain1),
					g.PlotLine("Strain 2", pipeline.Strain2),
					g.PlotLine("Strain 3", pipeline.Strain3),
					g.PlotLine("Strain 4", pipeline.Strain4),
					g.PlotLine("Strain 5", pipeline.Strain5),
					g.PlotLine("Strain 6", pipeline.Strain6),
				),
				g.Plot("IMU Orientation (Deg)").AxisLimits(0, 100, -50, 50, g.ConditionOnce).XTicks(lineTicks, false).Plots(
					g.PlotLine("yaw", pipeline.Yaw),
					g.PlotLine("pitch", pipeline.Pitch),
					g.PlotLine("roll", pipeline.Roll),
				),
			},
		),
	)
	return w
}
