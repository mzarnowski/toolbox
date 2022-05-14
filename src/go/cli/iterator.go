package cli

type Iterator[A any] interface {
	Next() (A, error)
}
