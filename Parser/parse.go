package link

import (
	"io"
	"strings"
	"golang.org/x/net/html"
)

func Parse(body io.Reader, base string) []string {
  
	z := html.NewTokenizer(body)
	result := make([]string, 0)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return result
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if token.Data != "a" {
				continue
			}

			for _, a := range token.Attr {
				if a.Key == "href" && matchDomain(a.Val,base) {
          result = append(result,filter(a.Val,base) )
				}
			}

		}
	}
}

func filter(links string, base string) string {
	if strings.HasPrefix(links, "/") {
		return base + links
	} else {
		return links
	}
}

func matchDomain(links string, base string) bool {
	return strings.HasPrefix(links,base) || strings.HasPrefix(links,"/")
}
