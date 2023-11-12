package loan

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
)

func (loanApp *Loan) CreateLoan(id uuid.UUID, loanId string, borrowerId string, w http.ResponseWriter) (*sql.Stmt, error) {
	//  prepare query statement to create loan application in db
	stmt, err := database.LoanDb.Prepare("INSERT INTO applications (id,loanId, borrowerId,type,term,amount,purpose, has_collateral, collateral, collateral_doc, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	return stmt, err
}

func GetLoans(userId string, userRole string, statusCondition string, w http.ResponseWriter) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	if userRole == "borrower" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications WHERE borrowerId = $1"+statusCondition, userId)
	} else if userRole == "lender" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications" + statusCondition)
	}

	return rows, err
}
