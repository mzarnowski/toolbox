package cli

import (
	"errors"
	"io"
	"unicode/utf8"
)

func NewRuneIterator(source io.Reader) Iterator[rune] {
	var buffer [4096]byte
	return &runeIterator{
		source: source,
		buffer: &buffer,
		offset: 0,
		limit:  0,
	}
}

type runeIterator struct {
	source io.Reader
	buffer *[4096]byte
	offset int
	limit  int
}

func (it *runeIterator) Next() (rune, error) {
	source, buffer := it.source, it.buffer
	available := it.limit - it.offset
	if available < 4 {
		copy(buffer[0:available], buffer[it.offset:it.limit])
		it.offset = 0

		len, err := source.Read(buffer[available:])
		it.limit = available + len

		if err != nil && (err != io.EOF || it.limit == 0) {
			return utf8.RuneError, err
		}
	}

	rune, len := utf8.DecodeRune(buffer[it.offset:it.limit])
	it.offset += len

	if rune == utf8.RuneError {
		return utf8.RuneError, errors.New("malformed input")
	} else {
		return rune, nil
	}
}
