package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeFetcher struct {
	mu      sync.Mutex
	fetched map[string]bool
	wg      sync.WaitGroup
}

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, f *SafeFetcher) {
	if depth <= 0 {
		return
	}
	f.mu.Lock()
	if f.fetched[url] == false {
		body, urls, err := fetcher.Fetch(url)
		f.fetched[url] = true
		f.mu.Unlock()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go Crawl(u, depth-1, fetcher, f)
		}
	} else {
		fmt.Printf("visited: %s\n", url)
		f.mu.Unlock()
	}
	return
}

func main() {
	f := SafeFetcher{fetched: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &f)
	time.Sleep(2 * time.Second)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
