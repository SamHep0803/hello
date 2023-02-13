package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/samhep0803/hello/internal/creds"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
)

var cfgFile string

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
	_, err := creds.GetCreds()
	if err == nil {
		fmt.Println("GitHub Already Initialized!")
	}

	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			fmt.Print("Enter GitHub Token: ")
			githubToken, _ := terminal.ReadPassword(0)
			err := creds.SetCreds("github", string(githubToken))
			if err != nil {
				fmt.Printf("failed to set github token: %s", err)
			}
		} else {
			fmt.Printf("failed to get github token: %s\n", err)
		}
	}
}

func checkWeatherExists() {
	_, err := creds.GetCreds()
	if err == nil {
		fmt.Println("Weather Already Initialized!")
	}

	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			fmt.Print("Enter OpenWeatherAPI Token: ")
			weatherToken, _ := terminal.ReadPassword(0)
			err := creds.SetCreds("weather", string(weatherToken))
			if err != nil {
				fmt.Printf("failed to set weather token: %s", err)
			}
		} else {
			fmt.Printf("failed to get weather token: %s\n", err)
		}
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find config directory.
		config, err := os.UserConfigDir()
		cobra.CheckErr(err)

		if _, err := os.Stat(config + "/hello"); os.IsNotExist(err) {
			err := os.Mkdir(config+"/hello", os.ModePerm)
			if err != nil {
				fmt.Sprintln("Error creating config file:", err)
			}
		}

		viper.AddConfigPath(config + "/hello")
		viper.SetConfigType("yaml")
		viper.SetConfigName("hello")

		// populate config
		viper.SetDefault("weather.cityId", "")

		viper.SafeWriteConfig()
	}
	viper.AutomaticEnv() // read in environment variables that match
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	initCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config/hello/hello.yaml)")

	rootCmd.AddCommand(initCmd)
}
