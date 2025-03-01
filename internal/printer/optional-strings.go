package printer

import (
	"database/sql"
)

func OptionalParamStr(param sql.NullString) string {
	if param.Valid {
		return param.String
	}
	return "N/A"
}
