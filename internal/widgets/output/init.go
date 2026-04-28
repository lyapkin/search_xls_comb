package output

import (
	"fmt"
	"strconv"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lyapkin/search_comg_xls/internal/state"
)

var noData = "Данных нет"

var headers = []string{"№", "Дата", "Сумма", "Чет", "Нечет"}

type output struct {
	table  *widget.Table
	text   *widget.Label
	Widget *fyne.Container
	data   []state.Row
}

func New(data []state.Row) *output {
	o := output{}

	o.data = data

	o.text = widget.NewLabel(noData)

	o.Widget = container.New(layout.NewBorderLayout(nil, nil, nil, nil), o.text)

	return &o
}

func (o *output) showData(cols int) {
	o.Widget.RemoveAll()
	o.Widget.Add(o.newTable(cols))
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
			o.showData(len(o.data[0].Data))
			return
		}

		o.showText(noData)

	case state.Error:
		o.showText(s.Error.Error())
	}
}

func (o *output) newTable(dataCols int) *widget.Table {
	t := widget.NewTableWithHeaders(
		func() (rows int, cols int) {
			if o.data == nil {
				return 0, 0
			}

			rows = len(o.data)
			cols = dataCols + len(headers)

			return rows, cols
		},
		func() fyne.CanvasObject {
			l := widget.NewLabel("")
			l.Alignment = fyne.TextAlignCenter
			return l
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			switch i.Col {
			case 0:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Number))
			case 1:
				obj.(*widget.Label).SetText(o.data[i.Row].Date.Format("02.01.06"))
			case dataCols + 2:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Sum))
			case dataCols + 3:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Even))
			case dataCols + 4:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Odd))
			default:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Data[i.Col-2]))
			}
		},
	)

	// set column length for № cell
	t.SetColumnWidth(0, 24)
	// set column length for date cell
	t.SetColumnWidth(1, 100)
	// set column length for data cells
	dataLength := len(o.data[0].Data)
	for i := 2; i <= dataLength+2; i++ {
		t.SetColumnWidth(i, 32)
	}
	// set column length for the rest cells
	for i := dataLength + 2; i <= dataLength+4; i++ {
		t.SetColumnWidth(i, 80)
	}

	t.CreateHeader = func() fyne.CanvasObject {
		l := widget.NewLabel("")
		l.Alignment = fyne.TextAlignCenter

		return l
	}

	t.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row == -1 && (id.Col == 0 || id.Col == 1) {
			// set header for first 2 columns
			template.(*widget.Label).SetText(headers[id.Col])
		} else if id.Row == -1 && id.Col > dataLength+1 {
			// set header for last 3 columns
			template.(*widget.Label).SetText(headers[id.Col-dataLength])
		} else if id.Row == -1 {
			// set header for the data columns
			template.(*widget.Label).SetText(fmt.Sprint(id.Col - 1))
		} else {
			// set row count
			// template.(*widget.Label).SetText(fmt.Sprintf("%d", id.Row+1))
		}
	}

	return t
}
