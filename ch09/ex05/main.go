package main

import (
	"log"
	"os"
	"time"
)

func main() {
	first := make(chan string)
	second := make(chan string)
	done := make(chan struct{})

	count := 0
	tmpCount := 0

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			f := <-first
			second <- "hi" + f
			count++
		}
	}()

	go func() {
		for {
			s := <-second
			first <- "hi" + s
			count++
		}
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	first <- ""

loop:
	for {
		select {
		case <-ticker.C:
			log.Printf("%d count / second", count-tmpCount)
			tmpCount = count
		case <-done:
			ticker.Stop()
			break loop
		}
	}

	println("finish")
}
