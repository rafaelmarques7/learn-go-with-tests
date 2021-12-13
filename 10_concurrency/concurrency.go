package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool, len(urls))
	// a channel is a way for goroutines to communicate with each other
	// remember that the main thread is also a goroutine
	// channels are usually two way endpoints:
	// you write to it (channel <- message) (in one goroutine)
	// you read from it (message := <- channel) (in another goroutine)
	resultsChannel := make(chan result)

	// parallelize the calls to WebsiteChecker
	for _, url := range urls {
		go func(u string) {
			// the "<-" symbol indicates that the "result" should be sent
			// to the resultsChannel
			resultsChannel <- result{u, wc(u)}
		}(url)
	}

	// run the writing to map linearly
	// this ensures there are no panics thrown,
	// since this could happen when trying to write a map at the same time
	for i := 0; i < len(urls); i++ {
		// receive expression
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
