package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
)

type azureFlagSet struct {
	*azureFlags
}

func newAzureFlagSet(flags *azureFlags) *azureFlagSet {
	return &azureFlagSet{flags}
}

func (f *azureFlagSet) flags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.StringVar(&f.accountName, "azure-account-name",
		"", "Azure storage account name")
	flagSet.StringVar(&f.accountKeyPath, "azure-account-key-path",
		"", "Path to the file containing the access key")
	flagSet.StringVar(&f.container, "azure-container",
		"", "Azure storage container name")
	flagSet.StringVar(&f.endpoint, "azure-endpoint",
		"", "Azure storage endpoint URL")

	return flagSet
}

type azureFlags struct {
	accountName    string
	accountKeyPath string
	container      string
	endpoint       string
}

func (f *azureFlags) validate() error {
	if f.accountName != "" && f.accountKeyPath == "" {
		return errors.New("--azure-account-key-path must be set")
	}
	if f.accountKeyPath != "" && f.accountName == "" {
		return errors.New("--azure-account-name must be set")
	}
	if f.container == "" {
		return errors.New("--azure-container is required")
	}
	if f.endpoint == "" {
		f.endpoint = fmt.Sprintf("https://%s.blob.core.windows.net", f.accountName)
	}

	return nil
}
