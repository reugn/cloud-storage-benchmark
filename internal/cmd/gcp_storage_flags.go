package cmd

import (
	"errors"

	"github.com/spf13/pflag"
)

type gcpFlagSet struct {
	*gcpFlags
}

func newGcpFlagSet(flags *gcpFlags) *gcpFlagSet {
	return &gcpFlagSet{flags}
}

func (f *gcpFlags) flags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.StringVar(&f.keyPath, "gcp-key-path",
		"", "Path to the file containing service account JSON key")
	flagSet.StringVar(&f.bucket, "gcp-bucket",
		"", "GCP Storage bucket")
	flagSet.StringVar(&f.endpoint, "gcp-endpoint-override",
		"", "GCP Storage alternate endpoint URL")

	return flagSet
}

type gcpFlags struct {
	keyPath  string
	bucket   string
	endpoint string
}

func (f *gcpFlags) validate() error {
	if f.keyPath == "" {
		return errors.New("--gcp-key-path is required")
	}
	if f.bucket == "" {
		return errors.New("--gcp-bucket is required")
	}

	return nil
}
