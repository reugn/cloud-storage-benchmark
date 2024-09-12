package util

import "fmt"

var (
	UnitsRate = []string{"B/s", "kB/s", "MB/s", "GB/s", "TB/s", "PB/s"}
	UnitsSize = []string{"B", "kB", "MB", "GB", "TB", "PB"}
)

func FormatUnit(num float64, units []string) string {
	base := 1024.0
	unitsLimit := len(units)
	i := 0
	for num >= base && i < unitsLimit {
		num /= base
		i++
	}

	f := "%.0f %s"
	if i > 1 {
		f = "%.2f %s"
	}

	return fmt.Sprintf(f, num, units[i])
}
