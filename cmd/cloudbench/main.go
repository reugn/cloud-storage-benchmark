package main

import (
	"os"

	"github.com/reugn/cloud-storage-benchmark/internal/cmd"
)

var version = "dev"

func main() {
	rootCmd := cmd.NewCloudBench().NewCommand(version)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
