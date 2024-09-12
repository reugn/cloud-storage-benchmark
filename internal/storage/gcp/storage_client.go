//nolint:revive
package gcp

import (
	"math/rand"
	"time"

	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
)

type StorageClient struct {
	generator data.Provider
}

func NewStorageClient(generator data.Provider) *StorageClient {
	return &StorageClient{
		generator: generator,
	}
}

var _ storage.Client = (*StorageClient)(nil)

func (c *StorageClient) Read(key string) error {
	n := time.Duration(rand.Intn(1000))
	time.Sleep(n * time.Millisecond)
	return nil
}

func (c *StorageClient) Write(key string) error {
	n := time.Duration(rand.Intn(1000))
	time.Sleep(n * time.Millisecond)
	return nil
}

func (c *StorageClient) Delete(key string) error {
	return nil
}
