package user

import (
	"database/sql"
	"reflect"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
)

// finc borrower by id
func (brw *Borrower) FindBorrowerById() sql.Row {
	row := database.LoanDb.QueryRow("SELECT * FROM borrowers WHERE id = $1", brw.ID)
	return *row
}

func (brw *Borrower) UpdateBorrower() (*sql.Stmt, error) {
	query := "UPDATE borrowers SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8,  job = $9, job_term = $10, income = $11, deck = $12, has_criminal_record = $13, offences = $14, jail_time = $15, kin = $16, guarantor = $17, nin = $18, bvn = $19, bank_name = $20, account_number = $21, identification = $22, progress = $23, credit_score= $24 WHERE id = $25"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

// finc user by email
func (user *User) FindUserById() *sql.Row {
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, role FROM users WHERE id = $1", user.ID)
	return row
}

func (user *User) UpdateUser() (*sql.Stmt, error) {
	query := "UPDATE users SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8, identification = $9 WHERE id = $10"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

// get all kins associated with borrower
func (brw *Borrower) GetBorrowerKins() (*sql.Rows, error) {
	query := "SELECT * FROM kins WHERE borrowerid = $1"
	rows, err := database.LoanDb.Query(query, brw.ID)
	return rows, err
}

// get all guarantors associated with borrower
func (brw *Borrower) GetBorrowerGuarantors() (*sql.Rows, error) {
	query := "SELECT * FROM guarantors WHERE borrowerid = $1"
	rows, err := database.LoanDb.Query(query, brw.ID)
	return rows, err
}

// method to create a new ukin
func (cred *NextOfKin) CreateKin() (*sql.Stmt, error) {
	// prepare query statement to insert new kin into db
	query := "INSERT INTO kins(id, borrowerid, firstname, lastname, email, phone, gender, relationship, address) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

func (user *NextOfKin) UpdateKin() (*sql.Stmt, error) {
	query := "UPDATE kins SET firstname=$1, lastname=$2, email=$3, phone=$4, gender=$5, relationship=$6, address=$7 WHERE id=$8 AND borrowerid=$9"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

// method to create a new guarantor
func (cred *Guarantor) CreateGuarantor() (*sql.Stmt, error) {
	// prepare query statement to insert new kin into db
	query := "INSERT INTO guarantors (id, borrowerid, firstname, lastname, email, phone, gender, nin, income, signature, address) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

func (user *Guarantor) UpdateGuarantor() (*sql.Stmt, error) {
	query := "UPDATE guarantors SET firstname=$1, lastname=$2, email=$3, phone=$4, gender=$5, nin=$6, income=$7, signature=$8, address=$9 WHERE id=$10 AND borrowerid=$11"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

func (brw *Borrower) CalculateProgress() uint8 {
	var progress uint8 = 0
	// Get the reflect.Value of the struct by dereferencing the pointer
	val := reflect.ValueOf(*brw)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		key := val.Type().Field(i).Name

		// exclude non-compulsory fields
		if key != "ID" && key != "FirstName" && key != "LastName" && key != "Email" && key != "LoanIds" && key != "Progress" {
			// check if value is present
			switch field.Kind() {
			case reflect.String:
				if field.String() != "" {
					progress += 5 // add to user progress
				}
			case reflect.Ptr, reflect.Interface:
				if !field.IsNil() {
					progress += 5 // add to user progress
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if field.Uint() != 0 {
					progress += 5 // add to user progress
				}
			}
		}
	}

	return progress
}

// func (kins NextOfKin) Value() (driver.Value, error) {
// 	return json.Marshal(kins)
// }

// func (kin *NextOfKin) Scan(value interface{}) error {
// 	bry, ok := value.([]byte)
// 	if !ok {
// 		return errors.New("type assertion to []byte failed")
// 	}
// 	return json.Unmarshal(bry, &kin)
// }
