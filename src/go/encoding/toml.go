package encoding

import (
	"io"

	"github.com/BurntSushi/toml"
)

type tomlEncoding[A any] struct{}

func (t *tomlEncoding[A]) Encode(value *A, out io.Writer) error {
	return toml.NewEncoder(out).Encode(value)
}

func (t *tomlEncoding[A]) Decode(receiver *A, in io.Reader) error {
	_, err := toml.NewDecoder(in).Decode(receiver)
	return err
}

func TOML[A any]() Encoding[A] {
	return &tomlEncoding[A]{}
}

func NewTomlFile[A any](path string) Encoded[A] {
	return NewFile(path, TOML[A]())
}
