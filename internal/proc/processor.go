package proc

import (
	"slices"
	"sync"
	"time"

	"github.com/reugn/cloud-storage-benchmark/internal/cli"
	"github.com/reugn/cloud-storage-benchmark/internal/model"
	"github.com/reugn/cloud-storage-benchmark/internal/util"
)

type Task struct {
	name string
	job  func(int) error
}

func NewTask(name string, job func(int) error) *Task {
	return &Task{
		name: name,
		job:  job,
	}
}

type Processor interface {
	Run(task *Task, progressBar cli.ProgressBar) (model.Stats, error)
}

type processor struct {
	tokens     chan struct{}
	errors     chan error
	iterations int
	objectSize int64
}

var _ Processor = (*processor)(nil)

func New(iterations, parallelism int, objectSize int64) Processor {
	processor := &processor{
		tokens:     make(chan struct{}, parallelism),
		errors:     make(chan error, parallelism),
		iterations: iterations,
		objectSize: objectSize,
	}

	for i := 0; i < parallelism; i++ {
		processor.tokens <- struct{}{}
	}

	return processor
}

func (p *processor) Run(task *Task, progressBar cli.ProgressBar) (model.Stats, error) {
	var (
		wg      sync.WaitGroup
		mtx     sync.Mutex
		latency []time.Duration
	)

	for i := 0; i < p.iterations; i++ {
		wg.Add(1)

		select {
		case <-p.tokens:
		case err := <-p.errors:
			return model.Stats{}, err
		}

		go func(n int) {
			defer func() {
				p.tokens <- struct{}{}
				wg.Done()
			}()

			t := time.Now()
			if err := task.job(n); err != nil {
				p.errors <- err
				return
			}
			elapsed := time.Since(t)

			mtx.Lock()
			latency = append(latency, elapsed)
			mtx.Unlock()

			progressBar.Add(1)
		}(i)
	}

	wg.Wait()
	progressBar.Clear()

	return p.buildStats(task.name, latency)
}

func (p *processor) buildStats(taskName string, latency []time.Duration) (model.Stats, error) {
	avgL := sum(latency) / time.Duration(p.iterations)
	minL := slices.Min[[]time.Duration](latency)
	maxL := slices.Max[[]time.Duration](latency)
	// calculate the average throughput rate in bytes per second
	avgThroughput := (float64(p.objectSize)) / (avgL.Seconds())

	floatLatency := util.ToFloat(latency)
	p50, _ := util.Percentile(floatLatency, 0.50)
	p95, _ := util.Percentile(floatLatency, 0.95)
	p99, _ := util.Percentile(floatLatency, 0.99)

	return model.Stats{
		Benchmark:     taskName,
		Min:           minL,
		Avg:           avgL,
		P50:           time.Duration(p50),
		P95:           time.Duration(p95),
		P99:           time.Duration(p99),
		Max:           maxL,
		AvgThroughput: util.FormatUnit(avgThroughput, util.UnitsRate),
	}, nil
}

func sum(duration []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range duration {
		total += d
	}
	return total
}
