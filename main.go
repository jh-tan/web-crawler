package main

import (
	"Crawler/Crawler"
	"sync"
)

func main() {
	urlChannel := make(chan string)
	processChannel := make(chan string)
	monitorChannel := make(chan int)
	var wg sync.WaitGroup

	go func() {
		urlChannel <- "https://en.wikipedia.org/wiki/Main_Page"
		monitorChannel <- 1
	}()

	go crawl.Process(urlChannel, processChannel, monitorChannel)
	go crawl.Monitor(urlChannel, processChannel, monitorChannel)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go crawl.Crawl(urlChannel, processChannel, monitorChannel, &wg)
	}
	wg.Wait()
}
