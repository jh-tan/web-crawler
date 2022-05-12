package main

import (
	"Crawler/Crawler"
	"fmt"
	"sync"
)

/*
  Things to do:
  1. Separate into different function
  2. Visit
  3. Prevent loop
  4. Multithreaded
  5. ..
*/

func main() {
	spider := crawl.NewSpider()
	var wg sync.WaitGroup

	go func() {
		spider.UrlChannel <- "https://youtube.com"
	}()

	wg.Add(1)
	go spider.Start(&wg)
	// go spider.Monitor(&wg)

	wg.Wait()
	// for _, url := range spider.Result.GetAllResult() {
	// 	fmt.Println(url)
	// }
	fmt.Println(len(spider.Result.GetAllResult()))
}
