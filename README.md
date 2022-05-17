# web-crawler
Trying out Go's concurrency features to parallelize a web crawler. 
Number of site/pages able to crawl in N seconds vary from site to site.
Tested Site:

 - Wikipedia - ~10 seconds for 500k pages
 - Youtube - ~ 30 seconds for 100k pages
 - Medium - ~ 30 seconds for 50k pages, and get rate limited
 - Investing - ~ 30seconds for 40k pages, and get rate limited
 - Baike.Baidu - ~ 35 seconds for 50k pages

### Current Issue
 - Stability
 - Unable to close the channel and stop the program properly
 - Get rate limited easily, and the program will just stuck there.
