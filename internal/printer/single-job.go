package printer

import (
	"fmt"

	"github.com/valentino7504/jobtrack/internal/db"
)

func PrintJob(job *db.Job) {
	var s string
	location := OptionalParamStr(job.Location)
	salaryRange := OptionalParamStr(job.SalaryRange)
	jobPostingURL := OptionalParamStr(job.JobPostingURL)
	s += fmt.Sprintf("Job ID: %d\nCompany: %s\nPosition: %s\n", job.ID, job.Company, job.Position)
	s += fmt.Sprintf("Status: %s\nLocation: %s\n", job.Status, location)
	s += fmt.Sprintf("Applied On: %s\n", db.FormatDateTime(*job.AppliedAt, true))
	s += fmt.Sprintf("Salary Range: %s\nJob Posting: %s", salaryRange, jobPostingURL)
	fmt.Println(s)
}
