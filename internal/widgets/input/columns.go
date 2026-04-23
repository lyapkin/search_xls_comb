package input

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/shared/numerical"
)

func newColumns(options []string, valuesInput *fyne.Container) (*widget.Select, fyne.CanvasObject) {
	s := widget.NewSelect(options, func(value string) {
		n, err := strconv.Atoi(value)
		if err != nil {
			log.Panicf("can not parse choices value into a number: %v", err)
		}

		valuesInput.RemoveAll()
		for range n {
			valuesInput.Add(numerical.NewNumericalEntry())
		}
	})

	s.SetSelectedIndex(0)

	w := container.New(layout.NewVBoxLayout(), s)

	return s, w
}
