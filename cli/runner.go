package cli

import (
	"fmt"

	"github.com/mstarkman/ghsummary.go/github"
)

func Run(githubUsername *string) {
	githubClient := github.NewClient()

	user, err := githubClient.Users.Get(githubUsername)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user.Name)
	}
}
