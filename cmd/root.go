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
	Use:     "jobtrack",
	Short:   "A CLI tool built in Go to track job applications efficiently.",
	Version: "1.0.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

func SetDB(dbHandle *sql.DB) {
	SqliteDB = dbHandle
}
