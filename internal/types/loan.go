package types

type LoanApplication struct {
	LoanID         string   `json:"loanId"`
	BorrowerId        string   `json:"borrowerId "`
	Borrower    Borrower `json:"borrower"`
	LoanType    string   `json:"loan_type"`
	LoanTerm    string   `json:"loan_term"`
	LoanAmount  string   `json:"loan_amount"`
	LoanPurpose string   `json:"loan_purpose"`
}
