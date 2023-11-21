package loan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lib/pq"
	fis "github.com/miriam-samuels/loan-management-backend/internal/fuzzy-engine"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/loan"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/user"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

const unauthorized string = "unauthorized"

func GetLoans(w http.ResponseWriter, r *http.Request) {
	// get loan status query parameter (status)
	status := r.URL.Query().Get("status")

	// get id of user making request
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey)

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
	rows, err := loan.GetLoans(currentUser, statusCondition, w)
	if err != nil {
		if err.Error() == unauthorized {
			helper.SendResponse(w, http.StatusUnauthorized, false, "can't view this information", nil, err)
			return
		}
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

// logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	application := &loan.Loan{}

	err := helper.ParseRequestBody(w, r, application)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil)
		return
	}

	// TODO: Validate request body

	// generate UUID for loan
	id := helper.GenerateUUID()

	// generate easily identifiable id for loan
	loanId := "Loan" + helper.GenerateUniqueId(6)

	var brw user.Borrower
	var kin []byte
	var guarantor []byte

	// get borrower id from request context
	brw.ID = r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey).Id

	// get borrower information
	row := brw.FindBorrowerById()
	err = row.Scan(
		&brw.ID,
		&brw.FirstName,
		&brw.LastName,
		&brw.Email,
		&brw.Phone,
		&brw.BirthDate,
		&brw.Gender,
		&brw.Nationality,
		&brw.StateOrigin,
		&brw.Address,
		&brw.Passport,
		&brw.Signature,
		&brw.Job,
		&brw.JobTerm,
		&brw.Income,
		&brw.Deck,
		&brw.HasCriminalRec,
		pq.Array(&brw.Offences),
		&brw.JailTime,
		&kin,
		&guarantor,
		&brw.Nin,
		&brw.Bvn,
		&brw.BankName,
		&brw.AccountNumber,
		&brw.Identification,
		pq.Array(&brw.LoanIds),
		&brw.Progress,
		&brw.CreditScore)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// Unmarshal kin JSON data into structs
	if err := json.Unmarshal(kin, &brw.Kin); err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// Unmarshal guarantor JSON data into structs
	if err := json.Unmarshal(guarantor, &brw.Guarantor); err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	//  access creditwothiness of application ... dereferenced application
	creditworthiness := fis.AccessCreditworthiness(brw, *application)
	fmt.Printf("User Creditworthiness :: %v", creditworthiness)

	// create loan application
	stmt, err := application.CreateLoan(id, loanId, brw.ID, w)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		id,
		loanId,
		brw.ID,
		application.Type,
		application.Term,
		application.Amount,
		application.Purpose,
		application.HasCollateral,
		application.Collateral,
		application.CollateralDocs,
		application.Status,
	)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
		return
	}

	// Form response object
	res := map[string]interface{}{
		"id":     id,
		"loanId": loanId,
		"data":   application,
	}
	helper.SendResponse(w, http.StatusOK, true, "Loan application successfully created", res)

}
