package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	max           = 1.0
	min           = -0.3
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		""+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// if i == j || i+1 == j || i == j+1 {
			// 	continue
			// }
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if math.IsNaN(ax) ||
				math.IsNaN(ay) ||
				math.IsNaN(bx) ||
				math.IsNaN(by) ||
				math.IsNaN(cx) ||
				math.IsNaN(cy) ||
				math.IsNaN(dx) ||
				math.IsNaN(dy) {
				continue
			}
			color := getColor(az, bz, cz, dz)
			fmt.Printf("<polygon fill='#%s' stroke='#%s' points='%g,%g %g,%g %g,%g %g,%g '/>\n", color, color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
}

func getColor(h1, h2, h3, h4 float64) string {
	height := (h1 + h2 + h3 + h4) / 4

	delta := uint32((max - height) / (max - min) * 255)

	c := (0xff0000 - delta<<16) + delta
	return prependZeros(fmt.Sprintf("%X", c))
}

func prependZeros(hex string) string {
	result := hex
	for i := len(hex); i < 6; i++ {
		result = "0" + result
	}
	return result
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
