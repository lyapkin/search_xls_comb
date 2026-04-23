package numerical

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type NumericalEntry struct {
	widget.Entry
}

func NewNumericalEntry() *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)

	return entry
}

func (e *NumericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') && len(e.Text) < 2 {
		e.Entry.TypedRune(r)
	}
}

func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.Atoi(content); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *NumericalEntry) CreateRenderer() fyne.WidgetRenderer {
	r := e.Entry.CreateRenderer()

	for _, obj := range r.Objects() {
		if t, ok := obj.(*canvas.Text); ok {
			t.Alignment = fyne.TextAlignCenter
		}
	}

	return r
}
