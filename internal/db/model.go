package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Job struct {
	Company       string     `json:"company" db:"company"`
	Position      string     `json:"position" db:"position"`
	Status        JobStatus  `json:"status" db:"status"`
	Location      NullString `json:"location" db:"location"`
	SalaryRange   NullString `json:"salary_range" db:"salary_range"`
	JobPostingURL NullString `json:"job_posting_url" db:"job_posting_url"`
	AppliedAt     *time.Time `json:"applied_at" db:"applied_at"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
	ID            int        `json:"id" db:"id"`
}

func AddJob(sqliteDB *sql.DB, job *Job) {
	const createQuery = `INSERT INTO jobs
		(company, position, status, location, applied_at, salary_range, job_posting_url)
		VALUES
		(?, ?, ?, ?, ?, ?, ?);`

	defaultAppliedAt := time.Now()
	if job.AppliedAt == nil {
		job.AppliedAt = &defaultAppliedAt
	}

	result, err := sqliteDB.Exec(
		createQuery,
		job.Company,
		job.Position,
		job.Status,
		ToSQLValue(&job.Location),
		FormatDateTime(*job.AppliedAt, true),
		ToSQLValue(&job.SalaryRange),
		ToSQLValue(&job.JobPostingURL),
	)
	if err != nil {
		fmt.Println("Error in adding job", err)
		return
	}

	jobDBId, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error getting new job ID", err)
		return
	}
	job.ID = int(jobDBId)
	fmt.Printf("New job application (ID: %d) added, good luck!\n", jobDBId)
}

func DeleteJobByID(sqliteDB *sql.DB, jobID int) {
	const deleteQuery = `DELETE FROM jobs
		WHERE id = ?
		RETURNING company, position;`

	var companyName, position string
	err := sqliteDB.QueryRow(deleteQuery, jobID).Scan(&companyName, &position)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No job found with the provided ID")
			return
		}
		fmt.Println("Error deleting job:", err)
		return
	}

	fmt.Printf("Application for %s at %s (ID: %d) deleted\n", position, companyName, jobID)
}

func GetJobByID(sqliteDB *sql.DB, id int) (*Job, error) {
	const selectQuery = `SELECT
		id, company, position, status, location, salary_range, job_posting_url, applied_at, created_at, updated_at
		FROM jobs WHERE id = ?;`

	var job Job
	var appliedAt, createdAt, updatedAt string
	err := sqliteDB.QueryRow(selectQuery, id).Scan(
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
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	job.AppliedAt, _ = ParseDateTime(appliedAt, true)
	job.CreatedAt, _ = ParseDateTime(createdAt, false)
	job.UpdatedAt, _ = ParseDateTime(updatedAt, false)
	return &job, nil
}

func getJobs(sqliteDB *sql.DB, query string, params ...any) ([]*Job, error) {
	rows, err := sqliteDB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	jobs, err := FetchJobsFromRows(rows)
	return jobs, err
}

func GetAllJobs(sqliteDB *sql.DB) ([]*Job, error) {
	const selectQuery = `SELECT
		id, company, position, status, location, salary_range, job_posting_url, applied_at, created_at, updated_at
		FROM jobs;`
	jobs, err := getJobs(sqliteDB, selectQuery)
	return jobs, err
}

func GetJobsByStatus(sqliteDB *sql.DB, jobStatus JobStatus) ([]*Job, error) {
	const selectQuery = `SELECT
		id, company, position, status, location, salary_range, job_posting_url, applied_at, created_at, updated_at
		FROM jobs WHERE status = ? ORDER BY applied_at ASC;`
	jobs, err := getJobs(sqliteDB, selectQuery, jobStatus)
	return jobs, err
}

func GetJobsByDate(sqliteDB *sql.DB, before string, after string) ([]*Job, error) {
	beforeTime, err := ParseDateTime(before, true)
	if err != nil {
		return nil, err
	}
	afterTime, err := ParseDateTime(after, true)
	if err != nil {
		return nil, err
	}
	if afterTime.After(*beforeTime) {
		return nil, err
	}
	selectQuery := `SELECT
		id, company, position, status, location, salary_range, job_posting_url, applied_at, created_at, updated_at
		FROM jobs WHERE applied_at >= ? AND applied_at <= ?;`
	jobs, err := getJobs(sqliteDB, selectQuery, afterTime, beforeTime)
	return jobs, err
}

type UpdatedJobParams struct {
	Company       *string
	Position      *string
	Status        *JobStatus
	Location      *string
	SalaryRange   *string
	JobPostingURL *string
	AppliedAt     *time.Time
}

func ToSQLValue[T any](ptr *T) any {
	if ptr == nil {
		return nil
	}
	switch v := any(*ptr).(type) {
	case time.Time:
		return FormatDateTime(v, true)
	case string:
		if v == "" {
			return nil
		}
		return *ptr
	default:
		return *ptr
	}
}

func UpdateJob(sqliteDB *sql.DB, jobID int, updates UpdatedJobParams) (*Job, error) {
	const updateQuery = `UPDATE jobs
		SET
		company = COALESCE(?, company),
		position = COALESCE(?, position),
		status = COALESCE(?, status),
		location = COALESCE(?, location),
		salary_range = COALESCE(?, salary_range),
		job_posting_url = COALESCE(?, job_posting_url),
		applied_at = COALESCE(?, applied_at),
		updated_at = (CURRENT_TIMESTAMP)
		WHERE id = ?
		RETURNING id, company, position, status, location, salary_range, job_posting_url, applied_at, created_at, updated_at;`

	row := sqliteDB.QueryRow(
		updateQuery,
		ToSQLValue(updates.Company),
		ToSQLValue(updates.Position),
		ToSQLValue(updates.Status),
		ToSQLValue(updates.Location),
		ToSQLValue(updates.SalaryRange),
		ToSQLValue(updates.JobPostingURL),
		ToSQLValue(updates.AppliedAt),
		jobID,
	)
	job, err := ParseRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No job found with that id")
			return nil, nil
		}
		return nil, err
	}
	return job, nil
}
