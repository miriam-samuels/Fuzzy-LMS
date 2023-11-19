package user

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
)

// finc borrower by id
func (brw *Borrower) FindBorrowerById() sql.Row {
	row := database.LoanDb.QueryRow("SELECT * FROM borrowers WHERE id = $1", brw.ID)
	return *row
}

func (brw *Borrower) UpdateBorrower() (*sql.Stmt, error) {
	query := "UPDATE borrowers SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8,  job = $9, job_term = $10, income = $11, deck = $12, has_criminal_record = $13, offences = $14, jail_time = $15, kin = $16, guarantor = $17, nin = $18, bvn = $19, bank_name = $20, account_number = $21, identification = $22, progress = $23 WHERE id = $24"
	stmt, err := database.LoanDb.Prepare(query)
	return stmt, err
}

// finc user by email
func (user *User) FindUserById() sql.Row {
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, role FROM users WHERE id = $1", user.ID)
	return *row
}

func (user *User) UpdateUser() (*sql.Stmt, error) {
	query := "UPDATE users SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8, identification = $9 WHERE id = $10"
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
		log.Println(key)

		// exclude non-compulsory fields
		if key != "ID" && key != "FirstName" && key != "LastName" && key != "Email" && key != "LoanIds" && key != "Progress" {
			// check if value is present
			switch field.Kind() {
			case reflect.String:
				if field.String() != "" {
					progress += 5 // add to user progress
					log.Println(progress)
				}
			case reflect.Ptr, reflect.Interface:
				if !field.IsNil() {
					progress += 5 // add to user progress
					log.Println(progress)
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if field.Uint() != 0 {
					progress += 5 // add to user progress
					log.Println(progress)
				}
			default:
				log.Println(progress)
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
