package encoding

type Encoded[A any] interface {
	Overwrite(value *A) error
	Read(receiver *A) error
}
