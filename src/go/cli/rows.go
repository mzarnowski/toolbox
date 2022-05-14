package cli

import (
	"io"
	"strings"
)

type Row []string

type Table struct {
	Header   []string
	Iterator Iterator[Row]
}

func NewTable(header []string, iterator Iterator[rune]) Table {
	return Table{
		Header:   header,
		Iterator: NewRowIterator(len(header), iterator),
	}
}

func (r *Table) Next() (Row, error) {
	return r.Iterator.Next()
}

func (r *Table) Columns() int {
	return len(r.Header)
}

type RowIterator struct {
	Iterator[Row]

	columns int
	runes   Iterator[rune]
}

func NewRowIterator(columns int, runes Iterator[rune]) Iterator[Row] {
	return &RowIterator{columns: columns, runes: runes}
}

func (table *RowIterator) Next() (Row, error) {
	row := make([]string, table.columns)
	for i := 0; i < table.columns-1; i++ {
		if value, err := table.nextValue('\t'); err != nil {
			return nil, err
		} else {
			row[i] = value
		}
	}

	if value, err := table.nextValue('\n'); err != nil && err != io.EOF {
		return nil, err
	} else {
		row[table.columns-1] = value
	}

	return row, nil
}

func (table *RowIterator) nextValue(expected rune) (string, error) {
	var builder strings.Builder
	for {
		if rune, err := table.runes.Next(); err != nil && err != io.EOF {
			return "", err
		} else if rune == expected {
			return builder.String(), nil
		} else if err == io.EOF {
			return builder.String(), io.EOF
		} else {
			builder.WriteRune(rune)
		}
	}
}

func RowsFromReader(in *io.Reader, header []string) *Table {
	if header == nil {
		header = make([]string, 0) // TODO parse
	}

	return &Table{
		Header:   header,
		Iterator: nil,
	}
}

// func Parse(in *io.Reader, columns int) Iterator[Row] {
// 	return &RowIterator{columns: columns, source: in}
// }
