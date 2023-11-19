package user

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

type Borrower struct {
	ID             string      `json:"id"`
	FirstName      string      `json:"firstname"`
	LastName       string      `json:"lastname"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	BirthDate      string      `json:"birth_date"`
	Gender         string      `json:"gender"`
	Nationality    string      `json:"nationality"`
	StateOrigin    string      `json:"state_origin"`
	Address        string      `json:"address"`
	Passport       string      `json:"passport"`
	Signature      string      `json:"signature"`
	Job            string      `json:"job"`
	JobTerm        uint8       `json:"job_term"`
	Income         string      `json:"income"`
	Deck           string      `json:"deck"`
	HasCriminalRec bool        `json:"has_criminal_record"`
	Offences       []string    `json:"offences"`
	JailTime       uint8       `json:"jail_time"`
	Kin            []NextOfKin `json:"kin"`
	Guarantor      []Guarantor `json:"guarantor"`
	Nin            string      `json:"nin"`
	Bvn            string      `json:"bvn"`
	BankName       string      `json:"bank_name"`
	AccountNumber  string      `json:"account_number"`
	Identification string      `json:"identification"`
	LoanIds        []string    `json:"loan_ids"`
	Progress			uint8			`json:"progress"`
}

type NextOfKin struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Relationship string `json:"relationship"`
	Address      string `json:"address"`
}

type Guarantor struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Income    string `json:"income"`
	Nin       string `json:"nin"`
	Signature string `json:"signature"`
	Address   string `json:"address"`
}
