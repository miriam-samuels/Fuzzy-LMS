package user

import (
	"encoding/json"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/repository/v1/user"
	"github.com/miriam-samuels/loan-management-backend/internal/types"

	"github.com/lib/pq"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {

	// get id and role of user making request
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey)

	if currentUser.Role == "borrower" {
		// variable to store borrower details
		var brw user.Borrower
		var kin []byte
		var guarantor []byte

		brw.ID = currentUser.Id

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

		res := map[string]interface{}{
			"user": brw,
		}
		helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	} else {
		// variable to store user details
		var user user.User
		user.ID = currentUser.Id

		// get user from db
		row := user.FindUserById()
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered::", nil, err)
			return
		}

		//form repsonse object
		res := map[string]interface{}{
			"user": user,
		}
		helper.SendResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	}

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// get id, role of user making request
	currentUser := r.Context().Value(types.AuthCtxKey{}).(types.AuthCtxKey)

	if currentUser.Role == "borrower" {
		var user user.Borrower

		// parse request body into user
		err := helper.ParseRequestBody(w, r, &user)
		if err != nil {
			helper.SendResponse(w, http.StatusBadRequest, false, "error parsing body:"+err.Error(), nil)
			return
		}

		// convert struct to json for kin
		kinJSON, err := json.Marshal(user.Kin)
		if err != nil {
			return
		}

		// convert struct to json for gurantor
		GuarantorJSON, err := json.Marshal(user.Guarantor)
		if err != nil {
			return
		}

		stmt, err := user.UpdateBorrower()
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
			return
		}

		defer stmt.Close()

		user.Progress = user.CalculateProgress()

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
			pq.Array(user.Offences), //concatenate struct to string seperated by comma
			user.JailTime,
			kinJSON,
			GuarantorJSON,
			user.Nin,
			user.Bvn,
			user.BankName,
			user.AccountNumber,
			user.Identification,
			user.Progress,
			currentUser.Id)
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db"+err.Error(), nil)
			return
		}
		// Form response object
		res := map[string]interface{}{
			"user": user,
		}
		helper.SendResponse(w, http.StatusOK, true, "User successfully updated", res)
	}

}
