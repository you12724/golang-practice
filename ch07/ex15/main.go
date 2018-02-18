package main

import (
	"bufio"
	"fmt"
	"os"

	"golang-practice/ch07/ex15/eval"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("式を入力してください\n")
		exp, _, err := reader.ReadLine()
		if err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(1)
			return
		}
		expr, err := eval.Parse(string(exp))
		if err != nil {
			fmt.Printf("%s", err.Error())
			os.Exit(1)
			return
		}
		println(expr.Eval(eval.Env{}))
	}
}
