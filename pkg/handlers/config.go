package handlers

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/matthewrobinsondev/lazyjira/pkg/config"
	"github.com/matthewrobinsondev/lazyjira/pkg/validate"
	"github.com/spf13/cobra"
)

func NewConfigAuthHandlerNewLazyJiraHandler(cmd *cobra.Command, args []string) error {
	configBuf := &config.Config{}

	var shouldOverride bool

	confSer := config.NewConfigService()

	confExists := confSer.Exists()

	fields := []huh.Field{
		huh.NewInput().
			Title("Whats you JIRA instance URL?").
			Value(&configBuf.JiraURL).
			Validate(func(val string) error {
				return validate.IsValidUrl(val)
			}),
		huh.NewInput().
			Title("Jira email address?").
			Value(&configBuf.Email),
		huh.NewInput().
			Title("Personal access token?").
			Value(&configBuf.AccessToken),
	}

	if confExists {
		fields = append(
			fields,
			huh.NewConfirm().
				Title("Should override existing?").
				Affirmative("Yes").
				Negative("No").
				Value(&shouldOverride),
		)
	}

	configForm := huh.NewForm(
		huh.NewGroup(fields...),
	).WithTheme(huh.ThemeBase())

	err := configForm.Run()
	if err != nil {
		return err
	}

	if confExists && !shouldOverride {
		fmt.Println("Skipping override")
		return nil
	}

	if err = confSer.Save(*configBuf); err != nil {
		cobra.CheckErr(err)
	}

	return nil
}
