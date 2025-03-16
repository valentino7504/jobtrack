package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
)

func processParam(param string) *string {
	if param == "" {
		return nil
	}
	return &param
}

func initializeUpdate(cmd *cobra.Command) *db.UpdatedJobParams {
	company, _ := cmd.Flags().GetString("company")
	position, _ := cmd.Flags().GetString("position")
	status, _ := cmd.Flags().GetString("status")
	location, _ := cmd.Flags().GetString("location")
	salaryRange, _ := cmd.Flags().GetString("salary-range")
	jobPostingURL, _ := cmd.Flags().GetString("job-posting-url")
	applied, _ := cmd.Flags().GetString("applied")
	appliedAt, err := db.ParseDateTime(applied, true)
	if err != nil && applied != "" {
		fmt.Println(err)
		return nil
	}
	if !db.IsValidStatus(db.JobStatus(status)) && status != "" {
		fmt.Println("Specified status is not valid")
		fmt.Println(
			"Valid statuses are: Applied, Interview, Offer, Accepted, \"Rejected Offer\" and Rejected",
		)
		return nil
	}

	updatedParams := db.UpdatedJobParams{
		Company:       processParam(company),
		Position:      processParam(position),
		Status:        (*db.JobStatus)(processParam(status)),
		Location:      processParam(location),
		SalaryRange:   processParam(salaryRange),
		JobPostingURL: processParam(jobPostingURL),
		AppliedAt:     appliedAt,
	}
	return &updatedParams
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing job application by specifying its ID and new details.",
	Long: `Update a job application in the database using its unique ID.

You can update details such as company name, position, status, location, salary range, job posting URL,
or the date you applied. Only the fields you specify will be changed, leaving other details untouched.

Examples:
  jobtrack update --id 3 --status "Interview"
  jobtrack update --id 5 --company "Google" --position "Software Engineer"
  jobtrack update --id 2 --salary-range "$80,000 - $100,000"`,
	Run: func(cmd *cobra.Command, args []string) {
		jobID, _ := cmd.Flags().GetInt("id")
		if jobID == -1 {
			fmt.Println("Please provide a valid job id")
			return
		}
		updatedParams := initializeUpdate(cmd)
		if updatedParams == nil {
			return
		}
		job, err := db.UpdateJob(SqliteDB, jobID, *updatedParams)
		if err != nil {
			fmt.Println("Error updating job:", err)
			return
		}
		if job == nil {
			return
		}
		fmt.Println("Job with id:", job.ID, "has been updated")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Int(
		"id",
		-1,
		"Specify the ID of the job to be updated",
	)
	updateCmd.Flags().String(
		"company",
		"",
		"Specify the name of the company where the job is",
	)
	updateCmd.Flags().String(
		"position",
		"",
		"Specify the position you are applying to",
	)
	updateCmd.Flags().String(
		"status",
		"",
		"Specify the stage of the hiring process you are at",
	)
	updateCmd.Flags().String("location", "", "The location of the job")
	updateCmd.Flags().String("salary-range", "", "The salary range of the job")
	updateCmd.Flags().String("job-posting-url", "", "The URL of the job posting")
	updateCmd.Flags().String(
		"applied",
		"",
		"The date of the application formatted YYYY-MM-DD",
	)
}
