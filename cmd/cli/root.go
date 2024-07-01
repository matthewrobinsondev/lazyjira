package cli

import (
	"fmt"
	"github.com/matthewrobinsondev/lazyjira/pkg/handlers"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lazyjira",
	Short: "For devs who put code before POs",
	RunE:  handlers.NewLazyJiraHandler,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
