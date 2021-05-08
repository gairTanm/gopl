package main

import (
	"fmt"
	"gopl/ch4/github/github"
	"os"
	"time"
)

type TimeIssue struct {
	Time  string
	Items []github.Issue
}

func printIssues(issue TimeIssue) {
	fmt.Printf("\n%s:\n", issue.Time)
	for _, item := range issue.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var monthOld, lessThanYearOld, old TimeIssue
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if time.Now().Year()-item.CreatedAt.Year() > 1 {
			old.Items = append(old.Items, *item)
		} else if time.Now().Month() == item.CreatedAt.Month() {
			monthOld.Items = append(monthOld.Items, *item)
		} else {
			lessThanYearOld.Items = append(lessThanYearOld.Items, *item)
		}
	}
	monthOld.Time = "Recent issues"
	lessThanYearOld.Time = "Issues less than a year old"
	old.Time = "Old issues"
	printIssues(monthOld)
	printIssues(lessThanYearOld)
	printIssues(old)
	fmt.Println()
}
