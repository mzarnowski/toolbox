package cli

import (
	"github.com/rivo/tview"
)

func NewFoo(header []string, rows int) *Foo {
	table := tview.NewTable()
	return &Foo{
		Table:  table,
		Header: header,
		Cells:  createCells(table, rows, len(header)),
	}
}

func createCells(table *tview.Table, rows, cols int) [][]*tview.TableCell {
	cells := make([][]*tview.TableCell, rows)
	for row := 0; row < rows; row++ {
		cells[row] = make([]*tview.TableCell, cols)
		for col := 0; col < cols; col++ {
			cells[row][col] = tview.NewTableCell("")
			table.SetCell(row, col, cells[row][col])
		}
	}
	return cells
}

type Foo struct {
	Table  *tview.Table
	Header []string
	Cells  [][]*tview.TableCell
}

func (f Foo) Columns() int {
	return len(f.Header)
}

// TODO need an intermediate structure...
func (f *Foo) Insert(row int, values []string) func() {
	return func() {
		for i, col := range values {
			f.Cells[row][i].SetText(col)
		}
	}
}
