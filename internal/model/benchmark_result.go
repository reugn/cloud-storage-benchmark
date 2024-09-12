package model

import (
	"runtime"
	"time"
)

type BenchmarkResult struct {
	Metadata Metadata `json:"metadata"`
	Stats    []Stats  `json:"stats"`
}

func NewBenchmarkResult() *BenchmarkResult {
	metadata := Metadata{
		OS:    runtime.GOOS,
		Cores: runtime.NumCPU(),
	}
	return &BenchmarkResult{
		Metadata: metadata,
		Stats:    make([]Stats, 0, 2), // read and write
	}
}

type Stats struct {
	Benchmark     string        `json:"benchmark"`
	Min           time.Duration `json:"min"`
	Avg           time.Duration `json:"avg"`
	P50           time.Duration `json:"p50"`
	P95           time.Duration `json:"p95"`
	P99           time.Duration `json:"p99"`
	Max           time.Duration `json:"max"`
	AvgThroughput string        `json:"avg-throughput"`
}

type Metadata struct {
	OS    string `json:"os"`
	Cores int    `json:"cores"`

	Provider    string `json:"provider"`
	Iterations  int    `json:"iterations"`
	Parallelism int    `json:"parallelism"`
	ObjectSize  string `json:"object-size"`
}
