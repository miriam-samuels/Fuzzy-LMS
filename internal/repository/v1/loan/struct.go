package loan

import "time"

type Loan struct {
	ID               string    `json:"id"`
	LoanID           string    `json:"loanId"`
	BorrowerId       string    `json:"borrowerid "`
	Type             string    `json:"type"`
	Term             string    `json:"term"`
	Amount           float64   `json:"amount"`
	Purpose          string    `json:"purpose"`
	HasCollateral    bool      `json:"has_collateral"`
	Collateral       string    `json:"collateral"`
	CollateralDocs   string    `json:"collateral_docs"`
	Status           string    `json:"status"`
	Creditworthiness float32   `json:"creditworthiness"`
	CreatedAt        time.Time `json:"created_at"`
}
