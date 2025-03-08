package jobPrinter

import "github.com/valentino7504/jobtrack/internal/db"

func OptionalParamStr(param db.NullString) string {
	if param.Valid {
		return param.String
	}
	return "N/A"
}
