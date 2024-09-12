package util

import (
	"errors"
	"sort"
	"time"
)

func ToFloat(slice []time.Duration) []float64 {
	floatSlice := make([]float64, len(slice))
	for i, v := range slice {
		floatSlice[i] = float64(v)
	}
	return floatSlice
}

func Percentile(slice []float64, percentile float64) (float64, error) {
	if len(slice) == 0 {
		return 0, errors.New("slice is empty")
	}
	if percentile < 0 || percentile > 1 {
		return 0, errors.New("percentile must be between 0 and 1")
	}

	sort.Float64s(slice)
	index := int(float64(len(slice)) * percentile)

	return slice[index], nil
}
