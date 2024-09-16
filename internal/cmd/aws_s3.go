package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	storage "github.com/reugn/cloud-storage-benchmark/internal/storage/aws"
	"github.com/spf13/cobra"
)

type AwsS3 struct {
	*commonFlags
	*awsFlags
}

func NewAwsS3(flags *commonFlags) *AwsS3 {
	return &AwsS3{
		commonFlags: flags,
		awsFlags:    &awsFlags{},
	}
}

//nolint:dupl
func (c *AwsS3) NewCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "aws-s3",
		Short: "Benchmark AWS S3 I/O operations",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := c.commonFlags.validate(); err != nil {
				return err
			}

			if err := c.awsFlags.validate(); err != nil {
				return err
			}

			s3Client, err := c.newS3Client(context.Background())
			if err != nil {
				return fmt.Errorf("failed to create S3 client: %w", err)
			}

			client := storage.NewS3Client(
				c.commonFlags.dataProvider(),
				s3Client,
				c.bucket,
			)

			output, err := executeCommand("AWS S3", client, c.commonFlags)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}

	flagSet := newAwsFlagSet(c.awsFlags)
	command.PersistentFlags().AddFlagSet(flagSet.flags())

	return command
}

func (c *AwsS3) newS3Client(ctx context.Context) (*s3.Client, error) {
	var opts []func(*config.LoadOptions) error
	opts = append(opts, config.WithDefaultRegion(c.region))
	if c.profile != "" {
		opts = append(opts, config.WithSharedConfigProfile(c.profile))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("unable to load aws config: %w", err)
	}

	var clientOpts []func(*s3.Options)
	if c.baseEndpoint != "" {
		clientOpts = append(clientOpts, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(c.baseEndpoint)
		})
	}
	if c.pathStyle {
		clientOpts = append(clientOpts, func(o *s3.Options) {
			o.UsePathStyle = c.pathStyle
		})
	}
	client := s3.NewFromConfig(cfg, clientOpts...)

	return client, nil
}
