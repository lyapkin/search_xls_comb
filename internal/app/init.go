package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/lyapkin/search_comg_xls/internal/features"
	"github.com/lyapkin/search_comg_xls/internal/state"
	"github.com/lyapkin/search_comg_xls/internal/widgets/export"
	"github.com/lyapkin/search_comg_xls/internal/widgets/input"
	"github.com/lyapkin/search_comg_xls/internal/widgets/loader"
	"github.com/lyapkin/search_comg_xls/internal/widgets/output"
	"github.com/lyapkin/search_comg_xls/internal/widgets/search"
)

var choices = [9]string{"4", "5", "6", "7", "8", "9", "10", "11", "12"}

func New(a fyne.App) fyne.Window {
	window := a.NewWindow("Поиcк комбинации")
	window.Resize(fyne.NewSize(860, 400))
	window.SetFixedSize(true)

	// init state
	s := state.State{}
	s.Status = state.Init

	// init input
	i := input.New(choices[:])

	// init actions
	search := search.New(func() {
		s.SetLoading()
		go func() {
			result, err := features.Find(i.Data(), i.Columns())
			if err != nil {
				fyne.Do(func() {
					s.SetError(err)
				})
				return
			}
			fyne.Do(func() {
				s.SetSuccess(result)
			})
		}()
	})
	exp := export.New(func() {
		s.SetLoading()
		if err := features.Export(s.Data); err != nil {
			s.SetError(err)
			return
		}
		s.SetSuccess(nil)
	})

	actions := container.New(layout.NewHBoxLayout(), search.Widget, layout.NewSpacer(), exp.Widget)

	// init output
	output := output.New(s.Data)

	// init loader
	loader := loader.New()

	// subscribe
	s.Subscribe(output.Refresh, search.Refresh, exp.Refresh, loader.Refresh)

	// init window
	top := container.New(layout.NewCustomPaddedVBoxLayout(16), i.Widget, actions)
	content := container.New(
		layout.NewBorderLayout(top, loader.Widget, nil, nil),
		top,
		output.Widget,
		loader.Widget,
	)

	window.SetContent(container.NewPadded(content))

	return window
}
