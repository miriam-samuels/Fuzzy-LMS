package types

type Loan struct {
	ID         string `json:"id"`
	LoanID     string `json:"loanId"`
	BorrowerId string `json:"borrowerid "`
	Type       string `json:"type"`
	Term       string `json:"term"`
	Amount     string `json:"amount"`
	Purpose    string `json:"purpose"`
	Status     string `json:"status"`
}
