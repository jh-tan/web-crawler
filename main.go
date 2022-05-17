package main

import (
	"Crawler/Crawler"
)

func main() {
	workChannel := make(chan []string)
	processedChannel := make(chan string)

	go func() {
		workChannel <- []string{"https://www.sigure.tw/"}
	}()

	crawl.Crawl(workChannel, processedChannel)

}
