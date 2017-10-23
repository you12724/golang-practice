package main

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
	"strings"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	cycles := getCycles(r)
	if cycles > 0 {
		lissajous(w, cycles)
	} else {
		fmt.Fprintf(w, "invalid params")
	}
}

func getCycles(r *http.Request) (cycles int) {
	query := strings.Split(r.URL.RawQuery, "&")
	result := 5
	for _, params := range query {
		param := strings.Split(params, "=")
		if param[0] == "cycles" {
			cycles, err := strconv.Atoi(param[1])
			if err != nil {
				result = -1
				break
			}
			result = cycles
		}
	}
	return result
}

func lissajous(out io.Writer, cycles int) {
	var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

	const (
		whiteIndex = 0
		greenIndex = 1
	)

	const (
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
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
