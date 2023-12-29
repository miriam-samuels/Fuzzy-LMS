package borrower

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/user"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func GetBorrowers(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey)

	// check if user making request is a lender
	if currentUser.Role != "lender" {
		helper.SendResponse(w, http.StatusUnauthorized, false, "You can't perform this action", nil)
		return
	}

	var brws []user.Borrower
	rows, err := user.GetBorrowers()
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
		return
	}

	for rows.Next() {
		var brw user.Borrower

		err := rows.Scan(
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
			pq.Array(make([]string, 2)),
			pq.Array(make([]string, 2)),
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

		brws = append(brws, brw)
	}
	// Form response object
	res := map[string]interface{}{
		"user": brws,
	}
	helper.SendResponse(w, http.StatusOK, true, "User successfully updated", res)
}

func GetBorrowerById(w http.ResponseWriter, r *http.Request) {

	// get id and role of user making request
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey)

	// check if user making request is a lender
	if currentUser.Role != "lender" {
		helper.SendResponse(w, http.StatusUnauthorized, false, "You can't perform this action", nil)
		return
	}

	// variable to store borrower details
	var brw user.Borrower
	var kin []string
	var guarantor []string

	// get url paramenters from request url
	vars := mux.Vars(r)
	brw.ID = vars["id"]

	row := brw.FindBorrowerById()
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
		&brw.CreditScore)
	if err != nil {
		helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
		return
	}

	// get kins from kins table  using borrower id
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

	// get gurantors from table using borrower id
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
		"user": brw,
	}
	helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
}
