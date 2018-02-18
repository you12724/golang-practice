package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type depthLinks struct {
	depth int
	urls  []string
}

var inputDepth = flag.Int("depth", 3, "depth of craw (default is 3)")

func crawl(depth int, url string) depthLinks {
	fmt.Println(url)
	if depth > *inputDepth {
		return depthLinks{depth + 1, nil}
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return depthLinks{depth + 1, list}
}

func main() {
	flag.Parse()
	fmt.Printf("input depth is %d\n", *inputDepth)
	worklist := make(chan depthLinks)
	n := 1

	go func() {
		worklist <- depthLinks{0, os.Args[1:]}
	}()

	for ; n > 0; n-- {
		seen := make(map[string]bool)
		list := <-worklist
		for _, url := range list.urls {
			if !seen[url] {
				seen[url] = true
				n++
				go func(depth int, url string) {
					worklist <- crawl(depth, url)
				}(list.depth, url)
			}
		}
	}
}
