package cmd

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"github.com/reugn/cloud-storage-benchmark/internal/storage/gcp"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

type GcpStorage struct {
	*commonFlags
	*gcpFlags
}

func NewGcpStorage(flags *commonFlags) *GcpStorage {
	return &GcpStorage{
		commonFlags: flags,
		gcpFlags:    &gcpFlags{},
	}
}

//nolint:dupl
func (c *GcpStorage) NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "gcp-storage",
		Short: "Benchmark GCP Storage I/O operations",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := c.commonFlags.validate(); err != nil {
				return err
			}

			if err := c.gcpFlags.validate(); err != nil {
				return err
			}

			storageClient, err := c.newStorageClient(context.Background())
			if err != nil {
				return fmt.Errorf("failed to create GCP Storage client: %w", err)
			}

			client := gcp.NewStorageClient(
				c.commonFlags.dataProvider(),
				storageClient,
				c.bucket,
			)

			output, err := executeCommand("GCP Storage", client, c.commonFlags)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}

	flagSet := newGcpFlagSet(c.gcpFlags)
	command.PersistentFlags().AddFlagSet(flagSet.flags())

	return command
}

func (c *GcpStorage) newStorageClient(ctx context.Context) (*storage.Client, error) {
	var opts []option.ClientOption

	if c.keyPath != "" {
		opts = append(opts, option.WithCredentialsFile(c.keyPath))
	}

	if c.endpoint != "" {
		opts = append(opts, option.WithEndpoint(c.endpoint), option.WithoutAuthentication())
	}

	gcpClient, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return gcpClient, nil
}
