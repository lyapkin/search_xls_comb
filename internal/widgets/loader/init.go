package loader

import (
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/state"
)

type loader struct {
	Widget *widget.ProgressBarInfinite
}

func New() *loader {
	s := loader{}
	s.Widget = widget.NewProgressBarInfinite()

	s.Widget.Hide()

	return &s
}

func (l *loader) Refresh(s *state.State) {
	switch s.Status {
	case state.Ready, state.Error:
		l.Widget.Stop()
		l.Widget.Hide()
	case state.Loading:
		l.Widget.Show()
		l.Widget.Start()
	}
	l.Widget.Refresh()
}
