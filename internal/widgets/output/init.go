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

var headers = []string{"№", "Дата", "Сумма", "Чет", "Нечет", "Кол-во"}

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
	amount := widget.NewLabel(fmt.Sprintf("Количество найденных строк: %d", len(o.data)))
	o.Widget.RemoveAll()
	o.Widget.Layout = layout.NewBorderLayout(amount, nil, nil, nil)
	o.Widget.Add(amount)
	o.Widget.Add(o.newTable(cols))
	o.Widget.Refresh()
}

func (o *output) showText(text string) {
	o.text.Text = text
	o.Widget.RemoveAll()
	o.Widget.Layout = layout.NewBorderLayout(nil, nil, nil, nil)
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
			l.Truncation = fyne.TextTruncateEllipsis
			return l
		},
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			switch i.Col {
			case 0:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Number))
			case 1:
				obj.(*widget.Label).SetText(o.data[i.Row].Date)
			case dataCols + 2:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Sum))
			case dataCols + 3:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Even))
			case dataCols + 4:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Odd))
			case dataCols + 5:
				val := strconv.Itoa(o.data[i.Row].Count)
				if val == "0" {
					val = "-"
				}
				obj.(*widget.Label).SetText(val)
			default:
				obj.(*widget.Label).SetText(strconv.Itoa(o.data[i.Row].Data[i.Col-2]))
			}
		},
	)

	// set column length for № cell
	t.SetColumnWidth(0, 48)
	// set column length for date cell
	t.SetColumnWidth(1, 100)
	// set column length for data cells
	dataLength := len(o.data[0].Data)
	for i := 2; i <= dataLength+2; i++ {
		t.SetColumnWidth(i, 36)
	}
	// set column length for the rest cells
	for i := dataLength + 2; i <= dataLength+5; i++ {
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
