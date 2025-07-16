package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var totalRequests int
	var concurrency int
	var url string

	flag.IntVar(&totalRequests, "n", 100, "Total number of request")
	flag.IntVar(&concurrency, "c", 10, "Number of concurrent Requests")
	flag.StringVar(&url, "url", "", "Target url")
	flag.Parse()

	if url == "" {
		fmt.Println("Please provide a URL with -url")
		return
	}

	var wg sync.WaitGroup
	requests := make(chan int, totalRequests)
	results := make(chan time.Duration, totalRequests)

	start := time.Now()

	// Lauch goroutines
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := http.Client{}

			for range requests {
				t0 := time.Now()
				resp, err := client.Get(url)
				if err == nil {
					resp.Body.Close()
				}
				results <- time.Since(t0)
			}
		}()
	}

	// sends "n" requests
	for i := 0; i < totalRequests; i++ {
		requests <- i
	}
	close(requests)

	wg.Wait()
	close(requests)

	totalDuration := time.Since(start)

	var sum time.Duration
	for r := range results {
		sum += r
	}

	fmt.Printf("Total Time: %v\n", totalDuration)
	fmt.Printf("Average Request Time: %v\n", sum/time.Duration(totalRequests))
	fmt.Printf("Requests/sec: %.2f\n", float64(totalRequests)/totalDuration.Seconds())
}
