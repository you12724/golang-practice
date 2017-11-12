package main

import (
	"bufio"
	"fmt"
	"golang-practice/ch02"
	"os"
	"strconv"
)

func main() {
	var input int
	if len(os.Args) == 1 {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		var err error
		input, err = strconv.Atoi(stdin.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "fm: %v\n", err)
			os.Exit(1)
		}
	} else {
		var err error
		input, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "fm: %v\n", err)
			os.Exit(1)
		}
	}
	feet := distconv.Feet(input)
	meter := distconv.Meter(input)
	fmt.Printf("%s = %s, %s = %s\n", meter, distconv.MToF(meter), feet, distconv.FToM(feet))
}
