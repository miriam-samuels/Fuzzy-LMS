package user

type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
}

type Borrower struct {
	ID             string      `json:"id,omitempty"`
	FirstName      string      `json:"firstname,omitempty"`
	LastName       string      `json:"lastname,omitempty"`
	Email          string      `json:"email,omitempty"`
	Phone          string      `json:"phone,omitempty"`
	BirthDate      string      `json:"birth_date,omitempty"`
	Gender         string      `json:"gender,omitempty"`
	Nationality    string      `json:"nationality,omitempty"`
	StateOrigin    string      `json:"state_origin,omitempty"`
	Address        string      `json:"address,omitempty"`
	Passport       string      `json:"passport,omitempty"`
	Signature      string      `json:"signature,omitempty"`
	Job            string      `json:"job,omitempty"`
	JobTerm        uint8       `json:"job_term,omitempty"`
	Income         float64     `json:"income,omitempty"`
	Deck           string      `json:"deck,omitempty"`
	HasCriminalRec bool        `json:"has_criminal_record,omitempty"`
	Offences       []string    `json:"offences,omitempty"`
	JailTime       uint8       `json:"jail_time,omitempty"`
	Kin            []NextOfKin `json:"kin,omitempty"`
	Guarantor      []Guarantor `json:"guarantor,omitempty"`
	Nin            string      `json:"nin,omitempty"`
	Bvn            string      `json:"bvn,omitempty"`
	BankName       string      `json:"bank_name,omitempty"`
	AccountNumber  string      `json:"account_number,omitempty"`
	Identification string      `json:"identification,omitempty"`
	LoanIds        []string    `json:"loan_ids,omitempty"`
	Progress       uint8       `json:"progress,omitempty"`
	CreditScore    uint16      `json:"credit_score,omitempty"`
}

type NextOfKin struct {
	ID           string `json:"id,omitempty"`
	BorrowerId   string `json:"borrowerid,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Relationship string `json:"relationship,omitempty"`
	Address      string `json:"address,omitempty"`
}

type Guarantor struct {
	ID         string `json:"id,omitempty"`
	BorrowerId string `json:"borrowerid,omitempty"`
	FirstName  string `json:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Income     string `json:"income,omitempty"`
	Nin        string `json:"nin,omitempty"`
	Signature  string `json:"signature,omitempty"`
	Address    string `json:"address,omitempty"`
}
