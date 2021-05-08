package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const omdbURL = "http://www.omdbapi.com/?apikey=[your api key here]&t="

func main() {
	q := strings.Join(os.Args[1:], "+")
	r, err := http.Get(omdbURL + q)
	if err != nil {
		fmt.Printf("Error occured %v", err)
	}
	if r.StatusCode != http.StatusOK {
		r.Body.Close()
		fmt.Println("Error", r.Status)
	}
	var result MovieSearchResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		r.Body.Close()
	}
	fmt.Print("Fetched\n")
	r.Body.Close()
	fmt.Println(result.Poster)
	r, err = http.Get(result.Poster)
	if err != nil {
		fmt.Println("Error occurred while fetching the poster", err)
		os.Exit(1)
	}
	fmt.Printf("Fetched the poster\n")
	file, err := os.Create(fmt.Sprintf("tmp/%s.jpg", q))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
	r.Body.Close()
}

type MovieSearchResult struct {
	Poster string
}
