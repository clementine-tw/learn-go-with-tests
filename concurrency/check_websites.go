package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(check WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, check(url)}
		}()
	}

	for range len(urls) {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
