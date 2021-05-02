package main

import (
	"gopl/ch3/surface"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var height, width int64
	var color string
	q := r.URL.Query()
	for param, arg := range q {
		switch param {
		case "height":
			height, _ = strconv.ParseInt(arg[0], 10, 64)
		case "width":
			width, _ = strconv.ParseInt(arg[0], 10, 64)
		case "color":
			color = arg[0]
		}
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Plot(w, height, width, color)
}
