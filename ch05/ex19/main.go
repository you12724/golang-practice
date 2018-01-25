package main

import "fmt"

func main() {
	println(example2())
}

func example() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("error: %v", p)
		}
	}()

	panic("returnなし")
}

func example2() (test string) {
	defer func() {
		if p := recover(); p != nil {
			test = fmt.Sprintf("return %v", p)
		}
	}()

	panic("テスト")
}
