package cmd

import (
	"testing"
	"context"
	"github.com/spf13/viper"
	"os"
)

func TestGetRepositories(t *testing.T){

	token := os.Getenv("TOKEN")
	viper.Set("github_token", token)
	ctx := context.Background()

	// Scenario 1: Valid Scenario
	t.Run("Valid Scenario", func(t *testing.T) {
		targetUser := "DelcioKelson"

		if res := len(getRepositories(targetUser,ctx)); res == 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})

	
	// Scenario 2: Empty Repository
	t.Run("Empty Result", func(t *testing.T) {
		targetUser := "Delibreezy"

		// Call the function
		if res := len(getRepositories(targetUser,ctx)); res != 0{
			t.Errorf("Expected different instead of %d", res)
		}
	})
	
}