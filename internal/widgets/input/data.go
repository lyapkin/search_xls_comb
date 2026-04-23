package input

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func newDataInput() (data *fyne.Container, widget *fyne.Container) {
	data = container.New(layout.NewHBoxLayout())
	widget = container.New(layout.NewVBoxLayout(), data)

	return data, widget
}
