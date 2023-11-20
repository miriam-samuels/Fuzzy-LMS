package loan

import (
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/loan"
)

// logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	loanApp := &loan.Loan{}

	err := helper.ParseRequestBody(w, r, loanApp)
	if err != nil {
		return
	}

	// TODO: Validate request body

	// generate UUID for loan
	id := helper.GenerateUUID()

	// generate easily identifiable id for loan
	loanId := "Loan" + helper.GenerateUniqueId(6)

	// get borrower id from request context
	borrowerId := r.Context().Value("userId").(string)

	// create loan application
	stmt, err := loanApp.CreateLoan(id, loanId, borrowerId, w)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		id,
		loanId,
		borrowerId,
		loanApp.Type,
		loanApp.Term,
		loanApp.Amount,
		loanApp.Purpose,
		loanApp.HasCollateral,
		loanApp.Collateral,
		loanApp.CollateralDocs,
		loanApp.Status,
	)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
		return
	}

	

	// Form response object
	res := map[string]interface{}{
		"id":     id,
		"loanId": loanId,
		"data":   loanApp,
	}
	helper.SendResponse(w, http.StatusOK, true, "Loan application successfully created", res)

}

func GetLoans(w http.ResponseWriter, r *http.Request) {
	// get loan status query parameter (status)
	status := r.URL.Query().Get("status")

	// get id of user making request
	userId := r.Context().Value("userId").(string)

	// get role of user making request
	userRole := r.Context().Value("userRole").(string)

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

	// get loans
	rows, err := loan.GetLoans(userId, userRole, statusCondition, w)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// slice to store all loan applications
	var loans []loan.Loan
	// process query
	for rows.Next() {
		var loan loan.Loan
		err := rows.Scan(&loan.LoanID, &loan.ID, &loan.BorrowerId, &loan.Type, &loan.Term, &loan.Amount, &loan.Purpose, &loan.Status, &loan.HasCollateral, &loan.CollateralDocs, &loan.Collateral)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error getting loans", nil, err)
			return
		}

		loans = append(loans, loan)
	}

	// Form response object
	res := map[string]interface{}{
		"loans": loans,
	}
	helper.SendResponse(w, http.StatusOK, true, "Loan application successfully created", res)
}
