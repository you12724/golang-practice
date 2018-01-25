package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	for _, url := range list {
		// TODO: 修正が必要
		if strings.Contains(url, "golang.org") {
			resp, err := http.Get(url)
			defer resp.Body.Close()
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			bytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Print(err.Error())
				os.Exit(1)
			}

			makeFile(url, bytes)
		}
	}
	return list
}

func makeFile(url string, body []byte) {
	dirname := "pages"
	// フォルダの存在確認
	_, err := os.Stat(dirname)
	if err != nil {
		if err := os.Mkdir("pages", 0777); err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
	}
	names := strings.Split(url, "/")
	name := names[len(names)-1]
	if !strings.HasSuffix(name, ".html") {
		name += ".html"
	}
	ioutil.WriteFile(fmt.Sprintf("%s/%s", dirname, name), body, os.ModePerm)
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
