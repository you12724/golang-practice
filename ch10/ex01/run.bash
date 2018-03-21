#!/usr/bin/env bash

go build gopl.io/ch3/mandelbrot
./mandelbrot > mandelbrot.png

# png
go run main.go -f jpg < mandelbrot.png > mandelbrot.jpg
go run main.go -f gif < mandelbrot.png > mandelbrot.gif

# gif
go run main.go -f jpg < mandelbrot.gif > mandelbrot2.jpg
go run main.go -f png < mandelbrot.gif > mandelbrot2.png

# jpg
go run main.go -f png < mandelbrot.jpg > mandelbrot2.png
go run main.go -f gif < mandelbrot.jpg > mandelbrot2.gif
