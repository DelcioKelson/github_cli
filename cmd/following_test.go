package cmd

import (
	"testing"
	"context"
	"github.com/spf13/viper"
	"os"

)

func TestGetFollowedRepositories(t *testing.T)  {

	token := os.Getenv("TOKEN")
	viper.Set("github_token", token)
	ctx := context.Background()

	// Scenario 1: Valid Scenario
	t.Run("Valid Scenario", func(t *testing.T) {
		targetUser := "DelcioKelson"

		if res := len(getFollowedRepositories(targetUser,ctx)); res == 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})

	/*
	// Scenario 2: Empty Result
	t.Run("Empty Result", func(t *testing.T) {
		targetUser := "DelcioKelson"

		// Call the function
		if res := len(getFollowedRepositories(targetUser,ctx)); res == 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})
	*/

}
