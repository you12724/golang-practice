package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("引数を入力してください")
		return
	}

	method := flag.String("method", "SHA256", "SHA384 or SHA512")
	flag.Parse()

	if *method == "SHA384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[len(os.Args)-1])))
	} else if *method == "SHA512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[len(os.Args)-1])))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[len(os.Args)-1])))
	}
}
