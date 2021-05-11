package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("http", url) {
			url = "http://" + url
		}
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s, %d\n", filename, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		if closeErr := f.Close(); err != nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}
	return local, n, err
}
