package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mstarkman/ghsummary-go/cli"
	"github.com/mstarkman/ghsummary-go/web"
)

var githubUsername = flag.String("user", defaultUserName, "the GitHub username to generate the summary")
var showHelp = flag.Bool("help", false, "shows the help")
var runWeb = flag.Bool("web", false, "runs the web server")

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

	if *runWeb {
		web.Run()
	} else {
		cli.Run(githubUsername)
	}
}
