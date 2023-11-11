package auth

type SignUpCred struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type SignInCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}