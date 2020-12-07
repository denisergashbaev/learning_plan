//https://tour.golang.org/concurrency/10
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlsRegistry struct {
	mu sync.Mutex
	v  map[string]bool
}

//Returns true if this url has not yet been processed and marks it for processing. If the url is already marked for processing, return false
func (f *UrlsRegistry) Process(key string) bool {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.v[key]; ok {
		//fmt.Printf("key %s exists, returning false\n", key)
		return false
	} else {
		//fmt.Printf("key %s does not exist, returning true\n", key)
		f.v[key] = true
		return true
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, urlsRegistry *UrlsRegistry, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer wg.Done()

	if depth <= 0 {
		return
	}
	if !urlsRegistry.Process(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		var wgInternal sync.WaitGroup
		wgInternal.Add(1)
		go Crawl(u, depth-1, fetcher, urlsRegistry, &wgInternal)
		wgInternal.Wait()
	}
	return
}

func main() {
	urlsRegistry := UrlsRegistry{v: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &urlsRegistry, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	} else {
		return "", nil, fmt.Errorf("not found: %s", url)
	}
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
