package dastco

import (
	"fmt"

	webinspect "github.com/mkorejo/dastco/pkg/webinspect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// webInspectGetScanStatusCmd
	scanID string
)

// Initialize `wi` command
var webInspectCmd = &cobra.Command{
	Use:   "wi",
	Short: "Work with Fortify WebInspect",
	Long:  "Fortify WebInspect CLI to list/start/stop/modify scans, watch for scan completion, generate reports, and send findings to Jira",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Add `wi` command to rootCmd and configure flags
func init() {
	rootCmd.AddCommand(webInspectCmd)

	webInspectCmd.AddCommand(webInspectListScansCmd)
	webInspectCmd.AddCommand(webInspectGetScanStatusCmd)

	webInspectGetScanStatusCmd.Flags().StringVar(&scanID, "scan-id", "", "Scan ID")
	// gkeInspectCmd.Flags().BoolVar(&skipVault, "skip-vault", false, "Skip Vault during cluster inspection")
}

// Subcommands for `wi`
var webInspectGetScanStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of a scan",
	Run: func(cmd *cobra.Command, args []string) {
		readConfig()
		if scanID == "" {
			log.Fatal("You must specify --scan-id.")
		}
		c := createHTTPClient()
		status := webinspect.GetScanStatus(c, url, username, password, scanID)
		fmt.Println(status)
	},
}

var webInspectListScansCmd = &cobra.Command{
	Use:   "list",
	Short: "List scans",
	Long:  "List Fortify WebInspect scans including name, ID, and current state",
	Run: func(cmd *cobra.Command, args []string) {
		readConfig()
		c := createHTTPClient()
		list := webinspect.ListScans(c, url, username, password)
		fmt.Println(list)
	},
}

var webInspectStartScanCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a scan",
	Run: func(cmd *cobra.Command, args []string) {
		readConfig()
		if scanID == "" {
			log.Fatal("You must specify --scan-id.")
		}
		c := createHTTPClient()
		list := webinspect.ListScans(c, url, username, password)
		fmt.Println(list)
	},
}

// readConfig initializes the WebInspect URL and username/password variables.
func readConfig() {
	if url == "" {
		url = viper.GetString("url")
		if url == "" {
			log.Fatal("WebInspect API URL not found.")
		}
	}
	if username == "" {
		username = viper.GetString("username")
		if username == "" {
			log.Fatal("WebInspect username not found.")
		}
	}
	// Do not read `password` from a configuration file.
	if password == "" {
		log.Fatal("WebInspect password not found.")
	}
}
