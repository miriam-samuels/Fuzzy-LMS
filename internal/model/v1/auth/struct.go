package auth

type SignUpCred struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role"`
}
 
type SignInCred struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}