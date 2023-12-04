package loan

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
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
	rows, err := loan.GetLoans(currentUser, statusCondition)
	if err != nil {
		if err.Error() == unauthorized {
			helper.SendResponse(w, http.StatusUnauthorized, false, "can't view this information", nil, err)
			return
		}
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// slice to store all loan applications
	loans := []loan.Loan{}
	// process query
	for rows.Next() {
		var loan loan.Loan
		err := rows.Scan(&loan.ID, &loan.LoanID, &loan.BorrowerId, &loan.Type, &loan.Term, &loan.Amount, &loan.Purpose, &loan.Status, &loan.Creditworthiness, &loan.HasCollateral, &loan.CollateralDocs, &loan.Collateral, &loan.CreatedAt)
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
	helper.SendResponse(w, http.StatusOK, true, "Loans fetched", res)
}

func GetLoanById(w http.ResponseWriter, r *http.Request) {
	// variable to store gotten loan
	var loan loan.Loan

	// get url paramenters from request url
	vars := mux.Vars(r)
	loan.ID = vars["id"]

	// get loan by id
	row := loan.GetLoanById()

	row.Scan(
		&loan.ID,
		&loan.LoanID,
		&loan.BorrowerId,
		&loan.Term,
		&loan.Type,
		&loan.Amount,
		&loan.Purpose,
		&loan.Status,
		&loan.Creditworthiness,
		&loan.HasCollateral,
		&loan.Collateral,
		&loan.CollateralDocs,
		&loan.CreatedAt,
	)

	// variable to store borrower details
	var brw user.Borrower
	var kin []string
	var guarantor []string

	brw.ID = loan.BorrowerId

	row = brw.FindBorrowerById()
	err := row.Scan(
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
		pq.Array(&kin),
		pq.Array(&guarantor),
		&brw.Nin,
		&brw.Bvn,
		&brw.BankName,
		&brw.AccountNumber,
		&brw.Identification,
		pq.Array(&brw.LoanIds),
		&brw.Progress,
		&brw.CreditScore,
	)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// get kins using borrower id
	if len(kin) > 0 {
		rows, err := brw.GetBorrowerKins()
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
			return
		}

		for rows.Next() {
			var kin user.NextOfKin
			err := rows.Scan(&kin.ID, &kin.BorrowerId, &kin.FirstName, &kin.LastName, &kin.Email, &kin.Phone, &kin.Gender, &kin.Relationship, &kin.Address)
			if err != nil {
				helper.SendResponse(w, http.StatusInternalServerError, false, "error getting kins", nil, err)
				return
			}
			brw.Kin = append(brw.Kin, kin)
		}
	}

	// get gurantors using borrower id
	if len(guarantor) > 0 {
		rows, err := brw.GetBorrowerGuarantors()
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
			return
		}

		for rows.Next() {
			var g user.Guarantor
			err := rows.Scan(&g.ID, &g.BorrowerId, &g.FirstName, &g.LastName, &g.Email, &g.Phone, &g.Gender, &g.Nin, &g.Income, &g.Signature, &g.Address)
			if err != nil {
				helper.SendResponse(w, http.StatusInternalServerError, false, "error getting kins", nil, err)
				return
			}
			brw.Guarantor = append(brw.Guarantor, g)
		}
	}

	res := map[string]interface{}{
		"loan": loan,
		"user": brw,
	}
	helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)

}

// logic to create a new Loan Application goes here
func CreateLoanApplication(w http.ResponseWriter, r *http.Request) {
	application := &loan.Loan{}

	err := helper.ParseRequestBody(w, r, application)
	if err != nil {
		helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil)
		return
	}

	// set status to pending by defaulr
	application.Status = "pending"

	// TODO: Validate request body

	// generate UUID for loan
	id := helper.GenerateUUID()

	// generate easily identifiable id for loan
	loanId := "Loan" + helper.GenerateUniqueId(6)

	var brw user.Borrower

	// get borrower id from request context
	brw.ID = r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey).Id

	// get borrower information
	row := brw.GetLoanDetails()
	err = row.Scan(
		&brw.CreditScore,
		&brw.Income,
		&brw.HasCriminalRec,
		&brw.JobTerm,
		pq.Array(&brw.Offences),
		&brw.Progress,
	)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// check user profile progres
	if brw.Progress < 90 {
		helper.SendResponse(w, http.StatusBadGateway, false, "Please complete profile", nil)
		return
	}

	//  access creditwothiness of application ... dereferenced application
	creditworthiness := fis.AccessCreditworthiness(brw, *application)
	fmt.Printf(" \n User Creditworthiness :: %v \n", creditworthiness)

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
