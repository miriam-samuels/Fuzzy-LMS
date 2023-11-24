package auth

import (
	"database/sql"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/user"
)

// method to check if user exists in the db
func (cred *SignUpCred) CheckUser() (user.User, error) {
	// TODO: Fix error when user does exist
	exists := user.User{}
	err := database.LoanDb.QueryRow("SELECT id FROM users WHERE email=$1", cred.Email).Scan(&exists.ID)
	return exists, err
}

// finc user by email
func (cred *SignInCred) FindUserByMail() sql.Row {
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, password, role FROM users WHERE email= $1", cred.Email)
	return *row
}

// method to create a new user
func (cred *SignUpCred) CreateUser() (*sql.Stmt, error) {
	// prepare query statement to insert new user into db -- this table is mostly used for authentication
	stmt, err := database.LoanDb.Prepare("INSERT INTO users (id, firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)")
	return stmt, err
}

// method to create a borrower if user role is borrower
func (cred *SignUpCred) CreateBorrower() (*sql.Stmt, error) {
	stmt, err := database.LoanDb.Prepare("INSERT INTO borrowers (id, firstname, lastname, email) VALUES ($1, $2, $3, $4)")
	return stmt, err
}
