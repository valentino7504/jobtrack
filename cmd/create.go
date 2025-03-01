package cmd

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func optionalSQL(param string) sql.NullString {
	var paramSQL sql.NullString
	if param != "" {
		paramSQL.String = param
		paramSQL.Valid = true
	}
	return paramSQL
}

func initializeJob(cmd *cobra.Command) *db.Job {
	caser := cases.Title(language.English)
	company, _ := cmd.Flags().GetString("company")
	position, _ := cmd.Flags().GetString("position")
	status, _ := cmd.Flags().GetString("status")
	status = caser.String(status)
	location, _ := cmd.Flags().GetString("location")
	salaryRange, _ := cmd.Flags().GetString("salary-range")
	jobPostingURL, _ := cmd.Flags().GetString("job-posting-url")
	applied, _ := cmd.Flags().GetString("applied")
	appliedAt, err := db.ParseDateTime(applied, true)
	if err != nil {
		fmt.Println("Applied date is not in the right format")
		return nil
	}
	if appliedAt.After(time.Now()) {
		fmt.Println("Applied date cannot be in the future")
		return nil
	}
	if company == "" {
		fmt.Println("Company not specified")
		return nil
	}
	if position == "" {
		fmt.Println("Position not specified")
		return nil
	}
	if !db.IsValidStatus(db.JobStatus(status)) {
		fmt.Println("Specified status is not valid")
		fmt.Println(
			"Valid statuses are: Applied, Interview, Offer, Accepted, \"Rejected Offer\" and Rejected",
		)
		return nil
	}
	job := db.Job{
		Company:       company,
		Position:      position,
		Status:        db.JobStatus(status),
		Location:      optionalSQL(location),
		SalaryRange:   optionalSQL(salaryRange),
		JobPostingURL: optionalSQL(jobPostingURL),
		AppliedAt:     appliedAt,
	}
	return &job
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create job applications in the database with optional details.",
	Long: `Creates a new job. The mandatory flags are --company and --position
	Other flags include:
	- --status - To define the status of the application: Applied, Interview, Offer,
	Accepted, Rejected Offer, Rejected
	- --location - To define the location of the company where the job is
	- --applied - Specify the date of application in YYYY-MM-DD format
	- --salary-range - Specify the salary
	- --job-posting-url - The URL to the job posting`,
	Run: func(cmd *cobra.Command, args []string) {
		job := initializeJob(cmd)
		if job == nil {
			return
		}
		db.AddJob(SqliteDB, job)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String(
		"company",
		"",
		"Specify the name of the company where the job is",
	)
	createCmd.Flags().String(
		"position",
		"",
		"Specify the position you are applying to",
	)
	createCmd.Flags().String(
		"status",
		"applied",
		"Specify the stage of the hiring process you are at",
	)
	createCmd.Flags().String("location", "", "The location of the job")
	createCmd.Flags().String("salary-range", "", "The salary range of the job")
	createCmd.Flags().String("job-posting-url", "", "The URL of the job posting")
	createCmd.Flags().String(
		"applied",
		time.Now().Format("2006-01-02"),
		"The date of the application formatted YYYY-MM-DD",
	)
}
