package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

func main() {
	urls := []string{
		"http://localhost:9999/1.txt",
		"http://localhost:9999/2.txt",
		"http://localhost:9999/3.txt",
		"http://localhost:9999/4.txt",
	}
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ul, err := url.Parse(u)
			fn := ul.Path[1:len(ul.Path)]
			res, err := http.Get(u)
			if err != nil {
				log.Println(err, u)
			}
			content, _ := ioutil.ReadAll(res.Body)
			ioutil.WriteFile(fn, content, 0644)
			res.Body.Close()
		}(u)
	}
	wg.Wait()
}
