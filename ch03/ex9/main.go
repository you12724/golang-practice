package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
		width, height          = 1024, 1024
	)

	var inputMaxX, inputMinX float64
	tmpX, err := strconv.ParseFloat(r.FormValue("x"), 64)
	if err == nil {
		inputMaxX = -tmpX
		inputMinX = tmpX
	} else {
		inputMaxX = xmax
		inputMinX = xmin
	}

	var inputMaxY, inputMinY float64
	tmpY, err := strconv.ParseFloat(r.FormValue("y"), 64)
	if err == nil {
		inputMaxY = -tmpY
		inputMinY = tmpY
	} else {
		inputMaxY = ymax
		inputMinY = ymin
	}

	var inputWidth, inputHeight int
	scale, err := strconv.ParseFloat(r.FormValue("scale"), 64)
	if err == nil {
		inputWidth = int(width * scale)
		inputHeight = int(height * scale)
	} else {
		inputWidth = width
		inputHeight = height
	}

	img := image.NewRGBA(image.Rect(0, 0, inputWidth, inputHeight))
	for py := 0; py < inputHeight; py++ {
		y := float64(py)/float64(inputHeight)*(inputMaxY-inputMinY) + inputMinY
		for px := 0; px < inputWidth; px++ {
			x := float64(px)/float64(inputWidth)*(inputMaxX-inputMinX) + inputMinX
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
