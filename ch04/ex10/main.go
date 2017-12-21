package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues: \n", result.TotalCount)

	onemonth := time.Now().AddDate(0, -1, 0)
	oneyear := time.Now().AddDate(-1, 0, 0)

	items := make(map[string]string)

	for _, item := range result.Items {
		str := fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		if oneyear.After(item.CreatedAt) {
			items["１年以上前"] += str
		} else if onemonth.Before(item.CreatedAt) {
			items["1ヶ月未満"] += str
		} else {
			items["１年未満"] += str
		}
	}

	for key, value := range items {
		fmt.Printf("%s\n%s\n", key, value)
	}
}
