# Cloud Storage Benchmark
[![Build](https://github.com/reugn/cloud-storage-benchmark/actions/workflows/build.yml/badge.svg)](https://github.com/reugn/cloud-storage-benchmark/actions/workflows/build.yml)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/reugn/cloud-storage-benchmark)](https://pkg.go.dev/github.com/reugn/cloud-storage-benchmark)
[![Go Report Card](https://goreportcard.com/badge/github.com/reugn/cloud-storage-benchmark)](https://goreportcard.com/report/github.com/reugn/cloud-storage-benchmark)

A CLI tool to evaluate performance of various cloud storage providers.

## Supported storage providers
- [x] [AWS S3](https://aws.amazon.com/s3/)
- [x] [GCP Storage](https://cloud.google.com/storage/)
- [x] [Azure Blob Storage](https://azure.microsoft.com/en-us/products/storage/blobs/)

## Build from source
Download and [install Go](https://golang.org/doc/install).

Install the application:

```sh
go install github.com/reugn/cloud-storage-benchmark/cmd/cloudbench@latest
```

See the [go install](https://go.dev/ref/mod#go-install) instructions for more information about the command.

## Usage
```console
Cloud Storage Benchmark

Usage:
  cloudbench [command]

Available Commands:
  aws-s3      Benchmark AWS S3 I/O operations
  azure-blob  Benchmark Azure Blob I/O operations
  completion  Generate the autocompletion script for the specified shell
  gcp-storage Benchmark GCP Storage I/O operations
  help        Help about any command

Flags:
  -f, --file string       Input file path to use for upload
  -h, --help              help for cloudbench
  -i, --iter int          The number of benchmark iterations (default 10)
  -j, --json              Return the benchmark result as JSON
      --no-reads          Skip read benchmark
  -o, --object int        The size of the benchmark object in bytes (default 5242880)
  -p, --parallelism int   The benchmark parallelism (default 1)
  -v, --version           version for cloudbench
```

## License
MIT
