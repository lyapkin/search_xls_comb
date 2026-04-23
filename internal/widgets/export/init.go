package export

import (
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/state"
)

type export struct {
	Widget *widget.Button
}

func New(f func()) *export {
	s := export{}
	s.Widget = widget.NewButton("Экспортировать", f)

	s.Widget.Disable()

	return &s
}

func (e *export) Refresh(s *state.State) {
	switch s.Status {
	case state.Ready:
		if len(s.Data) != 0 {
			e.Widget.Enable()
		}
	case state.Loading, state.Error:
		e.Widget.Disable()
	}
}
