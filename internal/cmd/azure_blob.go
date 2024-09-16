package cmd

import (
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/reugn/cloud-storage-benchmark/internal/storage/azure"
	"github.com/spf13/cobra"
)

type AzureBlob struct {
	*commonFlags
	*azureFlags
}

func NewAzureBlob(flags *commonFlags) *AzureBlob {
	return &AzureBlob{
		commonFlags: flags,
		azureFlags:  &azureFlags{},
	}
}

func (c *AzureBlob) NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "azure-blob",
		Short: "Benchmark Azure Blob I/O operations",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := c.commonFlags.validate(); err != nil {
				return err
			}

			if err := c.azureFlags.validate(); err != nil {
				return err
			}

			blobClient, err := c.newBlobClient()
			if err != nil {
				return fmt.Errorf("failed to create Azure Blob client: %w", err)
			}

			client := azure.NewBlobClient(
				c.commonFlags.dataProvider(),
				blobClient,
				c.container,
			)

			output, err := executeCommand("Azure Blob", client, c.commonFlags)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}

	flagSet := newAzureFlagSet(c.azureFlags)
	command.PersistentFlags().AddFlagSet(flagSet.flags())

	return command
}

func (c *AzureBlob) newBlobClient() (*azblob.Client, error) {
	var (
		blobClient *azblob.Client
		err        error
	)

	if c.accountName != "" && c.accountKeyPath != "" {
		var accountKey []byte
		accountKey, err = os.ReadFile(c.accountKeyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read account key file: %w", err)
		}
		var cred *azblob.SharedKeyCredential
		cred, err = azblob.NewSharedKeyCredential(c.accountName, string(accountKey))
		if err != nil {
			return nil, fmt.Errorf("failed to create shared key credentials: %w", err)
		}

		blobClient, err = azblob.NewClientWithSharedKeyCredential(c.endpoint, cred, nil)
	} else {
		blobClient, err = azblob.NewClientWithNoCredential(c.endpoint, nil)
	}

	if err != nil {
		return nil, err
	}

	return blobClient, nil
}
