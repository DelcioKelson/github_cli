package cmd

import (
	"fmt"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/spf13/viper"
	"golang.org/x/term"
	"os"
)

func getClient(ctx context.Context)(*github.Client){

	flag := 0

	for {
		if !viper.IsSet("github_token") {
			fmt.Println("Set you github token:")
			token, _ := term.ReadPassword(int(os.Stdin.Fd()))
			viper.Set("github_token", token)
			flag++
		}

		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: viper.GetViper().GetString("github_token")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client := github.NewClient(tc)

		_, _, err := client.Users.Get(ctx, "")

		if err != nil {
			_, ok := err.(*github.ErrorResponse) 
			if ok {
				fmt.Println("Invalid GitHub token. Please provide a valid token.")
				viper.Set("github_token", nil)
				continue
			}
		}	
		
		if(flag>1){
			err = viper.WriteConfig()
			if err != nil {
				  fmt.Printf("Error writing config file, %s", err)
			}
		}


		return client
	}
}

func printRepositories(repos []*github.Repository){
	// Iterate over the repositories and print their data
	for _, repo := range repos {
		fmt.Printf("Repository: %s\n",repo.GetFullName())
		fmt.Printf("Description: %s\n", repo.GetDescription())
		fmt.Printf("Language: %se\n", repo.GetLanguage())
		dateTime := repo.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Printf("Created at: %s\n",dateTime)
		fmt.Println("--------------")
	}
}

func getRepositories(targetUser string, ctx context.Context) ([]*github.Repository){

	client := getClient(ctx)

	// get all repositories for the target user
	repos, _, err := client.Repositories.List(ctx, targetUser, nil)
	if err != nil {
		fmt.Printf("Error retrieving repositories: %v\n", err)
		return nil
	}

	return repos
}

func print5PullRequests(pullRequests []*github.PullRequest){

	fmt.Println("Pull requests of Target user:")

	count := 0

	// Iterate over the pull requests and print relevant information
	for _, pr := range pullRequests{
		fmt.Printf("Repository: %s\n", pr.GetBase().GetRepo().GetFullName())
		fmt.Printf("Title: %s\n", pr.GetTitle())
		fmt.Printf("Number: %d\n", pr.GetNumber())
		fmt.Printf("State: %s\n", pr.GetState())

		dateTime := pr.CreatedAt.Format("2006-01-02 15:04:05")
		fmt.Printf("Created at: %s\n",dateTime)
		fmt.Println("--------------")

		count++
		if count==5{
			break;
		}
	}

	if count==0{
		fmt.Println("No pull requests")
	}
}
