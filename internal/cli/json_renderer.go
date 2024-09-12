package cli

import (
	"encoding/json"

	"github.com/reugn/cloud-storage-benchmark/internal/model"
)

type JSONRenderer struct{}

var _ Renderer = (*JSONRenderer)(nil)

func (r JSONRenderer) Render(result *model.BenchmarkResult) (string, error) {
	output, err := json.MarshalIndent(&result, "", "  ")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
