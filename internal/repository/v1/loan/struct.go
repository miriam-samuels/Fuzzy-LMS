package loan

type Loan struct {
	ID            string `json:"id"`
	LoanID        string `json:"loanId"`
	BorrowerId    string `json:"borrowerid "`
	Type          string `json:"type"`
	Term          string `json:"term"`
	Amount        string `json:"amount"`
	Purpose       string `json:"purpose"`
	HasCollateral bool   `json:"has_collateral"`
	Collateral    string `json:"collateral"`
	CollateralDocs string `json:"collateral_docs"`
	Status        string `json:"status"`
}
