package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
	"github.com/valentino7504/jobtrack/internal/jobPrinter"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job application by its ID.",
	Long: `Remove a job application from the database using its unique ID.

This action is irreversible, so use it with caution.

Examples:
  jobtrack delete --id 3     # Deletes the job with ID 3
  jobtrack delete --id 10    # Deletes the job with ID 10
`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		if id == -1 {
			fmt.Println("Specify the id of the job you want to delete")
			return
		}
		job, err := db.GetJobByID(SqliteDB, id)
		if err != nil {
			fmt.Println("Error accessing job with id", id)
			return
		}
		if job == nil {
			fmt.Println("No job found with ID:", id)
			return
		}
		force, _ := cmd.Flags().GetBool("force")
		if force {
			db.DeleteJobByID(SqliteDB, id)
			return
		}
		fmt.Println("Job to be deleted:")
		jobPrinter.PrintJob(job)
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nAre you sure? ([Y]/n): ")
		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		if strings.ToLower(response) == "n" {
			return
		}
		db.DeleteJobByID(SqliteDB, id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
	deleteCmd.Flags().Int("id", -1, "Specify the ID of the job you want to delete")
}
