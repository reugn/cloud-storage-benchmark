package data

import (
	"bytes"
	"io"
	"os"
)

// FileProvider reads data from a file.
type FileProvider struct {
	path string
}

var _ Provider = (*FileProvider)(nil)

// NewFileProvider returns a new FileProvider.
func NewFileProvider(path string) *FileProvider {
	return &FileProvider{
		path: path,
	}
}

func (f *FileProvider) Reader() (io.Reader, error) {
	data, err := os.ReadFile(f.path)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(data), nil
}
