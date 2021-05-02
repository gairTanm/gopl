package surface

import (
	"fmt"
	"io"
	"math"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Plot(out io.Writer, height int64, width int64, color string) {
	if height == 0 {
		height = 320
	}
	if width == 0 {
		width = 600
	}
	if color == "" {
		color = "gray"
	}
	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", color, width, height)
	io.WriteString(out, s)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j, float64(height), float64(width))
			bx, by, _ := corner(i, j, float64(height), float64(width))
			cx, cy, _ := corner(i, j+1, float64(height), float64(width))
			dx, dy, _ := corner(i+1, j+1, float64(height), float64(width))
			// ex 3.2, colored peaks and valleys
			stroke := color
			/*if az < 0 && bz < 0 && cz < 0 && dz < 0 {
				stroke = "#0000ff"
			}*/
			s = fmt.Sprintf("<polygon style='stroke:%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				stroke, ax, ay, bx, by, cx, cy, dx, dy)
			io.WriteString(out, s)
		}
	}
	io.WriteString(out, "</svg>")
}
func corner(i, j int, height, width float64) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	xyscale := width / 2 / xyrange
	zscale := height * 0.4
	xyrange := 30.0
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	//ex 3.1
	if r == 0 {
		return 0.0
	}
	return math.Sin(r) / r
}
