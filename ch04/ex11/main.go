package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	baseURL = "https://api.github.com/repos/you12724/golang-practice/issues"
)

type Issue struct {
	Number    int
	Title     string
	State     string
	CreatedAt time.Time `json:"created_at"`
	User      *User
}

type User struct {
	Name string `json:"login"`
}

func main() {
	// issues := indexIssues()
	// fmt.Print("number, title, username, state, created_at\n")
	// for _, issue := range issues {
	// 	fmt.Printf("%v, %v, %v, %v, %v\n", issue.Number, issue.Title, issue.User.Name, issue.State, issue.CreatedAt)
	// }
	// createIssue()
	println("2段階認証で力尽きました・・・。申し訳ないです・・・。")
}

func indexIssues() []Issue {
	url := baseURL
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		os.Exit(1)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		os.Exit(1)
	}
	resp.Body.Close()
	return result
}

func createIssue() {
	issue := Issue{Title: "faafsdafas"}
	data, err := json.Marshal(issue)
	if err != nil {
		println(err.Error())
		return
	}
	req, err := http.NewRequest("POST", baseURL, bytes.NewReader(data))
	if err != nil {
		println(err.Error())
		return
	}
	req.SetBasicAuth("", "")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%v", string(b))

}
