package types

import "time"

// ResponseStats is the report about benchmark result.
type ResponseStats struct {
	// Total represents total number of requests.
	Total int
	// Failures represents number of failure request.
	Failures int
	// Duration means the time of benchmark.
	Duration time.Duration
	// Latencies represents the latency distribution in seconds.
	//
	// NOTE: The key represents quantile.
	Latencies map[float64]float64
}
