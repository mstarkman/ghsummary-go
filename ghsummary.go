package main

import (
	"fmt"

	"github.com/mstarkman/ghsummary.go/github"
)

func main() {
	user, err := github.FindUser("mstarkman")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Name)
	}
}
