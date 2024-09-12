package data

import (
	"io"
)

// The Provider interface is used to abstract the source of binary data.
type Provider interface {
	Reader() (io.Reader, error)
}
