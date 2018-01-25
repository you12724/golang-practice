package main

import (
	"io"
	"os"
)

type WriterWrapper struct {
	c int64
	w io.Writer
}

func (ww *WriterWrapper) Write(p []byte) (int, error) {
	num, err := ww.w.Write(p)
	ww.c += int64(num)
	return num, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var ww WriterWrapper
	ww.w = w
	return &ww, &(ww.c)
}

func main() {
	w, num := CountingWriter(os.Stdout)
	w.Write([]byte("this is 10"))
	w.Write([]byte("plus 7."))
	println(*num)
}
