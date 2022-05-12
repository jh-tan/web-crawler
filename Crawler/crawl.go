package crawl

import (
	"Crawler/Parser"
	"net/http"
	"net/url"
	"sync"
)

type SafeMap struct {
	v   map[string]bool
	mux sync.Mutex
}

type SafeResult struct {
	url []string
	mux sync.Mutex
}

// SetVal sets the value for the given key.
func (m *SafeMap) SetVal(key string, val bool) {
	m.mux.Lock()
	m.v[key] = val
	m.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (m *SafeMap) GetVal(key string) bool {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.v[key]
}

func (r *SafeResult) AppendURL(url string) {
	r.mux.Lock()
	r.url = append(r.url, url)
	r.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (r *SafeResult) GetAllResult() []string {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.url
}

type Spider struct {
	UrlChannel      chan string
	progressChannel chan bool
	Result          *SafeResult
	seen            *SafeMap
}

func NewSpider() *Spider {
	return &Spider{
		UrlChannel:      make(chan string),
		progressChannel: make(chan bool),
		Result:          &SafeResult{url: make([]string, 0)},
		seen:            &SafeMap{v: make(map[string]bool)},
	}
}

func (spider *Spider) Start(wg *sync.WaitGroup) {

	for {
		select {
    // TO CHANGE
		case url := <-spider.UrlChannel:
			go Crawl(url, spider)
		case <-spider.progressChannel:
			wg.Done()
			return
		}
	}

}

func Crawl(rootURL string, spider *Spider) {
	parsedLink, err := fetchPage(rootURL)
	if err != nil {
		return
	}

	for _, url := range parsedLink {
		if spider.seen.GetVal(url) {
			continue
		}
		if len(spider.Result.GetAllResult()) >= 50000 {
			spider.progressChannel <- true
			break
		}
		spider.seen.SetVal(url, true)
		spider.Result.AppendURL(url)
		spider.UrlChannel <- url
	}

}

func fetchPage(links string) ([]string, error) {
	resp, err := http.Get(links)
	if err != nil {
		// fmt.Println("Server unable to reach")
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
