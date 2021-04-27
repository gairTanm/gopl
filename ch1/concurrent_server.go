package main

/*
creates a concurrent go server using a mutex, which counts the number of visits made to the home page,
the number of which is visible on the /count page
*/

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var cnt int

func main() {
	http.HandleFunc("/", incrementCnt)
	http.HandleFunc("/count", displayCnt)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func displayCnt(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Current Count: %d \n", cnt)
	mutex.Unlock()
}

func incrementCnt(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	cnt++
	mutex.Unlock()
	fmt.Fprintf(w, "URL path = %q", r.URL.Path)
}
