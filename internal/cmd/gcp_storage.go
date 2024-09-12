package cmd

import (
	"fmt"

	"github.com/reugn/cloud-storage-benchmark/internal/storage/gcp"
	"github.com/spf13/cobra"
)

type GcpStorage struct {
	*commonFlags
}

func NewGcpStorage(flags *commonFlags) *GcpStorage {
	return &GcpStorage{
		commonFlags: flags,
	}
}

func (c *GcpStorage) NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gcp-storage",
		Short: "Benchmark GCP Storage I/O operations",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := c.commonFlags.validate(); err != nil {
				return err
			}

			client := gcp.NewStorageClient(c.commonFlags.dataProvider())
			output, err := executeCommand("GCP Storage", client, c.commonFlags)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}
}
