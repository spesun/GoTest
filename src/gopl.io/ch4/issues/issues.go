package main

import (
	"os"
	"log"
	"fmt"
	"gopl.io/ch4/github"
)

func main() {

	second :=  10000000000
	result, err := github.SearchIssues(os.Args[1:], second)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)

	}
}
