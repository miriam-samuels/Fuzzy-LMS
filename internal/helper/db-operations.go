package helper

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
)

func Prepare(q string, w http.ResponseWriter) *sql.Stmt {
	stmt, err := database.LoanDb.Prepare(q)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		fmt.Printf("ERROR:: %v", err)
		return nil
	}

	return stmt
}

func Query(q string, w http.ResponseWriter) *sql.Rows {
	rows, err := database.LoanDb.Query(q)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, false, "error getting data", nil)
		return nil
	}
	return rows
}


