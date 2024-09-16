package gcp

import (
	"context"
	"errors"
	"io"

	gcpStorage "cloud.google.com/go/storage"
	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
)

type StorageClient struct {
	dataProvider data.Provider
	client       *gcpStorage.Client
	bucketHandle *gcpStorage.BucketHandle
}

func NewStorageClient(dataProvider data.Provider,
	client *gcpStorage.Client, bucket string) *StorageClient {
	return &StorageClient{
		dataProvider: dataProvider,
		client:       client,
		bucketHandle: client.Bucket(bucket),
	}
}

var _ storage.Client = (*StorageClient)(nil)

func (c *StorageClient) Read(key string) error {
	reader, err := c.bucketHandle.Object(key).NewReader(context.Background())
	if err != nil {
		return err
	}

	_, err = io.ReadAll(reader)

	return errors.Join(err, reader.Close())
}

func (c *StorageClient) Write(key string) error {
	dataReader, err := c.dataProvider.Reader()
	if err != nil {
		return err
	}

	writer := c.bucketHandle.Object(key).NewWriter(context.Background())
	_, err = io.Copy(writer, dataReader)

	return errors.Join(err, writer.Close())
}

func (c *StorageClient) Delete(key string) error {
	return c.bucketHandle.Object(key).Delete(context.Background())
}
