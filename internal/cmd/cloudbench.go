package cmd

import (
	"fmt"

	"github.com/reugn/cloud-storage-benchmark/internal/model"
	"github.com/reugn/cloud-storage-benchmark/internal/proc"
	"github.com/reugn/cloud-storage-benchmark/internal/storage"
	"github.com/reugn/cloud-storage-benchmark/internal/util"
	"github.com/spf13/cobra"
)

const (
	objectPrefix = "cloud-storage-benchmark/object"
)

type CloudBench struct {
}

func NewCloudBench() *CloudBench {
	return &CloudBench{}
}

func (c *CloudBench) NewCommand(version string) *cobra.Command {
	command := &cobra.Command{
		Use:     "cloudbench",
		Short:   "Cloud Storage Benchmark",
		Version: version,
	}

	flagSet := newCommonFlagSet()
	command.PersistentFlags().AddFlagSet(flagSet.flags())

	command.AddCommand(
		NewAwsS3(&flagSet.commonFlags).NewCommand(),
		NewGcpStorage(&flagSet.commonFlags).NewCommand(),
		NewAzureBlob(&flagSet.commonFlags).NewCommand(),
	)

	return command
}

func executeCommand(provider string, client storage.Client, flags *commonFlags) (string, error) {
	processor := proc.New(flags.iterations, flags.parallelism, flags.objectSize)

	result := model.NewBenchmarkResult()
	result.Metadata.Provider = provider
	result.Metadata.Iterations = flags.iterations
	result.Metadata.Parallelism = flags.parallelism
	result.Metadata.ObjectSize = util.FormatUnit(float64(flags.objectSize), util.UnitsSize)

	// run write benchmark
	writeStats, err := processor.Run(proc.NewTask(
		"Write",
		func(i int) error { return client.Write(fmt.Sprintf("%s%d", objectPrefix, i)) },
	), flags.progressBar("Writes "))
	if err != nil {
		return "", err
	}
	result.Stats = append(result.Stats, writeStats)

	// run read benchmark
	if !flags.noReads {
		readStats, err := processor.Run(
			proc.NewTask(
				"Read",
				func(i int) error { return client.Read(fmt.Sprintf("%s%d", objectPrefix, i)) },
			),
			flags.progressBar("Reads  "),
		)
		if err != nil {
			return "", err
		}
		result.Stats = append(result.Stats, readStats)
	}

	// clean up the objects
	if _, err = processor.Run(proc.NewTask(
		"Cleanup",
		func(i int) error { return client.Delete(fmt.Sprintf("%s%d", objectPrefix, i)) },
	), flags.progressBar("Cleanup")); err != nil {
		return "", err
	}

	// render the result
	output, err := flags.renderer().Render(result)
	if err != nil {
		return "", err
	}

	return output, nil
}
