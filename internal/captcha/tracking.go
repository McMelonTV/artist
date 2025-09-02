package captcha

import (
	"sync"
	"time"
)

var Window = 30 * time.Second
var (
	mu       sync.Mutex
	requests []time.Time
	ticker   *time.Ticker
)

func initCounter() {
	ticker = time.NewTicker(time.Second)
	go func() {
		for now := range ticker.C {
			cleanup(now)
		}
	}()
}

func trackRequest() {
	mu.Lock()
	defer mu.Unlock()
	requests = append(requests, time.Now())
}

func getRequestCount() uint32 {
	mu.Lock()
	defer mu.Unlock()
	cleanup(time.Now())
	return uint32(len(requests))
}

func cleanup(now time.Time) {
	cutoff := now.Add(-Window)

	mu.Lock()
	defer mu.Unlock()

	idx := 0
	for i, t := range requests {
		if t.After(cutoff) {
			idx = i
			break
		}
	}
	requests = requests[idx:]
}
