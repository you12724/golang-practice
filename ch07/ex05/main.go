package main

import (
	"errors"
	"io"
)

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{r, n}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type LimitedReader struct {
	R Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, errors.New("EOF")
	}
	if int64(len(p)) > l.N {
		p = p[0:l.N]
	}
	n, err = l.R.Read(p)
	l.N -= int64(n)
	return
}

func main() {
}
