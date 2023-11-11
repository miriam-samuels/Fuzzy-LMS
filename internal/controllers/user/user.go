package user

import (
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/database"
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

		row := database.LoanDb.QueryRow("SELECT * FROM borrowers WHERE id = $1", userId)
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
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered::"+err.Error(), nil)
			return
		}

		res := map[string]interface{}{
			"user": brw,
		}
		helper.SendJSONResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	} else {
		// variable to store user details
		var user user.User
		// get user from db
		row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, role FROM users WHERE id = $1", userId)
		err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Role)
		if err != nil {
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered::"+err.Error(), nil)
			return
		}

		res := map[string]interface{}{
			"user": user,
		}
		helper.SendJSONResponse(w, http.StatusOK, true, "Successfully fetched user profile", res)
	}

}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// get id of user making request
	userId := r.Context().Value("userId").(string)

	// get role of user making request
	userRole := r.Context().Value("userRole").(string)

	var user user.Borrower

	// parse request body into user
	helper.ParseRequestBody(w, r, &user)

	if userRole == "borrower" {
		query := "UPDATE borrowers SET phone = $1, birth_date = $2, gender = $3, nationality = $4, state_origin = $5, address = $6, passport = $7, signature = $8,  job = $9, job_term = $10, income = $11, deck = $12, has_criminal_record = $13, offences = $14, jail_time = $15, has_collateral = $16, collateral = $17, collateral_docs = $18, kin = $19, guarantor = $20, nin_slip = $21, nin = $22, bvn = $23, bank_name = $24, account = $25, identification = $26, loan_ids = $27, progress = $28 WHERE id = $29"

		stmt := helper.Prepare(query, w)

		defer stmt.Close()

		_, err := stmt.Exec(
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
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db"+err.Error(), nil)
			return
		}
		// Form response object
		helper.SendJSONResponse(w, http.StatusOK, true, "User successfully updated", nil)
	}

}
