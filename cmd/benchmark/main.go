package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// StatusReport carries info about one HTTP request
type StatusReport struct {
	Duration   time.Duration
	HTTPStatus int
}

func main() {
	// --- CLI flags ---
	var totalRequests int
	var concurrency int
	var url string

	flag.IntVar(&totalRequests, "requests", 100, "Number of requests")
	flag.IntVar(&concurrency, "concurrency", 10, "Concurrency level")
	flag.StringVar(&url, "url", "", "Target URL")
	flag.Parse()

	if url == "" {
		fmt.Println("Missing -url parameter")
		return
	}

	// --- Channels for coordination ---
	requests := make(chan int, totalRequests)
	results := make(chan StatusReport, totalRequests)

	var wg sync.WaitGroup

	// --- Workers ---
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := &http.Client{}

			for range requests {
				start := time.Now()
				resp, err := client.Get(url)
				duration := time.Since(start)

				report := StatusReport{Duration: duration, HTTPStatus: 0}

				if err == nil {
					report.HTTPStatus = resp.StatusCode
					resp.Body.Close()
				}

				results <- report
			}
		}()
	}

	// --- Feed jobs ---
	for i := 0; i < totalRequests; i++ {
		requests <- i
	}
	close(requests)

	// --- Close results after all workers finish ---
	go func() {
		wg.Wait()
		close(results)
	}()

	// --- Collect stats ---
	var total time.Duration
	var count int
	statusCounts := make(map[int]int)

	startWall := time.Now()
	for report := range results {
		total += report.Duration
		count++
		statusCounts[report.HTTPStatus]++
	}
	wallDuration := time.Since(startWall)

	// --- Report ---
	fmt.Printf("\nLoad test complete for %s\n", url)
	fmt.Printf("Total requests: %d\n", count)
	fmt.Printf("Total wall time: %v\n", wallDuration)
	fmt.Printf("Average time per request: %v\n", total/time.Duration(count))
	fmt.Printf("Requests/sec: %.2f\n", float64(count)/wallDuration.Seconds())

	fmt.Println("Status code distribution:")
	for code, qty := range statusCounts {
		fmt.Printf("  %d: %d\n", code, qty)
	}
}
