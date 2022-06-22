package encoding

import (
	"os"
)

func NewFile[A any](path string, encoding Encoding[A]) Encoded[A] {
	return &file[A]{path: path, encoding: encoding}
}

type file[A any] struct {
	path     string
	encoding Encoding[A]
}

func (e *file[A]) Overwrite(value *A) error {
	file, err := os.Open(e.path)
	if os.IsNotExist(err) {
		file, err = os.Create(e.path)
	}

	if err != nil {
		return err
	}

	return e.encoding.Encode(value, file)
}

func (e *file[A]) Read(receiver *A) error {
	if file, err := os.Open(e.path); os.IsNotExist(err) {
		return e.Overwrite(receiver)
	} else if err != nil {
		return err
	} else {
		return e.encoding.Decode(receiver, file)
	}
}
