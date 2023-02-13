package cmd

import (
	"errors"
	"fmt"

	"github.com/samhep0803/hello/internal/creds"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkGithubExists()
		checkWeatherExists()
	},
}

func checkGithubExists() {
	_, err := creds.GetGithubToken()
	if err == nil {
		fmt.Printf("GitHub Already Initialized!")
	}

	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			fmt.Print("Enter GitHub Token: ")
			githubToken, _ := terminal.ReadPassword(0)
			err := creds.SetCreds("github", string(githubToken))
			if err != nil {
				fmt.Printf("failed to set github token: %s", err)
			}
		}
	}
}

func checkWeatherExists() {
	_, err := creds.GetWeatherToken()
	if err == nil {
		fmt.Printf("Weather Already Initialized!")
	}

	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			fmt.Print("Enter OpenWeatherAPI Token: ")
			weatherToken, _ := terminal.ReadPassword(0)
			err := creds.SetCreds("weather", string(weatherToken))
			if err != nil {
				fmt.Printf("failed to set weather token: %s", err)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
