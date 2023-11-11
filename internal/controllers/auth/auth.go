package auth

import (
	// "fmt"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/miriam-samuels/loan-management-backend/internal/helper"
	"github.com/miriam-samuels/loan-management-backend/internal/model/v1/auth"
	"github.com/miriam-samuels/loan-management-backend/internal/model/v1/user"
)

func UserSignUp(w http.ResponseWriter, r *http.Request) {
	cred := &auth.SignUpCred{}
	helper.ParseRequestBody(w, r, cred)

	// TODO: Validate request body

	// check if user already exists
	var exists bool
	cred.CheckUser(&exists, w)
	if exists {
		helper.SendJSONResponse(w, http.StatusBadRequest, false, "user exist", nil)
		return
	}

	// encrypt user password
	encryptedPass, err := helper.Encrypt(cred.Password)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "error occured"+err.Error(), nil)
		log.Fatal(err)
		return
	}

	// generate uuid for new user
	userId := helper.GenerateUUID()

	// generate session token
	token, err := helper.SignJWT(userId.String(), cred.Role)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "unable to generate sesson", nil)
		fmt.Printf("Could not generate token for user:: %v", err)
		return
	}

	// Insert user into db
	cred.CreateUser(userId, encryptedPass, w)
	
	//  check if the user is a borrower then insert into the borrowes table
	if cred.Role == "borrower" {
		cred.CreateBorrower(userId, w)
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

}

func UserSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	cred := &auth.SignInCred{}
	helper.ParseRequestBody(w, r, cred)

	// TODO: Validate request body

	// variable to store data returned by db if user is found
	var user user.User
	var passwordHash string

	// get user from db
	row := cred.FindUserByMail(w)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &passwordHash, &user.Role)
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
	err = helper.CompareHashAndString(passwordHash, cred.Password)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusBadRequest, false, "incorrect password", nil)
		fmt.Printf("incorrect password: %v", err)
		return
	}

	// generate session token
	token, err := helper.SignJWT(user.ID, user.Role)
	if err != nil {
		helper.SendJSONResponse(w, http.StatusInternalServerError, false, "unable to generate sesson", nil)
		fmt.Printf("error generating token %v", err)
		return
	}

	res := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	helper.SendJSONResponse(w, http.StatusOK, true, "user login successful", res)
}
