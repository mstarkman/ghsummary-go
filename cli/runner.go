package cli

import (
	"fmt"

	"github.com/mstarkman/ghsummary-go/github"
)

func Run(githubUsername *string) {
	githubClient := github.NewClient()

	user, err := githubClient.Users.Get(githubUsername)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("Load time: %f secs.", user.LoadTime))
		fmt.Println(user.Name)
		fmt.Println(user.Login)
		fmt.Println(user.PublicRepoCount)
		for _, r := range *user.Repos {
			fmt.Println(r.Name)
		}
	}
}
