package loan

import (
	"log"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

//logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	loanApp := &types.LoanApplication{}
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

	result, err := stmt.Exec(id, loanId, borrowerId, loanApp.Type, loanApp.Term, loanApp.Amount, loanApp.Purpose)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
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

func GetAllLoanRequest(w http.ResponseWriter, r *http.Request) {
	rows := helper.Query("SELECT * FROM applications", w)

	// slice to store all loan applications
	var loanApplications []types.LoanApplication
	// process query
	for rows.Next() {
		var loan types.LoanApplication
		err := rows.Scan(&loan.ID, &loan.LoanID, &loan.BorrowerId, &loan.Type, &loan.Term, &loan.Amount, &loan.Purpose)
		if err != nil {
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting loans", nil)
			return
		}

		loanApplications = append(loanApplications, loan)
	}

	// Form response object
	res := map[string]interface{}{
		"loans": loanApplications,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "Loan application successfully created", res)
}
