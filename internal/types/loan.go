package types

type LoanApplication struct {
	ID         string `json:"id"`
	LoanID     string `json:"loanId"`
	BorrowerId string `json:"borrowerId "`
	Type       string `json:"type"`
	Term       string `json:"term"`
	Amount     string `json:"amount"`
	Purpose    string `json:"purpose"`
}
