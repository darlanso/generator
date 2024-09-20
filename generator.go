package gogenerator

import (
	"io"
	"net/http"
	"regexp"
)

// <-chan - canl somente leitura
func Title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			defer resp.Body.Close()
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[0]
		}(url)
	}
	return c
}
