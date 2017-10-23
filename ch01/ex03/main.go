package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start1 := time.Now()
	echo1(os.Args[1:])
	fmt.Println(time.Since(start1))
	start2 := time.Now()
	echo2(os.Args[1:])
	fmt.Println(time.Since(start2))
}

func echo1(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s := strings.Join(args, " ")
	fmt.Println(s)
}
