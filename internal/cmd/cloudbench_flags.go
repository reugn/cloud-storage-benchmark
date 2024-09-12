package cmd

import (
	"fmt"
	"os"

	"github.com/reugn/cloud-storage-benchmark/internal/cli"
	"github.com/reugn/cloud-storage-benchmark/internal/data"
	"github.com/spf13/pflag"
)

const (
	objectSizeDefault = 5 * 1024 * 1024
)

type commonFlagSet struct {
	commonFlags
}

func newCommonFlagSet() *commonFlagSet {
	return &commonFlagSet{}
}

func (f *commonFlagSet) flags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.IntVarP(&f.iterations, "iter", "i",
		10, "The number of benchmark iterations")
	flagSet.IntVarP(&f.parallelism, "parallelism", "p",
		1, "The benchmark parallelism")
	flagSet.Int64VarP(&f.objectSize, "object", "o",
		objectSizeDefault, "The size of the benchmark object in bytes")
	flagSet.BoolVarP(&f.jsonOutput, "json", "j",
		false, "Return the benchmark result as JSON")
	flagSet.BoolVar(&f.noReads, "no-reads",
		false, "Skip read benchmark")
	flagSet.StringVarP(&f.filePath, "file", "f",
		"", "Input file path to use for upload")

	return flagSet
}

type commonFlags struct {
	iterations  int
	parallelism int
	objectSize  int64
	jsonOutput  bool
	noReads     bool
	filePath    string
}

func (f *commonFlags) validate() error {
	if f.iterations < 1 {
		return fmt.Errorf("invalid iterations: %d", f.iterations)
	}
	if f.parallelism < 1 {
		return fmt.Errorf("invalid parallelism: %d", f.parallelism)
	}
	if f.objectSize < objectSizeDefault {
		return fmt.Errorf("invalid object size: %d", f.objectSize)
	}
	if f.filePath != "" {
		fileInfo, err := os.Stat(f.filePath)
		if err != nil {
			return fmt.Errorf("file %s stat error: %w", f.filePath, err)
		}
		// set the objects size
		f.objectSize = fileInfo.Size()
	}

	return nil
}

func (f *commonFlags) renderer() cli.Renderer {
	if f.jsonOutput {
		return cli.JSONRenderer{}
	}
	return cli.TableRenderer{}
}

func (f *commonFlags) progressBar(description string) cli.ProgressBar {
	if f.jsonOutput {
		return cli.NewNoProgressBar()
	}
	return cli.NewDefaultProgressBar(f.iterations, description)
}

func (f *commonFlags) dataProvider() data.Provider {
	if f.filePath != "" {
		return data.NewFileProvider(f.filePath)
	}
	return data.NewRandomGenerator(f.objectSize)
}
