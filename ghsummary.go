package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mstarkman/ghsummary.go/github"
)

var githubUsername = flag.String("user", defaultUserName, "the GitHub username to generate the summary")
var showHelp = flag.Bool("help", false, "shows the help")

const defaultUserName = ""

func parseCli() bool {
	flag.Parse()

	if *showHelp {
		return false
	}

	if *githubUsername == defaultUserName {
		return false
	}

	return true
}

func main() {
	ok := parseCli()

	if !ok {
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
