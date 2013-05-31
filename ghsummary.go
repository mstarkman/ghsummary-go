package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mstarkman/ghsummary.go/github"
)

const defaultUserName = ""

func main() {
	githubUsername := flag.String("user", defaultUserName, "the GitHub username to generate the summary")
	flag.Parse()

	if *githubUsername == defaultUserName {
		fmt.Println("")
		flag.Usage()
		os.Exit(0)
	}

	user, err := github.FindUser(*githubUsername)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Name)
	}
}
