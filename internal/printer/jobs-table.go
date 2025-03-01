package printer

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/valentino7504/jobtrack/internal/db"
)

// location, salary_range, job posting
func PrintJobsTable(jobs []*db.Job) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintf(w, "ID\tCompany\tPosition\tStatus\tLocation\tSalary Range\tApplied On\n")
	for _, job := range jobs {
		// location := "N/A"
		// fmt.Println(job.Location.Valid)
		// if job.Location.Valid {
		// 	location = job.Location.String
		// }
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\t%s\t%s\n",
			job.ID,
			job.Company,
			job.Position,
			job.Status,
			OptionalParamStr(job.Location),
			OptionalParamStr(job.SalaryRange),
			db.FormatDateTime(*job.AppliedAt, true),
		)
	}
	w.Flush()
}
