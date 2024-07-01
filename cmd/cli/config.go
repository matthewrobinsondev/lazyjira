package cli

import (
	"github.com/matthewrobinsondev/lazyjira/pkg/handlers"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration and setup",
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Setup and authenticate to your Jira account via an Access Token",
	RunE:  handlers.NewConfigAuthHandlerNewLazyJiraHandler,
}

func init() {
	configCmd.AddCommand(authCmd)
}
