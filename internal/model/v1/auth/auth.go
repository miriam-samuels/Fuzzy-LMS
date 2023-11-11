package auth

import (
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
)

// method to check if user exists in the db
func (cred *SignUpCred) CheckUser(exists *bool, w http.ResponseWriter) {
	err := database.LoanDb.QueryRow("SELECT 1 FROM users WHERE email=$1", cred.Email).Scan(&exists)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered", nil)
		return
	}
}

// finc user by email
func (cred *SignInCred) FindUserByMail(w http.ResponseWriter) sql.Row {
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, password, role FROM users WHERE email= $1", cred.Email)
	return *row
}

// method to create a new user
func (cred *SignUpCred) CreateUser(userId uuid.UUID, encryptedPass string, w http.ResponseWriter) {
	// prepare query statement to insert new user into db -- this table is mostly used for authentication
	stmt, err := database.LoanDb.Prepare("INSERT INTO users (id, firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	defer stmt.Close()

	//execute statement
	_, err = stmt.Exec(userId, cred.FirstName, cred.LastName, cred.Email, encryptedPass, cred.Role)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}
}

// method to create a borrower if user role is borrower
func (cred *SignUpCred) CreateBorrower(userId uuid.UUID, w http.ResponseWriter) {
	stmt, err := database.LoanDb.Prepare("INSERT INTO borrowers (id, firstname, lastname, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(userId, cred.FirstName, cred.LastName, cred.Email)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}
}
