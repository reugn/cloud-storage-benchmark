package aws

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
)

type S3Client struct {
	dataProvider data.Provider
	client       *s3.Client
	bucket       string
}

func NewS3Client(dataProvider data.Provider, client *s3.Client,
	bucket string) *S3Client {
	return &S3Client{
		dataProvider: dataProvider,
		client:       client,
		bucket:       bucket,
	}
}

var _ storage.Client = (*S3Client)(nil)

func (c *S3Client) Read(key string) error {
	getObjectOutput, err := c.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &c.bucket,
		Key:    &key,
	})
	if err != nil {
		return fmt.Errorf("failed to get object %s: %w", key, err)
	}

	_, err = io.ReadAll(getObjectOutput.Body)

	return errors.Join(err, getObjectOutput.Body.Close())
}

func (c *S3Client) Write(key string) error {
	dataReader, err := c.dataProvider.Reader()
	if err != nil {
		return err
	}

	if _, err = c.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &c.bucket,
		Key:    &key,
		Body:   dataReader,
	}); err != nil {
		return fmt.Errorf("failed to put object %s: %w", key, err)
	}

	return nil
}

func (c *S3Client) Delete(key string) error {
	if _, err := c.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: &c.bucket,
		Key:    &key,
	}); err != nil {
		return fmt.Errorf("failed to delete object %s: %w", key, err)
	}

	return nil
}
