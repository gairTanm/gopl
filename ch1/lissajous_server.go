package main

/*
go run lissajous_server.go
localhost:8000/?cycles=<number of cycles>
*/

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var paletteS = []color.Color{color.White, color.Black}

const (
	whiteIndexS = 0
	blackIndexS = 1
)

func main() {
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//ex 1.12 custom cycles
func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q) > 0 {
		cycles, err := strconv.Atoi(q["cycles"][0])
		if err != nil {
			fmt.Println(err)
		}
		cycle := float64(cycles)
		lissajousS(w, cycle)

	}
	lissajousS(w, 5)
}
func lissajousS(out io.Writer, cycle float64) {
	var (
		cycles  = cycle
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, paletteS)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndexS)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
