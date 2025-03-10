package jobPrinter

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/valentino7504/jobtrack/internal/db"
)

func PrintJobsTable(jobs []*db.Job) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintf(w, "ID\tCompany\tPosition\tStatus\tLocation\tSalary Range\tApplied On\n")
	for _, job := range jobs {
		location := db.NullString(job.Location)
		salaryRange := db.NullString(job.SalaryRange)
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\t%s\t%s\n",
			job.ID,
			job.Company,
			job.Position,
			job.Status,
			OptionalParamStr(location),
			OptionalParamStr(salaryRange),
			db.FormatDateTime(*job.AppliedAt, true),
		)
	}
	w.Flush()
}
