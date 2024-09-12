package data

import (
	"bytes"
	"crypto/rand"
	"io"
)

// RandomGenerator generates random data.
type RandomGenerator struct {
	bytes int64
}

var _ Provider = (*RandomGenerator)(nil)

// NewRandomGenerator returns a new RandomGenerator.
func NewRandomGenerator(bytes int64) *RandomGenerator {
	return &RandomGenerator{
		bytes: bytes,
	}
}

func (g *RandomGenerator) Reader() (io.Reader, error) {
	data := make([]byte, g.bytes)
	if _, err := rand.Read(data); err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
