package user

import (
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/model/v1/user"

	"github.com/lib/pq"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {

	// get id of user making request
	userId := r.Context().Value("userId").(string)

	// get role of user making request
	userRole := r.Context().Value("userRole").(string)

	if userRole == "borrower" {
		// variable to store borrower details
		var brw user.Borrower
		var kin []byte
		var guarantor []byte
		var offences []byte
		var loanIds []byte
		brw.ID = userId

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
			&offences,
			&brw.JailTime,
			&kin,
			&guarantor,
			&brw.NinSlip,
			&brw.Nin,
			&brw.Bvn,
			&brw.BankName,
			&brw.AccountNumber,
			&brw.Identification,
			&loanIds,
			&brw.Progress)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
			return
		}

		res := map[string]interface{}{
			"user": brw,
		}
		helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	} else {
		// variable to store user details
		var user user.User
		user.ID = userId

		// get user from db
		row := user.FindUserById()
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
			return
		}

		res := map[string]interface{}{
			"user": user,
		}
		helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	}

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// get id of user making request
	userId := r.Context().Value("userId").(string)

	// get role of user making request
	userRole := r.Context().Value("userRole").(string)

	if userRole == "borrower" {
		var user user.Borrower

		// parse request body into user
		helper.ParseRequestBody(w, r, &user)

		stmt, err := user.UpdateBorrower()
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
			return
		}

		defer stmt.Close()

		_, err = stmt.Exec(
			user.Phone,
			user.BirthDate,
			user.Gender,
			user.Nationality,
			user.StateOrigin,
			user.Address,
			user.Passport,
			user.Signature,
			user.Job,
			user.JobTerm,
			user.Income,
			user.Deck,
			user.HasCriminalRec,
			pq.Array(user.Offences),
			user.JailTime,
			pq.Array(user.Kin),
			pq.Array(user.Guarantor),
			user.NinSlip,
			user.Nin,
			user.Bvn,
			user.BankName,
			user.AccountNumber,
			user.Identification,
			pq.Array(user.LoanIds),
			user.Progress, userId)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db"+err.Error(), nil)
			return
		}
		// Form response object
		helper.SendResponse(w, http.StatusOK, true, "User successfully updated", nil)
	}

}
