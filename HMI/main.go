package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Force Balance")
	w.Resize(fyne.NewSize(400, 400))

	dashHome := widget.NewLabel("Stewart Platform Force-Balance Dashboard!")
	fYaw := 0.0
	fRoll := 0.0
	fPitch := 0.0

	yawLabel, yawSlider := getNewSliderWithLabel(fYaw, -5, 5)
	rollLabel, rollSlider := getNewSliderWithLabel(fRoll, -5, 5)
	pitchLabel, pitchSlider := getNewSliderWithLabel(fPitch, -5, 5)

	/*data := binding.BindFloat(&f)
	yprSilder := widget.NewSliderWithData(-5, 5, data)
	label1 := widget.NewLabelWithData(
		binding.FloatToString(data),
	)*/
	//yprSilder.Orientation = 1
	//var sliderHeight float32 = 200
	//yprSilder.Resize(fyne.Size{Width: 2000})
	w.SetContent(container.NewVBox(
		dashHome,
		widget.NewLabel("Yaw"),
		yawLabel,
		yawSlider,
		widget.NewLabel("Pitch"),
		pitchLabel,
		pitchSlider,
		widget.NewLabel("Roll"),
		rollLabel,
		rollSlider,
	))

	w.ShowAndRun()
}

func getNewSliderWithLabel(num float64, min float64, max float64) (*widget.Label, *widget.Slider) {
	data := binding.BindFloat(&num)
	yprSilder := widget.NewSliderWithData(min, max, data)
	label := widget.NewLabelWithData(
		binding.FloatToString(data),
	)
	return label, yprSilder
}
