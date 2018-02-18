package main

import (
	"fmt"
	"golang-practice/ch07/ex16/eval"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/calc", calculation)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func calculation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	exps, ok := r.Form["exp"]
	if !ok {
		http.Error(w, "expに式を入力してください", http.StatusBadRequest)
		return
	}
	for _, exp := range exps {
		result, err := eval.Parse(exp)
		fmt.Printf("%s\n", exp)
		if err != nil {
			http.Error(w, "expに正しい式を入力してください", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "%s = %f\n", exp, result.Eval(eval.Env{}))
	}
}
