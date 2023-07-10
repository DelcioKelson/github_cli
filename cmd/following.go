/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"context"
	"github.com/spf13/cobra"
	"github.com/google/go-github/github"	
)

// followsRepoCmd represents the followsRepo command
var followsRepoCmd = &cobra.Command{
	Use:   "following",
	Short: "Get repos followed by the target user",
	Long: `The "following" command fetches a list of repositories that are followed by
	the target user on GitHub. It provides insights into the repositories the user finds
	interesting or wants to keep track of
Usage:
github_cli following [target user]
	
Example:
github_cli following delciokelson
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		repos := getFollowedRepositories(args[0],ctx)
		fmt.Println("Repositories followed by the Target user:")
		printRepositories(repos)
	},
}

func init() {
	rootCmd.AddCommand(followsRepoCmd)
}

func getFollowedRepositories(targetUser string, ctx context.Context) ([]*github.Repository){

	client := getClient(ctx)

	repos, _, err := client.Activity.ListWatched(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving followed repositories: %v\n", err)
		return nil
	}

	return repos
}