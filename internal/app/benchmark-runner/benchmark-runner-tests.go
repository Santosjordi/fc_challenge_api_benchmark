package benchmarkrunner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBenchmarkRunner_BasicSuccess(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := benchmark.BenchmarkConfig{
		TotalRequests: 10,
		Concurrency:   2,
		URL:           server.URL,
		Client:        &http.Client{},
	}

	runner := benchmark.NewBenchmarkRunner(cfg)
	result, err := runner.Run()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.TotalRequests != 10 {
		t.Errorf("Expected 10 total requests, got %d", result.TotalRequests)
	}

	if result.Failures != 0 {
		t.Errorf("Expected 0 failures, got %d", result.Failures)
	}

	if result.Successes != 10 {
		t.Errorf("Expected 10 successes, got %d", result.Successes)
	}

	if result.AvgLatency < 10*time.Millisecond {
		t.Errorf("Expected latency >= 10ms, got %v", result.AvgLatency)
	}
}
