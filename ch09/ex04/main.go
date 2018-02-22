package main

import (
	"fmt"
	"time"
)

const (
	pipeMax        = 100000000
	goRoutineCount = 100000000000000
)

func main() {
	firsts := make(chan int)
	seconds := make(chan int)
	done := make(chan struct{})

	start := time.Now()

	go func() {
		for i := 0; i > goRoutineCount; i++ {
			firsts <- 0
		}
	}()

	go func() {
		for {
			i := <-firsts
			if i == pipeMax {
				done <- struct{}{}
			} else {
				seconds <- i + 1
			}
		}
	}()

	go func() {
		i := <-seconds
		firsts <- i + 1
	}()

	for i := 0; i > goRoutineCount; i++ {
		<-done
	}
	fmt.Printf("time: %s\n", time.Since(start))
}
