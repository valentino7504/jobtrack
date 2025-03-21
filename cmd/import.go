package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/valentino7504/jobtrack/internal/db"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import job applications from a JSON or CSV file (must be .json or .csv).",
	Long: `Import job applications into the database from a JSON or CSV file.

The file format is automatically detected based on the extension (.json or .csv).
The import process will assign new IDs, ensuring no duplicates based on ID.
If a job already exists (matching company, position, and applied date), it will be skipped.

Examples:
  jobtrack import jobs.json   # Import from a JSON file
  jobtrack import jobs.csv    # Import from a CSV file

Only .json and .csv files are supported.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("No file path provided")
			return
		}
		fp := args[0]
		ext := filepath.Ext(fp)
		success, failed := 0, 0
		switch ext {
		case ".json":
			data, err := os.ReadFile(fp)
			if err != nil {
				log.Fatal("Error opening file:", err)
			}
			var jobs db.Jobs
			err = json.Unmarshal(data, &jobs)
			if err != nil {
				log.Fatal("Error unmarshalling JSON:", err)
			}
			for _, job := range jobs {
				err := db.AddJob(SqliteDB, job)
				if err != nil {
					failed++
				} else {
					success++
				}
			}
		case ".csv":
			f, err := os.Open(fp)
			if err != nil {
				fmt.Println("Error opening file:", err)
				return
			}
			defer f.Close()
			r := csv.NewReader(f)
			jobs, err := db.FromCSV(r)
			if err != nil {
				fmt.Println("Error reading CSV file")
				return
			}
			for _, job := range jobs {
				db.AddJob(SqliteDB, job)
				if err != nil {
					failed++
				} else {
					success++
				}
			}
		default:
			fmt.Println("File format", ext, "not supported. Use json or csv")
			return
		}
		fmt.Println("Import from", fp, "complete:", success, "jobs added,", failed, "failed.")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
