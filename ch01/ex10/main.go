// url複数指定時に順番がoutputへの追記順が入れ替わった時に同一判定が正常にできないバグあり
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"time"
)

func main() {
	output1Str := "output1.txt"
	output2Str := "output2.txt"

	// 出力ファイル初期化, エラーは無視
	os.Remove(output1Str)
	os.Remove(output2Str)

	fetchAll(output1Str)
	fetchAll(output2Str)

	// 1回目と2回目の出力データの同一判定
	output1, err := ioutil.ReadFile(output1Str)
	if err != nil {
		fmt.Printf("not found %s\n", output1Str)
	}
	output2, err := ioutil.ReadFile(output2Str)
	if err != nil {
		fmt.Printf("not found %s\n", output2Str)
	}
	fmt.Println(reflect.DeepEqual(output1, output2))
}

func fetchAll(filename string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, filename)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, filename string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	writeFile(filename, bytes)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, len(bytes), url)
}

func writeFile(filename string, bytes []byte) {
	readBytes, err := ioutil.ReadFile(filename)
	if err == nil {
		bytes = append(readBytes, bytes...)
	}
	ioutil.WriteFile(filename, bytes, os.ModePerm)
}
