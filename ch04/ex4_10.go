package main

import (
	"fmt"
	"log"
	"time"
	"os"
	"./github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	
	now := time.Now()
	
	fmt.Println("Issues less than 30 days old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	
	fmt.Println("Issues less than 1 year old:")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 <= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}	
}

//!-