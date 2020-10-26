package dastco

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cliName = "dastco"
)

var (
	cfgFile, url, username, password string
	insecure                         bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   cliName,
	Short: "CLI for DAST",
	Long:  `Versatile CLI for programmatic interaction with Dynamic Application Security Testing (DAST) tools.`,

	// Default action
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute adds all child commands to the rootCmd and sets flags appropriately.
// This is called by main.main() and only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// PersistentFlags will work for this command and all subcommands.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Configuration file (defaults to ~/.dastco)")
	rootCmd.PersistentFlags().BoolVarP(&insecure, "insecure", "k", false, "Trust all SSL certificates")

	// Flags will only run when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in a configuration file and any ENV variables.
func initConfig() {
	if cfgFile != "" {
		// Use file specified by --config
		viper.SetConfigFile(cfgFile)
	} else {
		// Find $HOME directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search for a default configuration in $HOME directory with name ".dastco" (no extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dastco")
	}

	// Read in any environment variables that match
	viper.AutomaticEnv()

	// Read configuration file if one is found
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Warning("No configuration file specified.")
	}
}
