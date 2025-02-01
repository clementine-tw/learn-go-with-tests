package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := range len(urls) {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for range b.N {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
