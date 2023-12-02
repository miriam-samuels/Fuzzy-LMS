package dashboard

import (
	"database/sql"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func GetApplicationsCount(currentUser types.AuthCtxKey, status string) *sql.Row {
	var row *sql.Row
	if currentUser.Role == "borrower" {
		row = database.LoanDb.QueryRow("SELECT COUNT(*) FROM applications WHERE status = $1 AND borrowerId = $2", status, currentUser.Id)
	} else {
		row = database.LoanDb.QueryRow("SELECT COUNT(*) FROM applications WHERE status = $1", status)
	}
	return row
}
func GetWeeklyPlotPoints() (*sql.Rows, error) {
	query := `
    SELECT
        DATE_TRUNC('day', created_at) AS day,
        COUNT(*) AS applications
    FROM
        applications
    WHERE
        created_at >= current_date - interval '7 days'
    GROUP BY
        day
    ORDER BY
        day
`

	rows, err := database.LoanDb.Query(query)
	return rows, err
}

func GetMonthlyPlotPoints() (*sql.Rows, error) {
	query := `
    SELECT
        DATE_TRUNC('month', created_at) AS month,
        COUNT(*) AS applications
    FROM
        applications
    WHERE
        created_at >= current_date - interval '12 months'
    GROUP BY
	 	month
    ORDER BY
	 	month
`

	rows, err := database.LoanDb.Query(query)
	return rows, err
}

func GetPlotPoints5Yrs() (*sql.Rows, error) {
	query := `
    SELECT
        DATE_TRUNC('year', created_at) AS year,
        COUNT(*) AS applications
    FROM
        applications
    WHERE
        created_at >= current_date - interval '5 years'
    GROUP BY
	 	year
    ORDER BY
	 	year
`

	rows, err := database.LoanDb.Query(query)
	return rows, err
}
