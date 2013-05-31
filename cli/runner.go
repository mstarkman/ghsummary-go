package cli

import (
	"fmt"

	"github.com/mstarkman/ghsummary.go/github"
)

func Run(githubUsername string) {
	user, err := github.FindUser(githubUsername)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Name)
	}
}
