package azure

import (
	"context"
	"errors"
	"io"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
)

type BlobClient struct {
	dataProvider data.Provider
	client       *azblob.Client
	container    string
}

func NewBlobClient(dataProvider data.Provider,
	client *azblob.Client, container string) *BlobClient {
	return &BlobClient{
		dataProvider: dataProvider,
		client:       client,
		container:    container,
	}
}

var _ storage.Client = (*BlobClient)(nil)

func (c *BlobClient) Read(key string) error {
	streamResponse, err := c.client.DownloadStream(context.Background(), c.container, key, nil)
	if err != nil {
		return err
	}

	_, err = io.ReadAll(streamResponse.Body)

	return errors.Join(err, streamResponse.Body.Close())
}

func (c *BlobClient) Write(key string) error {
	dataReader, err := c.dataProvider.Reader()
	if err != nil {
		return err
	}

	_, err = c.client.UploadStream(context.Background(), c.container, key, dataReader, nil)

	return err
}

func (c *BlobClient) Delete(key string) error {
	_, err := c.client.DeleteBlob(context.Background(), c.container, key, nil)
	return err
}
