package profile

import (
	"fmt"
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
		var kin []string
		var guarantor []string

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

		fmt.Printf("%v", len(user.Kin))

		// update kins table
		var kins []string
		if len(user.Kin) > 0 {
			for i, k := range user.Kin {
				// check if kin already has id
				if k.ID == "" {
					k.ID = helper.GenerateUUID().String()
					k.BorrowerId = currentUser.Id
					// edit original array ...
					user.Kin[i].ID = k.ID
					user.Kin[i].BorrowerId = k.BorrowerId
					// insert kin into table
					stmt, err := k.CreateKin()
					if err != nil {
						helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
						return
					}
					defer stmt.Close()

					// execute statement for either create or update
					_, err = stmt.Exec(
						k.ID,
						k.BorrowerId,
						k.FirstName,
						k.LastName,
						k.Email,
						k.Phone,
						k.Gender,
						k.Relationship,
						k.Address,
					)
					kins = append(kins, k.ID)
				} else {
					// update
					stmt, err := k.UpdateKin()
					if err != nil {
						helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
						return
					}
					defer stmt.Close()

					// execute statement for either create or update
					_, err = stmt.Exec(
						k.FirstName,
						k.LastName,
						k.Email,
						k.Phone,
						k.Gender,
						k.Relationship,
						k.Address,
					)
					kins = append(kins, k.ID)
				}

			}
		}

		// generate id for guarantors
		var guarantors []string

		if len(user.Guarantor) > 0 {
			for i, g := range user.Guarantor {
				// check if kin already has id
				if g.ID == "" {
					g.ID = helper.GenerateUUID().String()
					g.BorrowerId = currentUser.Id
					// edit original aray
					user.Guarantor[i].ID = g.ID
					user.Guarantor[i].BorrowerId = g.BorrowerId
					// insert guarantor into table
					stmt, err := g.CreateGuarantor()
					if err != nil {
						helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
					}
					defer stmt.Close()

					_, err = stmt.Exec(
						g.ID,
						g.BorrowerId,
						g.FirstName,
						g.LastName,
						g.Email,
						g.Phone,
						g.Gender,
						g.Nin,
						g.Income,
						g.Signature,
						g.Address,
					)

					guarantors = append(guarantors, g.ID)
				} else {
					// update
					stmt, err := g.UpdateGuarantor()
					if err != nil {
						helper.SendResponse(w, http.StatusInternalServerError, false, "error saving to db", nil, err)
						return
					}
					defer stmt.Close()

					_, err = stmt.Exec(
						g.FirstName,
						g.LastName,
						g.Email,
						g.Phone,
						g.Gender,
						g.Nin,
						g.Income,
						g.Signature,
						g.Address,
						g.ID,
						g.BorrowerId,
					)
					guarantors = append(guarantors, g.ID)
				}
			}
		}

		stmt, err := user.UpdateBorrower()
		if err != nil {
			helper.SendResponse(w, http.StatusInternalServerError, false, "error encoutered", nil, err)
			return
		}

		defer stmt.Close()

		//  calculate user profile progress
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
			pq.Array(user.Offences),
			user.JailTime,
			pq.Array(kins),
			pq.Array(guarantors),
			user.Nin,
			user.Bvn,
			user.BankName,
			user.AccountNumber,
			user.Identification,
			user.Progress,
			user.CreditScore,
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
