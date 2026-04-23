package search

import (
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/state"
)

type search struct {
	Widget *widget.Button
}

func New(f func()) *search {
	s := search{}
	s.Widget = widget.NewButton("Найти", f)

	return &s
}

func (search *search) Refresh(s *state.State) {
	switch s.Status {
	case state.Ready, state.Error:
		search.Widget.Enable()
	case state.Loading:
		search.Widget.Disable()
	}
}
