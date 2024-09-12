package cli

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/reugn/cloud-storage-benchmark/internal/model"
)

type TableRenderer struct{}

var _ Renderer = (*TableRenderer)(nil)

func (r TableRenderer) Render(result *model.BenchmarkResult) (string, error) {
	var builder strings.Builder
	builder.WriteRune('\n')
	builder.WriteString(table([]model.Metadata{result.Metadata}))
	builder.WriteRune('\n')
	builder.WriteString(table(result.Stats))

	return builder.String(), nil
}

func table[T any](data []T) string {
	if len(data) == 0 {
		return ""
	}

	head := data[0]
	val := reflect.ValueOf(head)
	titles := make([]string, val.NumField())
	length := make([]int, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		titles[i] = val.Type().Field(i).Name
		length[i] = len(val.Type().Field(i).Name)
	}

	for _, item := range data {
		val := reflect.ValueOf(item)
		for i := 0; i < val.NumField(); i++ {
			value := fmt.Sprint(reflect.ValueOf(item).Field(i).Interface())
			fieldLen := len(value)
			if fieldLen > length[i] {
				length[i] = fieldLen
			}
		}
	}

	var builder strings.Builder
	for i, title := range titles {
		builder.WriteString(pad(title, length[i]))
	}
	builder.WriteRune('\n')

	for i := range titles {
		separator := pad(strings.Repeat("-", length[i]), length[i])
		builder.WriteString(separator)
	}
	builder.WriteRune('\n')

	for _, item := range data {
		v := reflect.ValueOf(item)
		for i := 0; i < v.NumField(); i++ {
			value := fmt.Sprint(v.Field(i).Interface())
			builder.WriteString(pad(value, length[i]))
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

func pad(str string, length int) string {
	if length > len(str) {
		return str + strings.Repeat(" ", length-len(str)+1)
	}
	return fmt.Sprintf("%s ", str)
}
