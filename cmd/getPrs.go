/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"

	"github.com/spf13/cobra"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)


// getPrsCmd represents the getPrs command
var getPrsCmd = &cobra.Command{
	Use:   "getPrs",
	Short: "Get Pull requests",
	Long: `given a github user allows me to find it's latest 5 PRs`,
	Run: func(cmd *cobra.Command, args []string) {
		getPrs(args[0]);
	},
}

func init() {
	rootCmd.AddCommand(getPrsCmd)
}

func getPrs(targetUser string){

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)


	// List all repositories for the target user
	repos, _, err := client.Repositories.List(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving repositories: %v\n", err)
		return
	}

	flag := 0;
	count := 0;
		// Iterate over the repositories
	for _, repo := range repos {
		// Get pull requests for each repository
		pullRequests, _, err := client.PullRequests.List(ctx, targetUser, *repo.Name, nil)
		if err != nil {
			fmt.Printf("Error retrieving pull requests for repository %s: %v\n", *repo.Name, err)
			continue
		}

		
		if len(pullRequests) != 0 {
			flag = 1
		}else{
			continue
		}

		// Iterate over the pull requests and print relevant information
		for _, pr := range pullRequests {
			fmt.Printf("Repository: %s\n", *repo.Name)
			fmt.Printf("Title: %s\n", pr.GetTitle())
			fmt.Printf("Number: %d\n", pr.GetNumber())
			fmt.Printf("State: %s\n", pr.GetState())
			fmt.Println("--------------")

			count++;
			if count==5 {
				break;
			}
		}
	}

	if flag == 0 {
		fmt.Print("No pull requests found for repository", "\n")
	}
}
	



