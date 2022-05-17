package crawl

import (
	"Crawler/Parser"
	"fmt"
	"net/http"
	"net/url"
)

func Crawl(worklist chan []string, processedLink chan string) {
	i, n := 1, 1
	seen := make(map[string]struct{})

	for i := 0; i < 10000; i++ {
		go func() {
			for link := range processedLink {
				parsedLink, err := FetchPage(link)
				if err != nil {
					continue
				}
				go func() {
					worklist <- parsedLink
				}()
			}
		}()
	}

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if _, ok := seen[link]; !ok {
				seen[link] = struct{}{}
				fmt.Println(i, link)
				i++
				n++
				processedLink <- link
			}
		}
	}
	close(processedLink)
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
