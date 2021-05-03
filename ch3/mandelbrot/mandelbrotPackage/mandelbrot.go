package mandelbrotPackage

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func Mandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//ex3.5 full color mandelbrot
	const (
		red   = 19
		green = 22
		blue  = 17
		a     = 0
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - red*n, 255 - green*n, 255 - blue*n, 255 - a*n}
		}
	}
	return color.Black
}
