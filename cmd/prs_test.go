package cmd

import (
	"testing"
	"context"
	"github.com/spf13/viper"
	"os"
)

func TestGetPullRequests(t *testing.T){

	token := os.Getenv("TOKEN")
	viper.Set("github_token", token)
	ctx := context.Background()

	// Scenario 1: Valid Scenario

	t.Run("Valid Scenario", func(t *testing.T) {
		targetUser := "rybakov"

		repos := getRepositories(targetUser,ctx)

		if res := len(getPrs(repos,targetUser,ctx)); res != 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})
		
	// Scenario 2: Empty Result
	t.Run("Empty Result", func(t *testing.T) {
		targetUser := "DelcioKelson"

		repos := getRepositories(targetUser,ctx)

		if res := len(getPullRequests(repos,targetUser,ctx)); res != 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})
}