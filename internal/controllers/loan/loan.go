package loan

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

//logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	loanApp := &types.Loan{}
	helper.ParseRequestBody(w, r, loanApp)

	// TODO: Validate request body

	// generate UUID for loan
	id := helper.GenerateUUID()

	// generate easily identifiable id for loan
	loanId := helper.GenerateLoanID()

	// get borrower id from request context
	borrowerId := r.Context().Value("userId").(string)

	//  prepare query statement to create loan application in db
	stmt := helper.Prepare("INSERT INTO applications (id,loanId, borrowerId,type,term,amount,purpose) VALUES ($1, $2, $3, $4, $5, $6, $7)", w)

	defer stmt.Close()

	result, err := stmt.Exec(id, loanId, borrowerId, loanApp.Type, loanApp.Term, loanApp.Amount, loanApp.Purpose)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		fmt.Printf("ERROR:: %v", err)
		return
	}

	// Form response object
	res := map[string]interface{}{
		"id":     id,
		"loanId": loanId,
		"data":   loanApp,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Loan application successfully created", res)
	log.Println(result)
}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	// get loan status query parameter (status)
	status := r.URL.Query().Get("status")

	// get id of user making request
	userId := r.Context().Value("userId").(string)

	// variable to store user
	var user types.User

	// get user details from db
	err := database.LoanDb.QueryRow("SELECT role FROM users WHERE id = $1", userId).Scan(&user.Role)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil)
		return
	}

	//set status condition for query based on request query params
	var statusCondition string
	switch status {
	case "pending":
		statusCondition = " AND status = 'pending'"
	case "reviewing":
		statusCondition = " AND status = 'reviewing'"
	case "approved":
		statusCondition = " AND status = 'approved'"
	case "declined":
		statusCondition = " AND status = 'declined'"
	default:
		statusCondition = ""
	}

	var rows *sql.Rows
	if user.Role == "borrower" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications WHERE borrowerId = $1"+statusCondition, userId)
	} else if user.Role == "lender" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications" + statusCondition)
	}
	
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil)
		return
	}

	// slice to store all loan applications
	var loans []types.Loan
	// process query
	for rows.Next() {
		var loan types.Loan
		err := rows.Scan(&loan.ID, &loan.LoanID, &loan.BorrowerId, &loan.Type, &loan.Term, &loan.Amount, &loan.Purpose, &loan.Status)
		if err != nil {
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting loans", nil)
			return
		}

		loans = append(loans, loan)
	}

	// Form response object
	res := map[string]interface{}{
		"loans": loans,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Loan application successfully created", res)
}
