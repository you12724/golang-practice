package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	xkcdURL  = "https://xkcd.com/"
	infoJson = "info.0.json"
)

type Comic struct {
	Num        int
	Year       string
	Month      string
	Day        string
	Title      string
	Link       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
}

func main() {
	if os.Args[1] == "read" {
		num, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println(err.Error())
		}
		readComic(num)
		return
	}
	if os.Args[1] == "fetch" {
		num, err := getNumberOfComics()
		if err != nil {
			fmt.Printf("Failed(%v)\n", err)
			os.Exit(1)
		}
		fetchAllComics(num)
		fmt.Printf("Done\n")
	}
}

func fetchAllComics(num int) {
	var comics = make(map[int]*Comic)

	// MEMO: 全部取ってると時間かかるので100
	for i := 1; i < 100; i++ {
		resp, err := http.Get(fmt.Sprintf("%s%d/info.0.json", xkcdURL, i))
		if err != nil {
			println(err.Error())
			return
		}

		var result Comic
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			println(err.Error())
			resp.Body.Close()
			os.Exit(1)
		}
		resp.Body.Close()
		comics[result.Num] = &result
	}

	for key, value := range comics {
		ioutil.WriteFile(fmt.Sprintf("comics/comic%d.txt", key), []byte(fmt.Sprintf("%s%d\n%s", xkcdURL, value.Num, value.Transcript)), os.ModePerm)
	}
}

func getNumberOfComics() (int, error) {
	resp, err := http.Get(xkcdURL + infoJson)
	if err != nil {
		return -1, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return -1, fmt.Errorf("Get Failed: %s", resp.Status)
	}

	var comicInfo struct {
		Num int
	}
	jsonDecoder := json.NewDecoder(resp.Body)
	if err := jsonDecoder.Decode(&comicInfo); err != nil {
		return -1, err
	}
	return comicInfo.Num, nil
}

func readComic(num int) {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("comics/comic%d.txt", num))
	if err != nil {
		panic(nil)
	}

	fmt.Println(string(bytes))

}
