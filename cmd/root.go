package cmd

import (
	"database/sql"
	"os"

	"github.com/spf13/cobra"
)

// A pointer to the SqliteDB handler
var SqliteDB *sql.DB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jobtrack",
	Short: "A CLI tool built in Go to track job applications efficiently.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jobtrack.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Sets the application db to the Sqlite DB handle.
func SetDB(dbHandle *sql.DB) {
	SqliteDB = dbHandle
}
