package output

import (
	"fmt"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/state"
)

var noData = "Данных нет"

var headers = []string{"Дата", "Сумма", "Чет", "Нечет"}

type output struct {
	table  *widget.Table
	text   *widget.Label
	Widget *fyne.Container
	data   [][]string
}

func New(data [][]string) *output {
	o := output{}

	o.data = data

	o.text = widget.NewLabel(noData)

	o.Widget = container.New(layout.NewBorderLayout(nil, nil, nil, nil), o.text)

	return &o
}

func (o *output) showData() {
	o.Widget.RemoveAll()
	o.Widget.Add(o.newTable())
	o.Widget.Refresh()
}

func (o *output) showText(text string) {
	o.text.Text = text
	o.Widget.RemoveAll()
	o.Widget.Add(o.text)
	o.Widget.Refresh()
}

func (o *output) Refresh(s *state.State) {
	switch s.Status {
	case state.Ready:
		if unsafe.SliceData(s.Data) == unsafe.SliceData(o.data) {
			return
		}

		o.data = s.Data
		if len(o.data) > 0 {
			o.showData()
			return
		}

		o.showText(noData)

	case state.Error:
		o.showText(s.Error.Error())
	}
}

func (o *output) newTable() *widget.Table {
	t := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			if o.data == nil {
				return 0, 0
			}

			rows = len(o.data)
			cols = 0
			if rows > 0 {
				cols = len(o.data[0])
			}

			return rows, cols
		},
		func() fyne.CanvasObject {
			l := widget.NewLabel("")
			l.Alignment = fyne.TextAlignCenter
			return l
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(o.data[i.Row][i.Col])
		},
	)

	t.SetColumnWidth(0, 100)
	dataLength := len(o.data[0]) - 4
	for i := 1; i <= dataLength; i++ {
		t.SetColumnWidth(i, 32)
	}
	for i := dataLength + 1; i <= dataLength+3; i++ {
		t.SetColumnWidth(i, 80)
	}

	t.CreateHeader = func() fyne.CanvasObject {
		l := widget.NewLabel("")
		l.Alignment = fyne.TextAlignCenter

		return l
	}

	t.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row == -1 && id.Col == 0 {
			template.(*widget.Label).SetText(headers[id.Col])
		} else if id.Row == -1 && id.Col > dataLength {
			template.(*widget.Label).SetText(headers[id.Col-dataLength])
		} else if id.Row == -1 {
			template.(*widget.Label).SetText(fmt.Sprint(id.Col))
		} else {
			template.(*widget.Label).SetText(fmt.Sprintf("%d", id.Row+1))
		}
	}

	return t
}
