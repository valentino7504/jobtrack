package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
	"github.com/valentino7504/jobtrack/internal/jobPrinter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List job applications with optional filters and sorting.",
	Long: `Retrieve job applications from the database.

By default, this command lists all jobs. You can filter results by ID, status, or applied date,
and sort them by latest or oldest.

Examples:
  jobtrack list                           # List all job applications
  jobtrack list --status "Interview"      # List jobs with status "Interview"
  jobtrack list --latest                  # List jobs sorted by most recent first
  jobtrack list --after "2024-01-01"      # List jobs applied on or after Jan 1, 2024
`,
	Run: func(cmd *cobra.Command, args []string) {
		jobID, _ := cmd.Flags().GetInt("id")
		status, _ := cmd.Flags().GetString("status")
		after, _ := cmd.Flags().GetString("after")
		before, _ := cmd.Flags().GetString("before")
		caser := cases.Title(language.English)
		status = caser.String(status)

		switch {
		case jobID > -1:
			job, err := db.GetJobByID(SqliteDB, jobID)
			if err != nil {
				fmt.Println("Error getting job:", err)
				return
			}
			if job == nil {
				fmt.Println("No job found with ID:", jobID)
				return
			}
			jobPrinter.PrintJob(job)
		case status != "":
			jobs, err := db.GetJobsByStatus(SqliteDB, db.JobStatus(status))
			if err != nil {
				fmt.Println("Error getting jobs:", err)
				return
			}
			if len(jobs) == 0 {
				fmt.Println("No jobs found with that status")
				return
			}
			jobPrinter.PrintJobsTable(jobs)
		case cmd.Flags().Changed("after") || cmd.Flags().Changed("before"):
			jobs, err := db.GetJobsByDate(SqliteDB, before, after)
			if err != nil {
				fmt.Println(err)
				return
			}
			if len(jobs) == 0 {
				fmt.Println("No jobs found within the specified date range")
				return
			}
			jobPrinter.PrintJobsTable(jobs)
		default:
			jobs, err := db.GetAllJobs(SqliteDB, false)
			if err != nil {
				fmt.Println("Error getting jobs", err)
				return
			}
			if len(jobs) == 0 {
				fmt.Println("No job applications available")
				return
			}
			jobPrinter.PrintJobsTable(jobs)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Int("id", -1, "The integer index of the job")
	listCmd.Flags().String("status", "", "The status of the job")
	listCmd.Flags().String("after", "1970-01-01", "List jobs applied on or after this date")
	listCmd.Flags().String("before", db.FormatDateTime(time.Now(), true), "List jobs applied on or before this date")
}
