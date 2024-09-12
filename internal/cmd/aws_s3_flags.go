package cmd

import (
	"errors"

	"github.com/spf13/pflag"
)

type awsFlagSet struct {
	*awsFlags
}

func newAwsFlagSet(flags *awsFlags) *awsFlagSet {
	return &awsFlagSet{flags}
}

func (f *awsFlagSet) flags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.StringVar(&f.region, "aws-region",
		"us-east-1", "AWS region")
	flagSet.StringVar(&f.profile, "aws-profile",
		"minio", "AWS profile")
	flagSet.StringVar(&f.bucket, "aws-bucket",
		"", "AWS bucket")
	flagSet.StringVar(&f.baseEndpoint, "aws-base-endpoint",
		"", "AWS base endpoint")
	flagSet.BoolVar(&f.pathStyle, "aws-path-style",
		false, "AWS use path style")

	return flagSet
}

type awsFlags struct {
	region       string
	profile      string
	bucket       string
	baseEndpoint string
	pathStyle    bool
}

func (f *awsFlags) validate() error {
	if f.region == "" {
		return errors.New("--aws-region is required")
	}
	if f.bucket == "" {
		return errors.New("--aws-bucket is required")
	}

	return nil
}
