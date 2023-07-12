package cmd

import (
	"context"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestGetPullRequests(t *testing.T) {

	// Get github secrect
	token := os.Getenv("TOKEN")
	viper.Set("github_token", token)
	ctx := context.Background()

	// Scenario 1: Valid Scenario
	// jjohannes is a user with some pull requests
	t.Run("Valid Scenario", func(t *testing.T) {
		targetUser := "jjohannes"

		repos := getRepositories(targetUser, ctx)

		// Call the function
		if res := len(getPullRequests(repos, targetUser, ctx)); res == 0 {
			t.Errorf("Expected different instead of %d", res)
		}
	})

	// Scenario 2: Empty Result
	t.Run("Empty Result", func(t *testing.T) {
		targetUser := "DelcioKelson"

		repos := getRepositories(targetUser, ctx)

		// Call the function
		if res := len(getPullRequests(repos, targetUser, ctx)); res != 0 {
			t.Errorf("Expected different instead of %d", res)
		}
	})
}
