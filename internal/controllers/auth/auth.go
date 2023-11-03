package auth

import (
	// "fmt"
	"database/sql"
	"fmt"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
	"log"
	"net/http"
)

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	cred := &types.SignUpCred{}
	helper.ParseRequestBody(w, r, cred)

	// TODO: Validate request body

	// check if user already exists
	var exists bool
	err := database.LoanDb.QueryRow("SELECT 1 FROM users WHERE email=$1", cred.Email).Scan(&exists)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encoutered", nil)
		return
	}

	if exists {
		helper.SendJSONResponse(w, http.StatusBadRequest, false, "user exist", nil)
		return
	}

	// encrypt user password
	encryptedPass, err := helper.Encrypt(cred.Password)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error encrypting password:"+err.Error(), nil)
		log.Fatal(err)
		return
	}

	// generate uuid for new user
	userId := helper.GenerateUUID()

	// generate session token
	token, err := helper.SignJWT(userId.String())
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "unable to generate sesson", nil)
		fmt.Printf("Could not generate token for user:: %v", err)
		return
	}

	// prepare query statement to insert new user into db -- this table is mostly used for authentication
	stmt := helper.Prepare("INSERT INTO users (id, firstname, lastname, email, password, role) VALUES ($1, $2, $3, $4, $5, $6)", w)

	defer stmt.Close()

	//execute statement
	result, err := stmt.Exec(userId, cred.FirstName, cred.LastName, cred.Email, encryptedPass, cred.Role)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
		return
	}

	//  check if the user is a borrower then insert into the borrowes table
	if cred.Role == "borrower" {
		stmt = helper.Prepare("INSERT INTO borrowers (id, firstname, lastname, email) VALUES ($1, $2, $3, $4)", w)
		
		defer stmt.Close()

		result, err = stmt.Exec(userId, cred.FirstName, cred.LastName, cred.Email)
		if err != nil {
			helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error saving to db", nil)
			return
		}
	}
	// Form response object
	res := map[string]interface{}{
		"token":     token,
		"id":        userId,
		"firstname": cred.FirstName,
		"lastname":  cred.LastName,
		"email":     cred.Email,
		"role":      cred.Role,
	}
	helper.SendJSONResponse(w, http.StatusOK, true, "user signup successful", res)
	log.Println(result)
}

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cred := &types.SignInCred{}
	helper.ParseRequestBody(w, r, cred)

	// TODO: Validate request body

	// check if email exists in db
	row := database.LoanDb.QueryRow("SELECT id, firstname, lastname, email, password, role FROM users WHERE email= $1", cred.Email)

	// variable to store data returned by db if user is found
	var user types.User
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			// send response that user does not exist
			helper.SendJSONResponse(w, http.StatusBadRequest, false, "user does not exist", nil)
			return
		}
		// if unknow error
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error getting user", nil)
		fmt.Printf("error getting user %v", err)
		return
	}

	// confirm user password is correct
	err = helper.CompareHashAndString(user.Password, cred.Password)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusBadRequest, false, "incorrect password", nil)
		fmt.Printf("incorrect password: %v", err)
		return
	}

	// generate session token
	token, err := helper.SignJWT(user.ID)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "unable to generate sesson", nil)
		fmt.Printf("error generating token %v", err)
		return
	}

	res := map[string]interface{}{
		"token": token,
		"user": map[string]string{
			"id":        user.ID,
			"email":     user.Email,
			"firstname": user.FirstName,
			"lastname":  user.LastName,
			"role":      user.Role,
		},
	}

	helper.SendJSONResponse(w, http.StatusOK, true, "user login successful", res)
}
