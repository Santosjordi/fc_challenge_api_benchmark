package benchmarkrunner

import (
	"net/http"
	"time"
)

type BenchmarkResult struct {
	TotalRequests int
	Successes     int
	Failures      int
	AvgLatency    time.Duration
}

type BenchmarkConfig struct {
	TotalRequests int
	Concurrency   int
	URL           string
	Client        *http.Client // allows injection in tests
}

type BenchmarkRunner struct {
	config BenchmarkConfig
}

func NewBenchmarkRunner(cfg BenchmarkConfig) *BenchmarkRunner {
	return &BenchmarkRunner{config: cfg}
}

func (br *BenchmarkRunner) Run() (*BenchmarkResult, error) {
	// your logic goes here
	return &BenchmarkResult{}, nil
}
