package crawl

import (
	"Crawler/Parser"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

func Crawl(worklist, process chan string, monitor chan int, wg *sync.WaitGroup) {
	for link := range worklist {
		parsedLink, err := FetchPage(link)
		monitor <- len(parsedLink)
		if err != nil || len(parsedLink) == 0 {
			return
		}
		for _, link := range parsedLink {
			go func(link string) {
				process <- link
			}(link)
		}
	}
	wg.Done()
}

func Process(worklist, process chan string, monitor chan int) {
	i := 1
	seen := make(map[string]struct{})
	for link := range process {
		monitor <- -1
		if _, ok := seen[link]; !ok {
			seen[link] = struct{}{}
			fmt.Println(i, link)
			i++
			go func(link string) {
				worklist <- link
			}(link)
		}
	}
}

func Monitor(work, process chan string, monitor chan int) {
	count := 0
	for i := range monitor {
		count += i
		// fmt.Println("Remaining: ", count, "Channel: ", i)
		if count == 0 {
			close(work)
			close(process)
			close(monitor)
		}
	}
}

func FetchPage(links string) ([]string, error) {
	resp, err := http.Get(links)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()
	return link.Parse(resp.Body, base), nil
}
