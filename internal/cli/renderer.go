package cli

import "github.com/reugn/cloud-storage-benchmark/internal/model"

type Renderer interface {
	Render(result *model.BenchmarkResult) (string, error)
}
