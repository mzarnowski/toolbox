package encoding

import (
	"io"
)

type Encoding[A any] interface {
	Encode(value *A, out io.Writer) error
	Decode(receiver *A, in io.Reader) error
}
