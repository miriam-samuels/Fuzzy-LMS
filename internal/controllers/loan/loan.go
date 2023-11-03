package loan

import (
	"log"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

//logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	application := &types.Borrower{}
	helper.ParseRequestBody(w, r, application)

	// TODO: Validate request body

	// generate ID for borrower
	id := helper.GenerateUUID()
	loanId := helper.GenerateLoanID()

	// prepare query statement to insert new borrower into db
	stmt, err := database.LoanDb.Prepare("INSERT INTO borrowers (id, loanId, firstname, lastname, email, phone, birth_date, gender, nationality, state_origin, address, passport, signature) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)")
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	defer stmt.Close()

	//execute statement
	result, err := stmt.Exec(id, loanId, application.FirstName, application.LastName, application.Email, application.Phone, application.BirthDate, application.Gender, application.Nationality, application.StateOrigin, application.Address, application.Passport, application.Signature)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	//  prepare query statement to create loan application in db
	stmt, err = database.LoanDb.Prepare("INSERT INTO applications (loanId, borrowerId) VALUES ($1, $2)")
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	defer stmt.Close()

	result, err = stmt.Exec(loanId, id)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	// Form response object
	res := map[string]interface{}{
		"id":          id,
		"loanId":      loanId,
		"application": application,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Loan application successfully created", res)
	log.Println(result)
}

func GetAllLoanRequest(w http.ResponseWriter, r *http.Request) {
	rows, err := database.LoanDb.Query("SELECT * FROM applications")
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting loans", nil)
		return
	}

	var loans types.LoanApplication
	// process query
	for rows.Next() {
		err = rows.Scan(&loans.LoanID, &loans.BorrowerId)
		if err != nil {
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting loans", nil)
			return
		}
	}

	// row := database.LoanDb.QueryRow("SELECT * FROM borrowers WHERE loanId=$1", loans.LoanID)

	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting loans", nil)
		return
	}

	// Form response object
	res := map[string]interface{}{
		"loans": loans,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Loan application successfully created", res)
}
