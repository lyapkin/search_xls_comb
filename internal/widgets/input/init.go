package input

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/shared/numerical"
)

type input struct {
	cols   *widget.Select
	data   *fyne.Container
	Widget fyne.CanvasObject
}

func New(options []string) *input {
	i := input{}
	d, dWidget := newDataInput()
	s, sWidget := newColumns(options, d)

	i.cols = s
	i.data = d

	i.Widget = container.New(layout.NewCustomPaddedHBoxLayout(32), sWidget, dWidget)

	return &i
}

func (i *input) Columns() string {
	return i.cols.Selected
}

func (i *input) Data() map[string]struct{} {
	result := make(map[string]struct{})
	for _, o := range i.data.Objects {
		val := o.(*numerical.NumericalEntry).Text
		if val != "" {
			result[val] = struct{}{}
		}
	}
	return result
}
