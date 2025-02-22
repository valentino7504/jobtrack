package db

import (
	"database/sql"
	"errors"
	"time"

	_ "modernc.org/sqlite"
)

// JobStatus type
type JobStatus string

const (
	APPLIED        JobStatus = "Applied"
	INTERVIEW      JobStatus = "Interview"
	OFFER          JobStatus = "Offer"
	ACCEPTED       JobStatus = "Accepted"
	REJECTED_OFFER JobStatus = "Rejected Offer"
	REJECTED       JobStatus = "Rejected"
)

// Parses and returns a time.Time pointer from the format string
func ParseDateTime(strTime string, dateOnly bool) (*time.Time, error) {
	if strTime == "" {
		return nil, errors.New("No time string passed")
	}
	var format string
	if dateOnly {
		format = time.DateOnly
	} else {
		format = time.DateTime
	}
	t, err := time.Parse(format, strTime)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// Formats time.Time pointer to SQLite datetime
func FormatDateTime(t time.Time, dateOnly bool) string {
	var format string
	if dateOnly {
		format = time.DateOnly
	} else {
		format = time.DateTime
	}
	return t.Format(format)
}

// Parses a row and creates a job struct
func ParseRow(row *sql.Row) (*Job, error) {
	var job Job
	var appliedAt, createdAt, updatedAt string
	err := row.Scan(
		&job.ID,
		&job.Company,
		&job.Position,
		&job.Status,
		&job.Location,
		&job.SalaryRange,
		&job.JobPostingURL,
		&appliedAt,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}
	job.AppliedAt, _ = ParseDateTime(appliedAt, true)
	job.CreatedAt, _ = ParseDateTime(createdAt, false)
	job.UpdatedAt, _ = ParseDateTime(updatedAt, false)
	return &job, nil
}

// Extracts job structs from sql Rows
func FetchJobsFromRows(rows *sql.Rows) ([]*Job, error) {
	var jobs []*Job
	for rows.Next() {
		var job Job
		var appliedAt, createdAt, updatedAt string
		err := rows.Scan(
			&job.ID,
			&job.Company,
			&job.Position,
			&job.Status,
			&job.Location,
			&job.SalaryRange,
			&job.JobPostingURL,
			&appliedAt,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return jobs, err
		}
		job.AppliedAt, _ = ParseDateTime(appliedAt, true)
		job.CreatedAt, _ = ParseDateTime(createdAt, false)
		job.UpdatedAt, _ = ParseDateTime(updatedAt, false)
		jobs = append(jobs, &job)
	}
	if err := rows.Err(); err != nil {
		return jobs, err
	}
	return jobs, nil
}
