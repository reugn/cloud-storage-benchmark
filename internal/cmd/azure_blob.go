package cmd

import (
	"fmt"

	"github.com/reugn/cloud-storage-benchmark/internal/storage/azure"
	"github.com/spf13/cobra"
)

type AzureBlob struct {
	*commonFlags
}

func NewAzureBlob(flags *commonFlags) *AzureBlob {
	return &AzureBlob{
		commonFlags: flags,
	}
}

func (c *AzureBlob) NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "azure-blob",
		Short: "Benchmark Azure Blob I/O operations",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := c.commonFlags.validate(); err != nil {
				return err
			}

			client := azure.NewBlobClient(c.commonFlags.dataProvider())
			output, err := executeCommand("Azure Blob", client, c.commonFlags)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}
}
