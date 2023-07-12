/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

// getReposCmd represents the getRepos command
var getReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Retrieve information about repositories owned by the target user",
	Long: `The repos command allows you to retrieve a list of
	repositories owned by the target user from GitHub. It fetches
	information about the repositories, including their names,
	descriptions, languages used, and the number of stars and
	forks they have received.

Usage:
github_cli repos [target user]
	
Example:
github_cli repos delciokelson
`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		repos := getRepositories(args[0], ctx)
		fmt.Println("Repositories owned by the Target user:")
		printRepositories(repos)
	},
}

func init() {
	rootCmd.AddCommand(getReposCmd)
}
