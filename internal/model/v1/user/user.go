package user

import (
	"database/sql"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
)

// finc borrower by id
func (brw *Borrower) FindBorrowerById() sql.Row {
	row := database.LoanDb.QueryRow("SELECT * FROM borrowers WHERE id = $1", brw.ID)
	return *row
}

func (brw *Borrower) UpdateBorrower() (*sql.Stmt, error) {
	query := "UPDATE borrowers SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8,  job = $9, job_term = $10, income = $11, deck = $12, has_criminal_record = $13, offences = $14, jail_time = $15, has_collateral = $16, collateral = $17, collateral_docs = $18, kin = $19, guarantor = $20, nin_slip = $21, nin = $22, bvn = $23, bank_name = $24, account = $25, identification = $26, loan_ids = $27, progress = $28 WHERE id = $29"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

// finc user by email
func (user *User) FindUserById() sql.Row {
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, role FROM users WHERE id = $1", user.ID)
	return *row
}

func (user *User) UpdateUser() (*sql.Stmt, error) {
	query := "UPDATE users SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8, identification = $9, loan_ids = $10 WHERE id = $11"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}
