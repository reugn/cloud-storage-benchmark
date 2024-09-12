//nolint:revive
package azure

import (
	"math/rand"
	"time"

	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
)

type BlobClient struct {
	generator data.Provider
}

func NewBlobClient(generator data.Provider) *BlobClient {
	return &BlobClient{
		generator: generator,
	}
}

var _ storage.Client = (*BlobClient)(nil)

func (c *BlobClient) Read(key string) error {
	n := time.Duration(rand.Intn(1000))
	time.Sleep(n * time.Millisecond)
	return nil
}

func (c *BlobClient) Write(key string) error {
	n := time.Duration(rand.Intn(1000))
	time.Sleep(n * time.Millisecond)
	return nil
}

func (c *BlobClient) Delete(key string) error {
	return nil
}
