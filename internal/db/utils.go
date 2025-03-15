package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	_ "modernc.org/sqlite"
)

// NullString type
type NullString struct {
	sql.NullString
}

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

func IsValidStatus(status JobStatus) bool {
	caser := cases.Title(language.English)
	titleStatus := caser.String(string(status))
	validStatuses := map[JobStatus]struct{}{
		APPLIED:        {},
		INTERVIEW:      {},
		OFFER:          {},
		ACCEPTED:       {},
		REJECTED_OFFER: {},
		REJECTED:       {},
	}
	_, ok := validStatuses[JobStatus(titleStatus)]
	return ok
}

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
		return nil, errors.New("Time passed is not in the correct format")
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

// NullString implements MarshalJSON to ensure that the potentially null strings are marshalled
// properly.
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// Unmarshalling a NullString
func (ns *NullString) UnmarshalJSON(data []byte) error {
	ns.Valid = string(data) != "null"
	e := json.Unmarshal(data, &ns.String)
	return e
}

// Marshalling a NullString into CSV value
func nullToEmpty(ns NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// Unmarshalling a NullString from a CSV value
func emptyToNull(s string) NullString {
	var ns NullString
	if len(s) == 0 {
		ns.Valid = false
		ns.String = ""
		return ns
	}
	ns.Valid = true
	ns.String = s
	return ns
}
