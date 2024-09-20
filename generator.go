package gogenerator

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// <-chan - canl somente leitura
func Title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			defer resp.Body.Close()
			if err != nil {
				fmt.Errorf("failed: %w", err)
			}
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}
