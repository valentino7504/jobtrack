package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export job applications as JSON or CSV to a file or standard output.",
	Long: `Export job applications from the database to a file or standard output.

You can choose between JSON (default) and CSV formats using the --format flag.
Use --output to specify a file instead of printing to stdout.

Examples:
  jobtrack export --format json --output jobs.json   # Save jobs as JSON
  jobtrack export --format csv --output jobs.csv     # Save jobs as CSV
  jobtrack export                                    # Print JSON to stdout
  jobtrack export --format csv                       # Print CSV to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		exportFormat, _ := cmd.Flags().GetString("format")
		filename, _ := cmd.Flags().GetString("output")
		jobs, err := db.GetAllJobs(SqliteDB, true)
		if err != nil {
			fmt.Println("Error getting jobs:", err)
			return
		}
		if len(jobs) == 0 {
			fmt.Println("No job applications available")
			return
		}
		var f *os.File

		if filename == "" {
			f = os.Stdout
		} else {
			f, err = os.Create(filename)
			if err != nil {
				fmt.Println("Error creating export file:", err)
				return
			}
		}
		switch exportFormat {
		case "csv":
			// use db.Jobs to use the ToCSV method on it
			w := csv.NewWriter(f)
			err = w.WriteAll(db.Jobs(jobs).ToCSV())
			if err != nil {
				fmt.Println("Error writing to CSV file:", err)
				return
			}
		default:
			b, err := json.Marshal(jobs)
			if err != nil {
				fmt.Println("Error marshalling to JSON:", err)
				return
			}
			var out bytes.Buffer
			err = json.Indent(&out, b, "", "\t")
			if err != nil {
				fmt.Println("Error marshalling to JSON:", err)
			}
			_, err = out.WriteTo(f)
			if err != nil {
				fmt.Println("Error writing JSON to file:", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP(
		"format",
		"f",
		"json",
		"Specify the format you want the export to be in - csv or json",
	)
	exportCmd.Flags().StringP(
		"output",
		"o",
		"",
		"Specify the output file (leave empty to print to stdout)",
	)
}
